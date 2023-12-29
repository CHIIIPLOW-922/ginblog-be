package minio

import (
	"context"
	"ginblog-be/settings"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
)

const URL = "http://192.168.194.35:9000/test/"

func UploadObject(file multipart.File, fileReader *multipart.FileHeader) (string, error) {
	flag, _ := MinioClient.BucketExists(context.Background(), settings.Conf.BucketName)
	if flag != true {
		err := MinioClient.MakeBucket(context.Background(), settings.Conf.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return "创建Bucket出错:" + err.Error(), err
		}
	}
	date := time.Now().Format("2006-01-02")
	// 获取文件名
	objectName := date + "/" + fileReader.Filename
	info, err := MinioClient.PutObject(context.Background(), settings.Conf.BucketName, objectName, file, fileReader.Size, minio.PutObjectOptions{})
	if err != nil {

		return "上传文件出错:" + err.Error(), err
	}
	filePath := URL + info.Key
	return filePath, nil
}
