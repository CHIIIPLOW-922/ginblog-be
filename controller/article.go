package controller

import (
	"ginblog-be/dao/mysql"
	"ginblog-be/enum/code"
	"ginblog-be/result"
	"github.com/gin-gonic/gin"
)

func GetAllArticle(c *gin.Context) {
	article, total := mysql.GetAllArticle()
	if article == nil {
		result.ResErr(c, code.CodeBadRequest)
		return
	}
	result.ResOk(c, code.CodeSuccess, article, total)
}
