package routers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type Route struct {
	path       string
	httpMethod string
	Method     reflect.Value
	Args       []reflect.Type
}

var Routes = []Route{}

func Register(controller interface{}) bool {
	controllerName := reflect.TypeOf(controller).String()
	fmt.Print(controllerName)
	module := controllerName
	if strings.Contains(controllerName, ".") {
		module = controllerName[strings.Index(controllerName, ".")+1:]
	}
	fmt.Println("module=", module)
	v := reflect.ValueOf(controller)
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		action := v.Type().Method(i).Name
		path := fmt.Sprintf("/%s/%s", module, action)
		params := make([]reflect.Type, 0, v.NumMethod())
		fmt.Println(params)
		httpMethod := "POST"
		if strings.Contains(action, "Get") {
			httpMethod = "GET"
		}
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
			fmt.Println(method.Type().In(j))
		}
		fmt.Println("params=", params)
		fmt.Println("httpMethod=", httpMethod)
		fmt.Println("action=", action)
		route := Route{path: path, httpMethod: httpMethod, Method: method, Args: params}
		Routes = append(Routes, route)
	}
	fmt.Println("Routes=", Routes)
	return true
}

func Bind(r *gin.Engine) {
	for _, route := range Routes {
		if route.httpMethod == "GET" {
			r.GET(route.path, match(route.path, route))
		}
		if route.httpMethod == "POST" {
			r.POST(route.path, match(route.path, route))
		}
	}
}

func match(path string, route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := strings.Split(path, "/")
		fmt.Println("fields,len(fields)=", fields, len(fields))
		if len(fields) < 3 {
			return
		}
		if len(Routes) > 0 {
			arguments := make([]reflect.Value, 1)
			arguments[0] = reflect.ValueOf(c)
			fmt.Println("arguments=", arguments)
			route.Method.Call(arguments)
		}
	}
}

func SetupRoutes() *gin.Engine {

	r := gin.Default()
	// r.Group("api/v1")
	Bind(r)
	return r
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
	// return r
}
