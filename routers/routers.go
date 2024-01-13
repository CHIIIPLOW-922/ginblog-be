package routers

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type Route struct {
	path       string
	httpMethod string
	Method     reflect.Value
	Args       []reflect.Type
}

var Routes = []Route{}

/*func Register(controller interface{}) bool {
	controllerName := reflect.TypeOf(controller).String()
	fmt.Print(controllerName)
	var module string
	if strings.Contains(controllerName, ".") {
		module = controllerName
	}make()
}*/

func SetupRoutes() *gin.Engine {

	r := gin.Default()
	r.Group("api/v1")
	//{
	//	v1.GET()
	//	v1.GET("/getAllComments", controller.GetAllComments)
	//	v1.POST("/uploadFile", controller.UploadFile)
	//	v1.POST("/saveUser", controller.SaveUser)
	//	v1.GET("/getAllArticle", controller.GetAllArticle)
	//	v1.POST("/removeFile", controller.RemoveFile)
	//	v1.GET("/lock", controller.LockTest)
	//	v1.GET("/unlock", controller.UnLockTest)
	//	v1.GET("/downFile", controller.DownloadFile)
	//}
	return r
}
