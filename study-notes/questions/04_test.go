package main

import (
	"fmt"
	"testing"
)

func Test_example04(t *testing.T) {
	example04()
}

//range 表达式是副本参与循环，就是说例子中参与循环的是 a 的副本，而不是真正的 a。
//就这个例子来说，假设 b 是 a 的副本，则 range 循环代码是这样的：
//因此无论 a 被如何修改，其副本 b 依旧保持原值，并且参与循环的是 b，
//因此 v 从 b 中取出的仍旧是 a 的原值，而非修改后的值。
func example04() {
	var a = []int{1, 2, 3, 4, 5}
	r := make([]int,5)

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

//如果想要 r 和 a 一样输出，修复办法：
//修复代码中，使用 *[5]int 作为 range 表达式，其副本依旧是一个指向原数组 a 的指针，
//因此后续所有循环中均是 &a 指向的原数组亲自参与的，因此 v 能从 &a 指向的原数组中取
//出 a 修改后的值。
//func example04() {
//	var a = [5]int{1, 2, 3, 4, 5}
//	var r [5]int
//
//	for i, v := range &a {
//		if i == 0 {
//			a[1] = 12
//			a[2] = 13
//		}
//		r[i] = v
//	}
//	fmt.Println("r = ", r)
//	fmt.Println("a = ", a)
//}
