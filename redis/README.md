## REDIS

#### 一、如何使用锁

```
......

lock := redis.GetRedisLock(key, time.Duration(5)*time.Second)
err := lock.Lock()
if err != nil {
    logs.Fatal(err)
    return
}
defer lock.Unlock()

......
```
