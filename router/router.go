package router

import (
	"ginblog-be/api"
	"ginblog-be/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 初始化路由
	// 添加路由处理函数
	// 注册中间件
	// 启动服务器
	r := gin.Default()
	r.GET("/test", api.Test)
	r.Run(utils.HttpPort)
}
