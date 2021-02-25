package main

import (
	"testing"
)

func Test_example1001(t *testing.T) {
	example1001()
}

func example1001() {
	//Q1:
	//使用 cap() 获取 map 的容量。
	//1.使用 make 创建 map 变量时可以指定第二个参数，不过会被忽略。
	//2.cap() 函数适用于数组、数组指针、slice 和 channel，不适用于 map，可以使用 len() 返回 map 的元素个数

	//m := make(map[string]int,2)
	//cap(m) // error

	//Q2:
	//nil 用于表示 interface、函数、maps、slices 和 channels 的“零值”。如果不指定变量的类型，
	//编译器猜不出变量的具体类型，导致编译错误。
	//var x = nil //error
	//  _ = x
}
