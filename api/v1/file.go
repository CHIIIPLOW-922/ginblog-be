package v1

import (
	"context"
	"log"
	"net/http"

	"github.com/CHIIIPLOW-922/ginblog-be/ossconfig"
	"github.com/CHIIIPLOW-922/ginblog-be/utils"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	minioClient := ossconfig.InitMinioClient()
	flag, err := minioClient.BucketExists(context.Background(), utils.BucketName)
	if err != nil {
		log.Fatalf("判断Bucket是否存在错误：%s", err.Error())
	}
	if !flag {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Bucket不存在",
			"flag":  flag,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Bucket存在",
		"flag":    flag,
	})

}
