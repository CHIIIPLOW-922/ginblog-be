package router

import (
	"ginblog-be/api"
	"ginblog-be/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 初始化路由
	// 添加路由处理函数
	// 注册中间件
	// 启动服务器
	r := gin.Default()
	r.GET("/test", api.Test)
	r.GET("/getAllUser", api.GetUsers)
	r.GET("/getAllComments", api.GetComments)
	r.GET("/getAllCategories", api.GetCategories)
	r.POST("/uploadFile", api.UploadFile)
	r.GET("/getUserById", api.GetUserById)
	r.GET("/getFiles", api.GetFiles)
	r.GET("/downloadFile", api.DownloadFile)
	err := r.Run(utils.HttpPort)
	if err != nil {
		log.Fatal("启动服务器失败：", err)
	}
}
