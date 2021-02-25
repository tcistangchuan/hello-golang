package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int

type Lock struct {
	ch chan struct{}
}

func NewLock() *Lock {
	return &Lock{
		ch: make(chan struct{}, 1),
	}
}

func (t *Lock) Lock() {
	<-t.ch
}

func (t *Lock) Unlock() {
	t.ch <- struct{}{}
}

func add(n int) {
	tmp := counter
	tmp += n
	counter = tmp

}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	t := NewLock()
	t.ch <- struct{}{}

	for i := 0; i < 10; i++ {
		go func(j int) {
			t.Lock()
			defer t.Unlock()
			add(j)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(counter)
}