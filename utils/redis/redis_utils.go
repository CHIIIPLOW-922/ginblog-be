package redis

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisLock struct {
	client     *redis.Client
	key        string
	expiration time.Duration // 过期时间，防止宕机或者异常
}

func NewLock(key string, expiration time.Duration) *RedisLock {
	return &RedisLock{
		client:     client,
		key:        key,
		expiration: expiration,
	}
}

func GetLock(key string) *RedisLock {
	return &RedisLock{
		client:     client,
		key:        key,
		expiration: 0,
	}
}

const unLockScript = `
if (redis.call("get", KEYS[1]) == KEYS[2]) then
	redis.call("del", KEYS[1])
	return true
end
return false
`

func (l *RedisLock) RLock(id string) (bool, error) {
	return l.client.SetNX(l.key, id, l.expiration).Result()
}

func (l *RedisLock) RUnLock(id string) error {
	_, err := l.client.Eval(unLockScript, []string{l.key, id}).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	return nil
}
