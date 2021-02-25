package main

import "fmt"

/**
select 随机执行一个可运行的 case。每个 case 必须是一个 chan 通信操作，要么是发送要么是接收。
select 会循环检测条件，如果有满足则执行并退出，否则一直循环检测。

select {
    case communication clause  :
       statement(s);
    case communication clause  :
       statement(s);
	default : [可选]
	statement(s);
}
*/

func main(){
	select {
	default:
			fmt.Println(111)
	}
	fmt.Println("123")
}