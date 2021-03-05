package main

import (
	"fmt"
	"time"
)

func produce(ch chan<- int){
	for i:=1;i<=10;i++{
		fmt.Println("生产了一个：",i)
		ch<-i
	}
}

func consumer(ch <-chan int){
	for i := 1; i <= 10; i++ {
		v := <-ch
		fmt.Println("Receive:", v)
	}
	//for{
	//	select {
	//	case d:=<-ch:
	//		fmt.Println("消费了一个：",d)
	//	}
	//}


}

func main(){
	ch := make(chan int)
	go produce(ch)
	go consumer(ch)
	select {
		case <-time.NewTimer(2*time.Second).C:
			fmt.Println("结束")
	}
}