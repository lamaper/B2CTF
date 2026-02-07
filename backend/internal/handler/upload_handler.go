// ----------------------------------------------------------------------------
// Copyright (c) 2026 lamaper
// 创建日期: 2026-01-21
// 最后修改: 2026-02-08
// 描述: 文件上传相关的HTTP处理函数
// ----------------------------------------------------------------------------
package handler

import (
	"net/http"
	"os"
	"path/filepath"

	"B2CTF/backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadFile 处理文件上传（仅限管理员）
func UploadFile(c *gin.Context) {
	// 1. 权限检查：只有管理员能上传文件
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "没有权限，仅管理员可以上传文件",
		})
		return
	}

	// 2. 获取上传的文件
	// "file" 是前端表单里的字段名 (formData.append('file', ...))
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请提供文件",
		})
		return
	}

	// 3. 文件大小检查（限制为 500MB）
	const maxFileSize = 500 * 1024 * 1024 // 500MB
	if file.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "文件过大，最大允许 500MB",
		})
		return
	}

	// 4. 为了防止文件名冲突，使用 UUID + 原始扩展名
	ext := filepath.Ext(file.Filename) // 获取 .zip / .exe 等
	newFileName := uuid.New().String() + ext

	// 5. 确保存储目录存在
	savePath := "./uploads"
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		if err := os.Mkdir(savePath, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "无法创建上传目录",
			})
			return
		}
	}

	// 6. 保存文件到指定路径
	// 最终路径例如: ./uploads/550e8400-e29b-....zip
	dst := filepath.Join(savePath, newFileName)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "文件保存失败: " + err.Error(),
		})
		return
	}

	// 7. 返回可访问的 URL
	// 注意：这里返回的是相对路径，前端可以直接拼接到域名后
	fileURL := "/uploads/" + newFileName

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传成功",
		"data": gin.H{
			"url":      fileURL,
			"filename": newFileName,
		},
	})
}
