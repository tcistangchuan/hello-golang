package main

import (
	"log"
	"time"
)

// 向一个通道发送一个值来实现单对单通知
// 1.主协程通知子协程
// 2.子协程通知主协程
//https://gfw.go101.org/article/channel-use-cases.html
func main() {
	example_1()
}
func example_1() {
	// 子协程通知主协程
	ch := make(chan int)
	go func() {
		// ...
		time.Sleep(5 * time.Second)
		ch <- 1
	}()
	for {
		// 模拟一个后台程序
		select {
		case <-ch:
			log.Println("监控结束")
			return
		case <-time.Tick(2 * time.Second): //time.Sleep()会阻塞整个goroutine
			// ...
			log.Println("监控中...")
		}
	}
}

func example_2() {
	// 主协程通知子协程
	ch := make(chan int)
	go func() {
		// 模拟一个后台程序
		for {
			select {
			case <-ch:
				log.Println("监控结束")
				return
			case <-time.Tick(2 * time.Second): // time.Sleep()会阻塞整个goroutine
				// ...
				log.Println("监控中...")
			}
		}
	}()
	log.Println("start..")
	time.Sleep(time.Second)
	ch <- 1
}
