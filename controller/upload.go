package controller

import (
	"ginblog-be/enum/code"
	"ginblog-be/result"
	"ginblog-be/utils/minio"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, fileReader, err := c.Request.FormFile("file")
	if err != nil {
		result.ResErrWithMsg(c, code.CodeBadRequest, "文件参数错误")
		return
	}
	// 处理文件
	filaName, err := minio.UploadObject(file, fileReader)
	if err != nil {
		result.ResErrWithMsg(c, code.CodeBadRequest, "上传失败")
		return
	}
	result.ResOk(c, code.CodeSuccess, filaName, nil)
}
