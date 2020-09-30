package main

import (
	"fmt"
	"testing"
)

func Test_example0901(t *testing.T) {
	example0901()
}

//字面量初始化切片时候，可以指定索引，没有指定索引的元素会在前一个索引基础
//之上加一，所以输出[1 0 2 3]，而不是[1 3 2]。
func example0901() {
	var x = []int{4: 4, 3, 0: 1}
	fmt.Println(x)
}
