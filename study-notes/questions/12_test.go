package main

import (
	"fmt"
	"testing"
)

func Test_example1201(t *testing.T) {
	example1201()
}

type TT struct {
	n int
}

func (t *TT) Set(n int) {
	t.n = n
}

func getTT() TT {
	return TT{}
}

//1.直接返回的 T{} 不可寻址；
//2.不可寻址的结构体不能调用带结构体指针接收者的方法；
func example1201() {
	//getTT().Set(1) // error
	m := getTT()
	m.Set(1)
	fmt.Println(m)
}
