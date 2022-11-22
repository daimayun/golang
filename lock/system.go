package lock

import "time"

// memoryCache 内存缓存
var memoryCache map[interface{}]struct{} = make(map[interface{}]struct{})

// GetSet 读取并加锁
func GetSet(key interface{}, ttl ...uint64) (isExistLock bool) {
	var delTtl uint64
	if len(ttl) > 0 {
		delTtl = ttl[0]
	}
	_, isExistLock = memoryCache[key]
	if isExistLock {
		return
	}
	memoryCache[key] = struct{}{}
	if delTtl == 0 {
		return
	}
	go func(key interface{}, ttl uint64) {
		time.AfterFunc(time.Second*time.Duration(ttl), func() {
			delete(memoryCache, key)
		})
	}(key, delTtl)
	return
}

// Del 删除锁
func Del(key interface{}) {
	delete(memoryCache, key)
}

// Empty 清空所有锁
func Empty() {
	memoryCache = map[interface{}]struct{}{}
}
