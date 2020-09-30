package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel-example 获取最快的协程
func main() {
	rand.Seed(time.Now().UnixNano())
	t := time.Now()
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go makeChan(c)
	}
	fmt.Println("最快返回为：", <-c)
	fmt.Println("执行时间为：", time.Since(t))
	fmt.Println()
	fmt.Println("第二快返回为：", <-c)
	fmt.Println("第二快执行时间为：", time.Since(t))
}

func makeChan(c chan<- int) {
	b := rand.Int31n(20) + 1
	fmt.Println("运行的所耗费时间：", b)
	//b:=2
	time.Sleep(time.Second * time.Duration(b))
	c <- int(b)
}
