package main

import (
	"fmt"
	"reflect"
)

func ttt() {
	var user_id interface{}
	user_id = 123

	var id int
	id = 123

	//这里不能赋值，因为类型不一样
	//id = user_id

	//但是这里可以判断，为什么不同的类型可以判断相等？？？
	if user_id == id {
		fmt.Println("相等", user_id)
	} else {
		fmt.Println("不相等", user_id)
	}
}



type T interface {
	f1()
}

func main() {
	test1()
	test2()
}

func test1() {
	var t T
	var i interface{} = t
	fmt.Println(t == nil, i == t, i == nil)
	fmt.Println(reflect.TypeOf(t), reflect.TypeOf(i))
}

func test2() {
	var t *T
	var i interface{} = t
	fmt.Println(t == nil, i == t, i == nil) // I == (*T)(nil)
	fmt.Println(reflect.TypeOf(t), reflect.TypeOf(i))
}