package main

import "fmt"

type I interface {
	Get() int
	Set(int)
}

type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {
	s.Age = age
}

func f(i I) {
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%+v\n", i)
	fmt.Printf("%#v\n", i)
	i.Set(10)
}

func main() {
	//interface 的变量中存储的是实现了 interface 的类型的对象值
	//go 会自动进行 interface 的检查，并在运行时执行从其他类型到
	//interface 的自动转换，即使实现了多个 interface，go 也会在使
	//用对应 interface 时实现自动转换，这就是 interface 的魔力所在。
	s := S{}
	f(&s)
	fmt.Println(s.Age)
}
