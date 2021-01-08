package singleton

import (
	"sync"
)

var (
	mutex     sync.Mutex
	singletonInstance *singleton
	once      sync.Once
)

type singleton struct {
	Age int
	Name string
}

// 这种单列模式 弊端：会大量创建锁，释放锁，带来不必要的开销
func GetInstance() *singleton {
	mutex.Lock()
	defer mutex.Unlock()
	if singletonInstance == nil {
		return &singleton{}
	}
	return singletonInstance
}

// 正确的单列模式
func GetInstance2() *singleton {
	if singletonInstance == nil {
		once.Do(func() {
			singletonInstance = &singleton{}
		})
	}
	return singletonInstance
}
