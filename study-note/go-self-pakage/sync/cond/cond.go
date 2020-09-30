package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

//小结
//sync.Cond 不是一个常用的同步机制，在遇到长时间条件无法满足时，与使用 for {} 进行
//忙碌等待相比，sync.Cond 能够让出处理器的使用权。在使用的过程中我们需要注意以下问题：
//
//sync.Cond.Wait 方法在调用之前一定要使用获取互斥锁，否则会触发程序崩溃；
//sync.Cond.Signal 方法唤醒的 Goroutine 都是队列最前面、等待最久的 Goroutine；
//sync.Cond.Broadcast 会按照一定顺序广播通知等待的全部 Goroutine；

// 实现广播
func main() {
	//
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c)
	}
	time.Sleep(1 * time.Second)
	go broadcast(c)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	//c.Broadcast() // 唤醒全部
	c.Signal() // 唤醒等待时间最久的一个goroutine
	log.Println("start:")
	c.L.Unlock()
}

func listen(c *sync.Cond) {
	c.L.Lock()
	c.Wait()
	time.Sleep(time.Second * 2)
	log.Println("listen")
	c.L.Unlock()
}
