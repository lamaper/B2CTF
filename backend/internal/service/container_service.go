package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"strconv"

	"B2CTF/backend/internal/db"
	"B2CTF/backend/internal/model"
	dockerPkg "B2CTF/backend/internal/pkg/docker"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0 // 如果转换失败，返回0
	}
	return i
}

// LaunchContainer 启动靶机
func LaunchContainer(userID, challengeID uint) (*model.Container, error) {
	ctx := context.Background()

	// 1. 获取题目、比赛、用户信息
	var chal model.Challenge
	if err := db.DB.First(&chal, challengeID).Error; err != nil {
		return nil, errors.New("题目不存在")
	}
	if !chal.IsDynamic {
		return nil, errors.New("这不是动态题目")
	}

	var comp model.Competition
	db.DB.First(&comp, chal.CompetitionID)

	var user model.User
	db.DB.First(&user, userID)

	// 2. 复用检查
	var existingContainer model.Container
	var found bool

	if comp.Mode == 1 && user.TeamID != 0 {
		// --- 团队模式：查全队 ---
		// 逻辑：查找 containers 表，关联 users 表，条件是 users.team_id = 当前用户的team_id
		err := db.DB.Table("containers").
			Joins("INNER JOIN users ON users.id = containers.user_id").
			Where("users.team_id = ? AND containers.challenge_id = ? AND containers.status = ?", user.TeamID, challengeID, "running").
			Last(&existingContainer).Error
		if err == nil {
			found = true
		}
	} else {
		// --- 个人模式：查自己 ---
		err := db.DB.Where("user_id = ? AND challenge_id = ? AND status = ?", userID, challengeID, "running").
			Last(&existingContainer).Error
		if err == nil {
			found = true
		}
	}

	if found {
		// 检查是否过期 (虽然 status=running，但可能时间到了没来得及清理)
		if time.Now().Before(existingContainer.ExpireTime) {
			// 没过期，直接把旧的甩给前端
			// 前端看到这个数据，应该直接显示连接信息，而不是提示“启动成功”
			return &existingContainer, nil
		}
		// 如果过期了，代码继续往下走，下面会开新的（旧的会被清理逻辑处理）
	}

	// 3. 全局限制
	var runningCount int64

	if comp.Mode == 1 && user.TeamID != 0 {
		// 团队限制：全队加起来只能开 1 个 (或者 N 个)
		db.DB.Table("containers").
			Joins("INNER JOIN users ON users.id = containers.user_id").
			Where("users.team_id = ? AND containers.status = ?", user.TeamID, "running").
			Count(&runningCount)
	} else {
		// 个人限制：自己只能开 1 个
		db.DB.Model(&model.Container{}).
			Where("user_id = ? AND status = ?", userID, "running").
			Count(&runningCount)
	}

	limit := comp.ContainerLimit
		if limit <= 0 {
    	limit = 1
	}

	// 用配置的值进行判断
	if runningCount >= int64(limit) {
 		msg := fmt.Sprintf("资源限制：本次比赛限制同时只能开启 %d 个靶机，请先销毁旧的", limit)
    	if comp.Mode == 1 {
    	    msg = fmt.Sprintf("资源限制：本战队限制同时只能开启 %d 个靶机，请联系队友或先销毁旧的", limit)
    	}
    	return nil, errors.New(msg)
	}

	// =================================================================
	// 4. Docker 启动逻辑 (代码和之前一样)
	// =================================================================
	
	dynamicFlag := "flag{" + uuid.New().String() + "}"

	// ... 配置 Docker 参数 ...
	containerPort := nat.Port(fmt.Sprintf("%d/tcp", chal.ContainerPort))
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: ""}},
		},
		Resources: container.Resources{
			NanoCPUs: 500000000, 
			Memory:   128 * 1024 * 1024,
		},
	}
	config := &container.Config{
		Image: chal.ImageName,
		Env:   []string{"FLAG=" + dynamicFlag},
	}

	// ... 创建并启动 ...
	resp, err := dockerPkg.Cli.ContainerCreate(ctx, config, hostConfig, nil, nil, "")
	if err != nil {
		return nil, fmt.Errorf("创建容器失败: %v", err)
	}
	if err := dockerPkg.Cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return nil, fmt.Errorf("启动容器失败: %v", err)
	}

	// ... 获取端口 ...
	inspect, err := dockerPkg.Cli.ContainerInspect(ctx, resp.ID)
	if err != nil {
		return nil, err
	}
	bindings := inspect.NetworkSettings.Ports[containerPort]
	if len(bindings) == 0 {
		return nil, errors.New("端口映射失败")
	}
	assignedPort := bindings[0].HostPort

	// ... 写入数据库 ...
	newContainer := model.Container{
		UserID:      userID,
		ChallengeID: challengeID,
		ContainerID: resp.ID,
		Flag:        dynamicFlag,
		Port:        convertStringToInt(assignedPort),
		ExpireTime:  time.Now().Add(time.Hour * 1),
		Status:      "running",
	}
	if err := db.DB.Create(&newContainer).Error; err != nil {
		return nil, err
	}

	return &newContainer, nil
}

// TerminateContainer 销毁靶机
func TerminateContainer(userID, challengeID uint) error {
	ctx := context.Background()

	// 1. 查数据库记录
	var c model.Container
	result := db.DB.Where("user_id = ? AND challenge_id = ? AND status = ?", userID, challengeID, "running").First(&c)
	if result.Error != nil {
		return errors.New("没有运行中的靶机")
	}

	// 2. 调用 Docker API 强制删除容器 (Force: true 意味着即使在运行也强杀)
	// 即使 Docker 里容器已经没了（比如被手动删了），这一步报错也不要紧，重点是更新数据库
	_ = dockerPkg.Cli.ContainerRemove(ctx, c.ContainerID, types.ContainerRemoveOptions{Force: true})

	// 3. 更新数据库状态
	c.Status = "destroyed"
	return db.DB.Save(&c).Error
}