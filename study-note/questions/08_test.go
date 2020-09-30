package main

import (
	"fmt"
	"testing"
)

type Param map[string]interface{}

type Show struct {
	*Param
}

func Test_example0801(t *testing.T) {
	example0801()
}
func Test_example0802(t *testing.T) {
	example0802()
}

//字面量初始化切片时候，可以指定索引，没有指定索引的元素会在前一个索引基础
//之上加一，所以输出[1 0 2 3]，而不是[1 3 2]。
func example0801() {
	var x = []int{4: 4, 3, 0: 1}
	fmt.Println(x)
}

// 不能对 nil 的 map 直接赋值，需要使用 make() 初始化。
//但可以使用 append() 函数对为 nil 的 slice 增加元素。
func example0802() {
	var a []int
	a = append(a, 1)
	fmt.Println(a)

	var b map[int]int
	fmt.Println(b[1])
	b[1] = 1 // panic
}
