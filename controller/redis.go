package controller

import (
	"ginblog-be/enum/code"
	"ginblog-be/result"
	"ginblog-be/utils/redis"
	"github.com/gin-gonic/gin"
	"time"
)

func LockTest(c *gin.Context) {
	key := c.Query("lockKey")
	rlock := redis.NewLock(key, time.Second*30)
	res, _ := rlock.RLock(key)
	if !res {
		result.ResErrWithMsg(c, code.CodeBadRequest, "加锁失败")
		return
	}
	//defer rlock.RUnLock(key)
	result.ResOkWithMsg(c, code.CodeSuccess, "加锁成功", nil, nil)
}

func UnLockTest(c *gin.Context) {
	key := c.Query("lockKey")
	rlock := redis.GetLock(key)
	if err := rlock.RUnLock(key); err != nil {
		result.ResErrWithMsg(c, code.CodeBadRequest, "解锁失败")
		return
	}
	result.ResOkWithMsg(c, code.CodeSuccess, "解锁成功", nil, nil)
}
