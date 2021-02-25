package main

// 下面这段代码能否通过编译，不能的话原因是什么；如果能，输出什么。
// 答案：不能。
// 不能通过编译，new([]int) 之后的 list 是一个 *[]int 类型的指针，不能对指针执行 append 操作。
// 可以使用 make() 初始化之后再用。同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，
// 不要用 new() 。
func example01() {
	//t:=new([]int)
	//t = append(t,1,2,3,4)
	//log.Println(*t)
}
