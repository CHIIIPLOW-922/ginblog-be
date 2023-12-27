package ossconfig

import (
	"context"
	"log"

	"github.com/CHIIIPLOW-922/ginblog-be/utils"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinioClient() *minio.Client {
	// 初始化一个minio客户端对象
	minioClient, err := minio.New(utils.EndPoint, &minio.Options{
		Creds: credentials.NewStaticV4(utils.AccessKey, utils.SecretKey, ""),
	})
	if err != nil {
		log.Fatalf("初始化MinioClient错误：%s", err.Error())
	}
	return minioClient
}

func test() string {
	// 初始化一个minio客户端对象
	minioClient := InitMinioClient()
	flag, err := minioClient.BucketExists(context.Background(), utils.BucketName)
	if err != nil {
		log.Fatalf("判断Bucket是否存在错误：%s", err.Error())
	}
	if !flag {
		return "不存在"
	}
	return "存在"
}
