package controller

import (
	"ginblog-be/dao/mysql"
	"ginblog-be/enum/code"
	"ginblog-be/result"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {

	//查询所有用户
	users, total := mysql.GetAllUsers()
	if users == nil {
		result.ResErr(c, code.CodeBadRequest)
		return
	}
	result.ResOk(c, code.CodeSuccess, users, total)
}
