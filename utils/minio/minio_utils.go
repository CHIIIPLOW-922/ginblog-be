package minio

import (
	"context"
	"errors"
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

func GetObject(objectName string) (*minio.Object, minio.ObjectInfo) {
	object, _ := MinioClient.GetObject(context.Background(), settings.Conf.BucketName, objectName, minio.GetObjectOptions{})
	objectInfo, _ := object.Stat()
	if objectInfo.Size == 0 {
		return nil, objectInfo
	}
	return object, objectInfo
}

func DeleteObject(objectName string) error {
	var objectList []string
	flag, _ := MinioClient.BucketExists(context.Background(), settings.Conf.BucketName)
	if !flag {
		return errors.New("该Bucket不存在")
	}
	objects := MinioClient.ListObjects(context.Background(), settings.Conf.BucketName, minio.ListObjectsOptions{Recursive: true})
	for object := range objects {
		if objectName == object.Key {
			objectList = append(objectList, object.Key)
		}
	}
	if len(objectList) == 0 {
		return errors.New("该对象不存在")
	}
	MinioClient.RemoveObject(context.Background(), settings.Conf.BucketName, objectName, minio.RemoveObjectOptions{})
	return nil
}
