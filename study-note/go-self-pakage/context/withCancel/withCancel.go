package main

import (
	"context"
	"log"
	"time"
)

// context 用WithCancel实现通知
//context.WithCancel 函数能够从 context.Context 中衍生出一个新的子上下文并返回
//用于取消该上下文的函数（CancelFunc）。一旦我们执行返回的取消函数，当前上下文以及它
//的子上下文都会被取消，所有的 Goroutine 都会同步收到这一取消信号
//https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/

func main() {
	// 一对多的通知
	ctx, cancel := context.WithCancel(context.Background())
	go worker(1, ctx)
	go worker(2, ctx)
	go worker(3, ctx)
	<-time.After(4 * time.Second)
	log.Println("开始")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(1 * time.Second)
}

func worker(id int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("结束", id)
			return
		case <-time.Tick(1 * time.Second):
			log.Println("进行中...")
		}
	}
}

// 1对1通知
func forexample_1() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("结束")
				return
			case <-time.Tick(2 * time.Second):
				log.Println("进行中...")
			}
		}
	}()
	time.Sleep(10 * time.Second)
	cancel()
	log.Println("done")
}
