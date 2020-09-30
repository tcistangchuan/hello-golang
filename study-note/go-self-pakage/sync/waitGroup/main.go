package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		// ...
		log.Println("一号工作完成")
		wg.Done()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		// ...
		log.Println("二号工作完成")
		wg.Done()
	}(&wg)
	wg.Wait()
	log.Println("好了，大家都干完了")
}
