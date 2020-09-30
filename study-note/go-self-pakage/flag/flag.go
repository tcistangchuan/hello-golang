package main

import (
	"flag"
	"fmt"
	//"log"
	"time"
)

/**
Go语言内置的flag包实现了命令行参数的解析，flag包使得开发命令行工具更为简单。
参考链接：https://www.liwenzhou.com/posts/Go/go_flag/
# os.Args 编译后来获取命令行参数
# flag.Type(flag名, 默认值, 帮助信息)*Type
*/

//os.Args demo
func main() {
	//os.Args是一个[]string
	//if len(os.Args) > 0 {
	//	for index, arg := range os.Args {
	//		log.Printf("args[%d]=%v\n", index, arg)
	//	}
	//}

	//name := flag.String("name", "张三", "姓名")
	//age := flag.Int("age", 18, "年龄")
	//married := flag.Bool("married", false, "婚否")
	//delay := flag.Duration("d", 0, "时间间隔")
	//log.Println(*name)
	//log.Println(*age)
	//log.Println(*married)
	//log.Println(*delay)

	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
