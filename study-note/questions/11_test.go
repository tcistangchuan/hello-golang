package main

import (
	"fmt"
	"testing"
)

func Test_example1101(t *testing.T) {
	example1101()
}

type T struct {
	N []int
}

func foo(t T) {
	t.N[0] = 100
}

//调用 foo() 函数时虽然是传值，但 foo() 函数中，字段 ls 依旧可以看成是指向底层数组的指针。
func example1101() {
	t := T{[]int{1, 2, 3}}
	foo(t)
	fmt.Println(t)
}
