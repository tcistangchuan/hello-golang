package main

import (
	"log"
	"time"
)

func main() {
	//log.Println("start")
	//t := <-time.After(2*time.Second) // 阻塞
	//log.Println("end:", t)

	// 实现计时器
	for {
		select {
		case <-time.After(2 * time.Second):
			log.Println("done")
		}
	}
}
