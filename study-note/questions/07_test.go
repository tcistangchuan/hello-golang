package main

import (
	"fmt"
	"testing"
)

func Test_example0701(t *testing.T) {
	example0701()
}

//interface 的内部结构，我们知道接口除了有静态类型，还有动态类型和动态值，
//当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。这里的 x 的动
//态类型是 *int，所以 x 不为 nil。
// 解决：
// if == (*int)(nil)
func Foo1(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func example0701() {
	var x *int = nil
	fmt.Println(x == nil)
	Foo1(x)
}
