package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_example1301(t *testing.T) {
	example1301()
}

type N struct {
	age int
}

func (n N) test(i int) {
	n.age = i
}

func example1301() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, _ := range a {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}

func test(i int) {
	fmt.Println(i)
}
