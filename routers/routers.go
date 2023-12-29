package routers

import (
	"ginblog-be/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/getAllUsers", controller.GetAllUsers)
		v1.GET("/getAllComments", controller.GetAllComments)
		v1.POST("/uploadFile", controller.UploadFile)
		v1.POST("/saveUser", controller.SaveUser)
	}
	return r
}
