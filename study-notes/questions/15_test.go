package main

import (
	"fmt"
	"testing"
)

func Test15(test *testing.T){
	//t := []int{1,2} // 这种赋值 cap 为2，append 就会创建新的空间
	t := make([]int,0,3)
	t = append(t,1)
	t = append(t,2)
	var a [][]int
	a = append(a,t)
	a = append(a,t)
	a = append(a,t)
	// a[0] , a[1] ,a[2] 都是对 t 进行append 操作
	a[0] = append(a[0],0)
	a[1] = append(a[1],1)
	a[2] = append(a[2],2)
	//for k,v := range a{
	//	a = append(a,append(v,k))
	//}
	fmt.Println(a)
}
