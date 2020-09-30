package main

import (
	"fmt"
	"reflect"
)

//Go 语言的 reflect 包为我们提供的多种能力，包括如何使用反射来动态修改变量、
//判断类型是否实现了某些接口以及动态调用方法等功能，通过对反射包中方法原理的分
//析能帮助我们理解之前看起来比较怪异、令人困惑的现象。
// 参考：
//https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/
//https://juejin.im/post/6844903559335526407#heading-15

type A interface {
	test1(string) string
	test2(string) string
}

type B struct {
	name string `json:"name2" toml:"test"`
	old  int
}

func main() {
	b := B{}
	b.name = "tc"
	b.old = 18
	/*
		变量到反射对象

		使用 reflect.TypeOf 和 reflect.ValueOf 能够获取 Go 语言中的变量对应的反射对象。
		一旦获取了反射对象，我们就能得到跟当前类型相关数据和操作，并可以使用这些运行时获取的结构执行方法。
	*/
	//t := reflect.TypeOf(b)
	//v := reflect.ValueOf(b)
	//log.Println(t)
	//log.Println(v)
	//log.Println(v.FieldByName("name"))
	//log.Println(v.FieldByName("old"))

	// 反射对象到接口
	//log.Println(v.Interface().(B))

	// 修改结构体里的数据和切片里的数据
	/*
		i := 1
		v := &i
		*v = 10
	*/
	//v := reflect.ValueOf(&(b.name))
	//v.Elem().SetString("test")
	//fmt.Println(b)

	//test := []int{1,2,3}
	//v2 := reflect.ValueOf(&(test[1]))
	//v2.Elem().SetInt(223344)
	//fmt.Println(test)

	// typeof 的操作
	//fmt.Println(reflect.TypeOf(b).String())
	//fmt.Println(reflect.TypeOf(b).Kind())
	//t,_ := reflect.TypeOf(b).FieldByName("name")
	//fmt.Println( t)
	//fmt.Println( t.Tag.Get("toml"))

	// 判断结构体是否实现接口
	//typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	//customErrorPtr := reflect.TypeOf(&CustomError{})
	//customError := reflect.TypeOf(CustomError{})
	//fmt.Println(customErrorPtr.Implements(typeOfError)) // #=> true
	//fmt.Println(customError.Implements(typeOfError)) // #=> false

	// 调用方法
	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	fmt.Println(t)
	argv := make([]reflect.Value, t.NumIn())
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i)
	}
	fmt.Println(argv)
	result := v.Call(argv)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result)          // #=> 1
	fmt.Println(result[0].Int()) // #=> 1

}
func Add(a, b int) int { return a + b }

type CustomError struct{}

func (*CustomError) Error() string {
	return ""
}
