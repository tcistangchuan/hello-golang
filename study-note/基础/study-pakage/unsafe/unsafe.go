package main

import (
	"fmt"
	"unsafe"
)

// Pointer:
//任何类型的指针值都可以转换为 Pointer
//Pointer 可以转换为任何类型的指针值
//uintptr 可以转换为 Pointer
//Pointer 可以转换为 uintptr(非类型安全指针)

func createInt() *int {
	return new(int)
}
func main() {
	p0 := createInt()
	fmt.Printf("%p\n", p0)
	fmt.Println(unsafe.Pointer(p0))
	fmt.Println(uintptr(unsafe.Pointer(p0)))

	var i int
	// 此函数用来取得一个值的尺寸 单位字节
	fmt.Println(unsafe.Sizeof(i))
	// 一个值在内存中的地址对齐保证
	fmt.Println(unsafe.Alignof(i))
	//此函数用来取得一个结构体值的某个字段的地址相对于此结构体值的地址的偏移。
	type A struct {
		name string
		old  int
	}
	var a A
	fmt.Println(unsafe.Offsetof(a.old))

}
