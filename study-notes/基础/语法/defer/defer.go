package main

func main() {
	// defer在一个被延迟调用的函数值是在其调用被推入延迟调用堆栈之前被估值的
	//var f = func () {
	//	fmt.Println(false)
	//}
	//defer f()
	//f = func () {
	//	fmt.Println(true)
	//}

	// defer的在执行的时候panic,而不是被推入延迟调用堆栈的时候
	//defer func() {
	//	panic("123")
	//}()
	//fmt.Println("ok")

}

func test() string {
	defer func() string {
		return "123"
	}()
	return ""
}
