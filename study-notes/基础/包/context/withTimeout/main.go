package main

import (
	"context"
	"log"
	"time"
)

func main() {
	//我们创建了一个过期时间为 2s 的上下文, 如果超过过期时间当前上下文以及它的子上下文都会被取消，所有的 Goroutine 都会同步收到这一取消信号
	//多个 Goroutine 同时订阅 ctx.Done() 管道中的消息，一旦接收到取消信号就立刻停止当前正在执行的工作。
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) //context.WithDeadline() 方法与这个方法类似
	defer cancel()

	log.Println("start")
	go handle(ctx, 4*time.Second)
	select {
	case <-ctx.Done():
		log.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		log.Println("handle", ctx.Err())
	case <-time.After(duration):
		log.Println("process request with", duration)
	}

}
