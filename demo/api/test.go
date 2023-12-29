package api

import (
	"context"
	"ginblog-be/model"
	"ginblog-be/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

const URL = "http://192.168.194.35:9000/test/"

func Test(c *gin.Context) {
	bucketName := c.Query("bucket")
	minioClient := utils.InitMinioClient()
	flag, _ := minioClient.BucketExists(context.Background(), bucketName)
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "存在",
	})
}

func GetUsers(c *gin.Context) {
	data, total := model.GetAllUsers()
	if data != nil {
		utils.ResOk(c, data, total)
		return
	}
	utils.ResErr(c, utils.ERROR)
}

func GetComments(c *gin.Context) {
	data, total := model.GetAllComments()
	if data != nil {
		utils.ResOk(c, data, total)
		return
	}
	utils.ResErr(c, utils.ERROR)
}

func GetCategories(c *gin.Context) {
	data, total := model.GetAllCategories()
	if data != nil {
		utils.ResOk(c, data, total)
		return
	}
	utils.ResErr(c, utils.ERROR)
}

func UploadFile(c *gin.Context) {
	minioClient := utils.InitMinioClient()
	file, fileReader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "文件参数错误",
		})
		return
	}
	date := time.Now().Format("2006-01-02")
	// 获取文件名
	objectName := date + "/" + fileReader.Filename
	// 获取文件后缀
	info, err := minioClient.PutObject(context.Background(), utils.BucketName, objectName, file, fileReader.Size, minio.PutObjectOptions{})
	if err != nil {
		utils.ResErr(c, utils.ERROR)
		return
	}
	filePath := URL + info.Key
	utils.ResOk(c, filePath, 0)

}

func GetFiles(c *gin.Context) {
	minioClient := utils.InitMinioClient()
	var list []string
	info := minioClient.ListObjects(context.Background(), utils.BucketName, minio.ListObjectsOptions{Recursive: true})
	for object := range info {
		if object.Err != nil {
			utils.ResErr(c, utils.ERROR)
			return
		}
		list = append(list, URL+object.Key)

	}
	if list != nil {
		utils.ResOk(c, list, 0)
		return
	}
	utils.ResErr(c, utils.ERROR)

}

func DownloadFile(c *gin.Context) {
	minioClient := utils.InitMinioClient()
	objectName := c.Query("objectName")
	err := minioClient.FGetObject(context.Background(), utils.BucketName, objectName, objectName, minio.GetObjectOptions{})
	if err != nil {
		utils.ResErr(c, utils.ERROR)
		return
	}
	c.Header("Content-Disposition", "attachment; filename="+objectName)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "chunked")
	c.File(objectName)
	utils.ResOk(c, objectName, 0)
	return
}

func GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := model.GetUserById(id)
	if data == (model.User{}) {
		utils.ResErr(c, utils.ERROR)
		return
	}
	utils.ResOk(c, data, 0)
}
