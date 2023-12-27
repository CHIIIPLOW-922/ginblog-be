package router

import (
	"github.com/CHIIIPLOW-922/ginblog-be/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	// 初始化路由
	// 启动服务
	r.Run(utils.HttpPort)
}
