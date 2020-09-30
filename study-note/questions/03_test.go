package main

import (
	"testing"
)

func Test_example03(t *testing.T) {
	example03()
}

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

//编译报错 cannot assign to struct field m["foo"].x in map。
//错误原因：对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y
//的值赋给 X，但 go 中的 map 的 value 本身是不可寻址的。
func example03() {
	//m["foo"].x = 4
	//fmt.Println(m["foo"].x)

}

//有两个解决办法：
//
//1.使用临时变量
//
//
//func main() {
//	tmp := m["foo"]
//	tmp.x = 4
//	m["foo"] = tmp
//	fmt.Println(m["foo"].x)
//}
//2.修改数据结构
//
//type Math struct {
//	x, y int
//}
//
//var m = map[string]*Math{
//	"foo": &Math{2, 3},
//}
//
//func main() {
//	m["foo"].x = 4
//	fmt.Println(m["foo"].x)
//	fmt.Printf("%#v", m["foo"])   // %#v 格式化输出详细信息
//}
