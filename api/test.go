package api

import (
	"context"
	"ginblog-be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	minioClient := utils.InitMinioClient()
	flag, err := minioClient.BucketExists(context.Background(), utils.BucketName)
	if err != nil {
		return
	}
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "存在",
	})
}
