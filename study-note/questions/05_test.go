package main

import (
	"fmt"
	"testing"
)

func Test_example05(t *testing.T) {
	example05()
}

func Test_example0502(t *testing.T) {
	example0502()
}

func change(s ...int) {
	s = append(s, 3)
	fmt.Println("s:", s)
}

//知识点：可变函数、append()操作。Go 提供的语法糖...，可以将
//slice 传进可变函数，不会创建新的切片。第一次调用 change()
//时，append() 操作使切片底层数组发生了扩容，原 slice 的底层
//数组不会改变；第二次调用change() 函数时，使用了操作符[i,j]
//获得一个新的切片，假定为 slice1，它的底层数组和原切片底层数组
//是重合的，不过 slice1 的长度、容量分别是 2、5，所以在 change
//() 函数中对 slice1 底层数组的修改会影响到原切片。
//链接：https://mp.weixin.qq.com/s?__biz=MzI2MDA1MTcxMg==&mid=2648467030&idx=1&sn=9dacebc1e23d5b2dd689c71b0f49e68b&chksm=f2474039c530c92f023747a5774f19cfafd8f76a2bd658adeec51e2ad9df0e21afd3a69a16b5&scene=21#wechat_redirect
func example05() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice) // 12000
	change(slice[0:2]...)
	fmt.Println(slice) // 123
}

func example0502() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
