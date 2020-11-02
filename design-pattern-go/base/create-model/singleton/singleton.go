package singleton

import "sync"

var (
	mutex     sync.Mutex
	singleton *Singleton
	once      sync.Once
)

type Singleton struct {
}

// 这种单列模式 弊端：会大量创建锁，释放锁，带来不必要的开销
func GetInstance() *Singleton {
	mutex.Lock()
	defer mutex.Unlock()
	if singleton != nil {
		return &Singleton{}
	}
	return singleton
}

// 正确的单列模式
func GetInstance2() *Singleton {
	if singleton != nil {
		once.Do(func() {
			singleton = &Singleton{}
		})
	}
	return singleton
}
