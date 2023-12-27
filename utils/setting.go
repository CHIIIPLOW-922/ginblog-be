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
	DbPassword string
	DbName     string

	EndPoint   string
	AccessKey  string
	SecretKey  string
	BucketName string
	UseSSL     bool
)

func init() {
	//加载配置文件
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取失败，请检查文件路径:", err)
	}

	LoadServer(file)
	LoadDatabase(file)
	LoadOss(file)
}

func LoadServer(file *ini.File) {
	HttpPort = file.Section("server").Key("HttpPort").String()
}

func LoadDatabase(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").String()
	DbPort = file.Section("database").Key("DbPort").String()
	DbUser = file.Section("database").Key("DbUser").String()
	DbPassword = file.Section("database").Key("DbPassword").String()
	DbName = file.Section("database").Key("DbName").String()
}

func LoadOss(file *ini.File) {
	EndPoint = file.Section("oss").Key("EndPoint").String()
	AccessKey = file.Section("oss").Key("AccessKey").String()
	SecretKey = file.Section("oss").Key("SecretKey").String()
	BucketName = file.Section("oss").Key("BucketName").String()
	UseSSL = file.Section("oss").Key("UseSSL").MustBool(false)
}
