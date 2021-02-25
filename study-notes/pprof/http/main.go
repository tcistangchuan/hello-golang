package main

import (
	"log"
	"net/http"
	"stuty-go/study-notes/pprof/http/studyPprof"
)

// 默认 studyPprof 的配置
const (
	isOpenPprof = true
)

func main() {
	go test()
	if isOpenPprof {
		studyPprof.BootPprof()
	}
	http.ListenAndServe("0.0.0.0:6060", nil)
}

var a []string

func test() {
	// 模拟阻塞
	var ch chan int
	for {
		select {
		case <-ch:
			log.Printf("ch:%d", <-ch)
		default:
		}
	}
}
