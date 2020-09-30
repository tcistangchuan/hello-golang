package main

import (
	"log"
	"time"
)

// 使用timer定时器，超时后需要重置，才能继续触发。
func main() {
	t := time.NewTimer(time.Duration(2 * time.Second))
	defer t.Stop() // // 不会再发送，通道不会关闭
	log.Println("start")
	for {
		select {
		case <-t.C:
			log.Println("done")
			t.Reset(3 * time.Second)
		}
	}
}
