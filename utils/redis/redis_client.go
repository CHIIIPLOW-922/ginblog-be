package redis

import (
	"fmt"
	"ginblog-be/settings"
	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func InitRedis(config *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password:     config.Password, // no password set
		DB:           config.DB,       // use default DB
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
	})
	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = client.Close()
}
