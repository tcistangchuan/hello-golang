package main

import "fmt"

// 参考地址 https://gfw.go101.org/article/evaluation-orders.html
func main() {
	//m := map[string]int{"Go": 0}
	//s := []int{1, 1, 1}; olds := s
	//n := 2
	//p := &n
	//s, s[2] = []int{2, 2, 2}, 5

	s := []int{1, 1, 1}
	olds := s
	k1 := []int{2, 2, 2}
	k2 := 5
	v1 := &s
	v2 := &s[2]
	*v1 = k1
	*v2 = k2
	fmt.Println(s)
	fmt.Println(olds)
}
