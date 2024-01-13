package controller

import (
	"ginblog-be/enum/code"
	"ginblog-be/result"
	"ginblog-be/utils/minio"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
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

func DownloadFile(c *gin.Context) {
	fileName := c.Query("fileName")
	if fileName == "" {
		result.ResErrWithMsg(c, code.CodeBadRequest, "参数错误")
		return
	}
	object, objectInfo := minio.GetObject(fileName)
	if object == nil {
		result.ResErrWithMsg(c, code.CodeBadRequest, "下载失败")
		return
	}
	downloadFileName := time.Now().Format(time.Stamp) + "/" + objectInfo.Key
	data, _ := ioutil.ReadAll(object)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment;filename="+"\""+downloadFileName+"\"")
	c.Data(http.StatusOK, "application/octet-stream", data)
}

func RemoveFile(c *gin.Context) {
	fileName := c.Query("fileName")
	if fileName == "" {
		result.ResErrWithMsg(c, code.CodeBadRequest, "参数错误")
		return
	}
	if err := minio.DeleteObject(fileName); err != nil {
		result.ResErrWithMsg(c, code.CodeBadRequest, err.Error())
		return
	}
	result.ResOk(c, code.CodeSuccess, fileName, nil)
}
