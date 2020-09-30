package main

import (
	"log"
	"time"
)

// 一对多的通知
func main() {
	ready, done := make(chan int, 2), make(chan int, 2)

	go worker(1, ready, done, time.Second)
	go worker(2, ready, done, 2*time.Second)

	// ready<-1;ready<-2 等同于close()
	close(ready)
	<-done
	<-done
}

func worker(id int, ready <-chan int, done chan<- int, d time.Duration) {

	log.Println("start:", id)
	<-ready

	// ...
	<-time.After(d)

	log.Println("done:", id)
	done <- id
	//a := []int{}

}
