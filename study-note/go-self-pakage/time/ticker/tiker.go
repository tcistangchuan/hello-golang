package main

import (
	"log"
	"time"
)

func main() {
	//t := time.NewTicker(2*time.Second)
	//defer t.Stop() // 不会再发送，通道不会关闭
	//for{
	//	<-t.C
	//	log.Println("done")
	//}

	for {
		select {
		case <-time.Tick(2 * time.Second):
			log.Println("done")
		}
	}

}
