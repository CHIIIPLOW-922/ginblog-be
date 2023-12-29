package controller

import (
	"ginblog-be/dao/mysql"
	"ginblog-be/enum/code"
	"ginblog-be/models"
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

func SaveUser(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		result.ResErrWithMsg(c, code.CodeBadRequest, "入参错误")
		return
	}
	err := mysql.InsertUser(user)
	if err != nil {
		result.ResErrWithMsg(c, code.CodeBadRequest, "保存失败")
		return
	}
	result.ResOk(c, code.CodeSuccess, &user, nil)
}
