package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	HttpPort string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	EndPoint   string
	AccessKey  string
	SecretKey  string
	BucketName string
	UseSSL     bool
)

// 初始化
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadMinio(file)
	LoadDB(file)

}

func LoadServer(file *ini.File) {
	HttpPort = file.Section("server").Key("HttpPort").String()
}

func LoadDB(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").String()
	DbPort = file.Section("database").Key("DbPort").String()
	DbUser = file.Section("database").Key("DbUser").String()
	DbPassWord = file.Section("database").Key("DbPassWord").String()
	DbName = file.Section("database").Key("DbName").String()
}

func LoadMinio(file *ini.File) {
	EndPoint = file.Section("minio").Key("EndPoint").String()
	AccessKey = file.Section("minio").Key("AccessKey").String()
	SecretKey = file.Section("minio").Key("SecretKey").String()
	BucketName = file.Section("minio").Key("BucketName").String()
	UseSSL = file.Section("minio").Key("UseSSL").MustBool(false)
}
