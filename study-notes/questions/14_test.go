package main

import (
	"fmt"
	"testing"
)

type M struct{
	ids []int
}
func example14() {
	m := &M{}
	m.ids = append(m.ids, 1)
	fmt.Println(cap(m.ids))
	m.ids = append(m.ids, 2)
	fmt.Println(cap(m.ids))
	m.ids = append(m.ids, 3)
	fmt.Println(cap(m.ids))
	//m.ids = append(m.ids, 0)
	//fmt.Println(cap(m.ids))


	a := append(m.ids, 4)
	b := append(m.ids, 5)
	c := append(m.ids, 6)

	fmt.Println(m.ids)
	fmt.Println(a, b, c)

	//找到的一个解释：Go 会为你的 slice 提供比你需要的更多的容量，原因是在 slice 的底层，
	//有个不可变动的（immutable）数组（array）在实际起作用。当你要为 slice 添加元素从而让切片
	//的容量更大的时候，实际上是创建了一个新数组，把原来的切片元素和新添加的元素放到新的数组里，并
	//把这个数组作为新 slice 的底层。如果你添加很多数据到 slice 里，就会反复去创建和复制这些数据
	//，影响性能。所以运行时会分配比你期望的更多的容量到 slice，让复制数据这些操作变得不那么频繁。

}


func Test_example14(t *testing.T){
	example14()
}