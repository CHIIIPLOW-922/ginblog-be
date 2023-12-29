package main

import (
	"fmt"
	"ginblog-be/dao/mysql"
	"ginblog-be/routers"
	"ginblog-be/settings"
	"ginblog-be/utils/minio"
)

func main() {

	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := minio.InitMinioClient(settings.Conf.MinioConfig); err != nil {
		fmt.Printf("init minio failed, err:%v\n", err)
		return
	}
	if err := mysql.InitDB(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return // 初始化数据库连接
	}
	//程序退出
	defer mysql.Close()

	r := routers.SetupRoutes()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
