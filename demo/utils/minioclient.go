package utils

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinioClient() *minio.Client {
	minioClient, err := minio.New(EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AccessKey, SecretKey, ""),
		Secure: UseSSL,
	})
	if err != nil {
		log.Fatalf("初始化MinioClient错误：%s", err.Error())
	}
	return minioClient
}
