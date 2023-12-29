package mysql

import (
	"database/sql"
	"fmt"
	"ginblog-be/settings"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB // 声明一个全局的db变量

func InitDB(config *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DB)
	sqlDB, err := sql.Open(config.Driver, dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	// 注册单例
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		// 禁止自动给表名加 "s"
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	db = gormDB
	if err != nil {
		log.Fatal(err)
		return
	}
	return // 返回nil
}

func Close() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
