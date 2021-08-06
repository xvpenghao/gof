package singleton

import (
	"sync"
	"time"
)

var lazySingleton *Singleton

// GetLazyInstance 不支持并发
func GetLazyInstance() *Singleton {
	// 当同时有两个，协程都满足13行时，那么就会创建两个示例
	if lazySingleton == nil {
		// 加这个时间，是用来测试的，并发是否安全的
		time.Sleep(5 * time.Millisecond)
		lazySingleton = &Singleton{}
	}
	return lazySingleton
}

// 利用加锁机制，实现双重检测
var lock sync.Mutex

// GetLazyInstanceV2 支持并发，双重监测
func GetLazyInstanceV2() *Singleton {
	if lazySingleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if lazySingleton == nil {
			time.Sleep(5 * time.Millisecond)
			lazySingleton = &Singleton{}
			return lazySingleton
		}
	}
	return lazySingleton
}

var once sync.Once

// GetLazyInstanceV3 v3 利用 sync.Once来实现
func GetLazyInstanceV3() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			// 改函数只会执行一次
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
