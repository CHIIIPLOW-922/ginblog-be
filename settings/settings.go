package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Mode      string `mapstructure:"mode"`
	Port      int    `mapstructure:"port"`
	Name      string `mapstructure:"name"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int    `mapstructure:"machine_id"`
	// *LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*MinioConfig `mapstructure:"minio"`
}

type MinioConfig struct {
	Endpoint   string `mapstructure:"endpoint"`
	AccessKey  string `mapstructure:"access_key"`
	SecretKey  string `mapstructure:"secret_key"`
	UseSSL     bool   `mapstructure:"usessl"`
	BucketName string `mapstructure:"bucketname"`
}

type MySQLConfig struct {
	Driver       string `mapstructure:"driver"`
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

// type LogConfig struct {
// 	Level      string `mapstructure:"level"`
// 	Filename   string `mapstructure:"filename"`
// 	MaxSize    int    `mapstructure:"max_size"`
// 	MaxAge     int    `mapstructure:"max_age"`
// 	MaxBackups int    `mapstructure:"max_backups"`
// }

var Conf = new(AppConfig)

func Init() error {
	//读取配置文件
	viper.SetConfigFile("./config/config.yaml")
	//读取环境变量
	viper.WatchConfig()
	//监控配置文件变化
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件变化")
		viper.Unmarshal(&Conf)
	})
	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件错误:%d", err))
	}
	//读取配置信息反序列化到Conf变量
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("反序列化配置文件错误:%d", err))
	}
	return err

}
