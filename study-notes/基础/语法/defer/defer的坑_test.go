package main

import (
	"fmt"
	"testing"
)

// 参考：
//defer执行原理
//要想知道为何是这个结果，就得先回答前面的2个问题：
//
//(1)defer和defer后的函数什么时候执行吗？
//(2)defer后函数里的变量值是什么时候计算的吗？
//依次来回答，这2个问题。
//
//问题1：defer在defer语句处执行，defer的执行结果是把defer后的函数压入到栈，等待return或者函数panic后，
//再按先进后出的顺序执行被defer的函数。
//
//问题2：defer的函数的参数是在执行defer时计算的，defer的函数中的变量的值是在函数执行时计算的。
//https://segmentfault.com/a/1190000017043508
//https://juejin.cn/post/6844903679519096846


func test1() (x int) {
	defer fmt.Printf("in defer: x = %d\n", x)
	x = 7
	return 9
}

func test2() (x int) {
	x = 7
	defer fmt.Printf("in defer2: x = %d\n", x)
	return 9
}

func test3() (x int) {
	defer func() {
		fmt.Printf("in defer3: x = %d\n", x)
	}()

	x = 7
	return 9
}

func test4() (x int) {
	defer func(param int) {
		fmt.Printf("in defer x as parameter: x = %d\n", param)
		fmt.Printf("in defer x after return: x = %d\n", x) // 这里的x是值外部的变量
	}(x) //此时 x = 0

	x = 7
	return 9
}

func TestM(t *testing.T) {
	//fmt.Println("test1")
	//fmt.Printf("in main: x = %d\n", test1())
	//fmt.Println("test2")
	//fmt.Printf("in main: x = %d\n", test2())
	//fmt.Println("test3")
	//fmt.Printf("in main: x = %d\n", test3())
	//fmt.Println("test4")
	//fmt.Printf("in main: x = %d\n", test4())
	fmt.Println(returnValues())
	fmt.Println(namedReturnValues())
}


func returnValues() int {
	var result int
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return result
}

func namedReturnValues() (result int) {
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return result
}


