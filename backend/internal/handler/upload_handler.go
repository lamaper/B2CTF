// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-01-21
// 描述: 文件上传相关的HTTP处理函数
// ----------------------------------------------------------------------------
package handler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadFile 处理文件上传
func UploadFile(c *gin.Context) {
	// 1. 获取上传的文件
	// "file" 是前端表单里的字段名 (formData.append('file', ...))
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传文件"})
		return
	}

	// 2. 为了防止文件名冲突（比如两个题目都叫 crackme.exe），我们需要重命名文件
	// 策略：UUID + 原始扩展名
	ext := filepath.Ext(file.Filename) // 获取 .zip / .exe
	newFileName := uuid.New().String() + ext
	
	// 3. 确保存储目录存在
	savePath := "./uploads"
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		os.Mkdir(savePath, os.ModePerm)
	}

	// 4. 保存文件到指定路径
	// 最终路径例如: ./uploads/550e8400-e29b-....zip
	dst := filepath.Join(savePath, newFileName)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}

	// 5. 返回可访问的 URL
	// 注意：这里返回的是相对路径，前端可以直接拼接到域名后
	fileURL := "/uploads/" + newFileName
	
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传成功",
		"data": gin.H{
			"url": fileURL,
		},
	})
}