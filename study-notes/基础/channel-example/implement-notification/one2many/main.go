package main

import (
	"fmt"
	"time"
)

// 一对多的通知
func main() {
	N := 10
	exit := make(chan struct{})
	done := make(chan struct{}, N)

	//start N worker goroutines
	for i := 0; i < N; i++ {
		// 开启10个协程工作
		go func(n int) {
			for {
				select {
				// wait for exit signal
				case <-exit: // 等待结束信号（要么推送消息到exit通道，要么close掉exit通道）
					fmt.Printf("worker goroutine #%d exit\n", n)
					done <- struct{}{} //发送完成任务信号
					return // 结束工作
				case <-time.Tick(time.Second):
					fmt.Printf("worker goroutine #%d is working...\n", n)
				}
			}
		}(i)
	}

	time.Sleep(3 * time.Second)
	//广播通知
	close(exit)

	//wait for all worker goroutines exit
	for i := 0; i < N; i++ {
		<-done
	}

	fmt.Println("main goroutine exit")

}

