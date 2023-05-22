package redis

import (
	"github.com/go-redsync/redsync"
	"time"
)

// GetRedisLock 获取REDIS锁
func GetRedisLock(key string, expireTime time.Duration) *redsync.Mutex {
	return redisLock.NewMutex(key, redsync.SetExpiry(expireTime), redsync.SetTries(1000))
}
