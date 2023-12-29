package controller

import (
	"ginblog-be/dao/mysql"
	"ginblog-be/enum/code"
	"ginblog-be/result"

	"github.com/gin-gonic/gin"
)

func GetAllComments(c *gin.Context) {
	comments, total := mysql.GetAllComments()
	if comments == nil {
		result.ResErr(c, code.CodeBadRequest)
		return
	}
	result.ResOk(c, code.CodeSuccess, comments, total)

}
