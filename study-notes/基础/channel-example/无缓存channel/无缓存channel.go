package main

import (
	"fmt"
	"sync"
)

func main(){
	w := sync.WaitGroup{}
	w.Add(1)
	ch:= make(chan int)
	//fmt.Println(<-ch) // 死锁

	// 必须同时生产和消费，所以开一个新的 goroutine
	go func() {
		ch<-1
		w.Done()
	}()
	fmt.Println(<-ch)
	w.Wait()

}

