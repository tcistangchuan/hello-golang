- new()和make的区别（传入参数和返回参数类型不一样）

  ```go
  只声明未赋值的变量，golang都会自动为其初始化为零值，基础数据类型的零值比较简单，引用类型和指针的零值都为nil，nil类型不能直接赋值，因此需要通过new,make开辟一个内存，或指向一个变量。
  
  new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。
  
  new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。
  
  make(T,args) 返回初始化之后的 T 类型的值。make() 只适用于引用类型 slice、map 和 channel.
  ```


- Struct的值接收者和指针接收者的区别

  ```
  如果方法的接收者是值类型，无论调用者是对象还是对象指针，修改的都是对象的副本，不影响调用者；如果方法的接收者是指针类型，则调用者修改的是指针指向的对象本身。
  例如：
  (1)接收者是值类型,值调用者的func里面无法修改到struct的属性
  (2)接收者是值类型,值调用者不会自动实现接收者是指针类型的方法（实现interface）
  
  使用指针作为方法的接收者的理由：
  方法能够修改接收者指向的值。
  避免在每次调用方法时复制该值，在值的类型为大型结构体时，这样做会更加高效。
  ```

- 切片浅拷贝和深拷贝

  ```
  浅拷贝：
  源切片和目的切片共享同一底层数组空间，源切片修改，目的切片同样被修改。
  方式：直接赋值（引用底层同一个数组）
  
  深拷贝：
  源切片和目的切片各自都有彼此独立的底层数组空间，复制一份，各自的修改，彼此不受影响
  方式：copy() append()
  ```

- 指针之间的比较
  指针之间也是可以进行相等测试的，只有当它们指向同一个变量或全部是nil时才相等。

  ```
  var x, y int
  fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
  ```

  ```
  var p = f()
  
  func f() *int {
      v := 1
      return &v
  }
  
  fmt.Println(f() == f()) // "false"
  ```

  ```
  func incr(p *int) int {
      *p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
      return *p
  }
  
  v := 1
  incr(&v)              // side effect: v is now 2
  fmt.Println(incr(&v)) // "3" (and v is 3)
  ```

- 变量垃圾回收&内存逃逸

  ```
  https://docs.hacknode.org/gopl-zh/ch2/ch2-03.html
  ```

- 赋值
  如果map查找、类型断言或通道接收出现在赋值语句的右边，它们都可能会产生两个结果，有一个额外的布尔结果表示操作是否成功：

  ```
  v, ok = m[key]             // map lookup
  v, ok = x.(T)              // type assertion
  v, ok = <-ch               // channel receive
  ```

- 类型转换

  对于每一个类型T，都有一个对应的类型转换操作T(x)，用于将x转为T类型（译注：如果T是指针类型，可能会需要用小括弧包装T，比如`(*int)(0)`）。只有当两个类型的底层基础类型相同时，才允许这种转型操作，或者是两者都是指向相同底层结构的指针类型，这些转换只改变类型而不会影响值本身。

  ```go
  type A int
  type B int
  type C float32
  type D A
  func main(){
  	var i B = 1
  	var b *B = &i
  	var a *A = (*A)(b)
  	var d *D = (*D)(b)
  	fmt.Printf("%v,type:%T\n", a , a)
  	fmt.Printf("%v,type:%T\n", d , d)
  }
  ```

- 类型之间的比较

  比较运算符`==`和`<`也可以用来比较一个命名类型的变量和另一个有相同类型的变量，或有着相同底层类型的未命名类型的值之间做比较。但是如果两个值有着不同的类型，则不能直接进行比较：

  ```go
  type Celsius int    // 摄氏温度
  type Fahrenheit int // 华氏温度
  func main(){
  	var c Celsius
  	var f Fahrenheit
  	fmt.Println(c == 0)          // "true"
  	fmt.Println(f >= 0)          // "true"
  	//fmt.Println(c == f)          // compile error: type mismatch
  	fmt.Println(c == Celsius(f)) // "true"!
  }
  ```

- Go语言的字符有以下两种：

  一种是uint8类型,或者叫byte型,代表了ASCII码的一个字符。
  另一种是rune类型,代表一个UTF-8字符。当需要处理中文、日文或者其他复合字符时,
  则需要用到rune类型。rune类型实际是一个int32。

  UTF-8是编码规则,将 Unicode中字符的ID以某种方式进行编码，utf-8是unicode的一种实现。

- 字符串遍历

  ```
  len() :获取字符串ASCLL码长度或者字节长度
  ASCII 字符串长度使用len() 长度
  Unicode 字符串长度使用utf8.RuneCountInString()
  
  1.遍历每一个 ASCII 字符
  直接使用for i:=0;i<len(str) ;i++ {}
  
  2.按 Unicode 字符遍历字符串
  使用 for _, s := range str{}
  ```

- 单引号和双引号

  ```
  而在Go中，双引号是用来表示字符串string，其实质是一个byte类型的数组，单引号表示rune类型。还有一个反引号，用来创建原生的字符串字面量，它可以由多行组成，但不支持任何转义序列。因此，当把两个不同类型的变量进行拼接时，就会报错。
  ```

- 数组

  如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算。

  ```
  q := [...]int{1, 2, 3}
  fmt.Printf("%T\n", q) // "[3]int"
  ```

  数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。

  如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的，这时候我们可以直接通过==比较运算符来比较两个数组，只有当两个数组的所有元素都是相等的时候数组才是相等的。不相等比较运算符!=遵循同样的规则。

  ```go
  a := [2]int{1, 2}
  b := [...]int{1, 2}
  c := [2]int{1, 3}
  fmt.Println(a == b, a == c, b == c) // "true false false"
  d := [3]int{1, 2}
  fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
  ```

- 函数参数传递

  ```go
  当调用一个函数的时候，函数的每个调用参数将会被赋值给函数内部的参数变量，所以函数参数变量接收的是一个复制的副本，并不是原始调用的变量。因为函数参数传递的机制导致传递大的数组类型将是低效的，并且对数组参数的任何的修改都是发生在复制的数组上，并不能直接修改调用时原始的数组变量。在这个方面，Go语言对待数组的方式和其它很多编程语言不同，其它编程语言可能会隐式地将数组作为引用或指针对象传入被调用的函数。
  ```

- slice

  ```go
  # 指针的结构：
  一个slice由三个部分构成：指针、长度和容量。指针指向第一个slice元素对应的底层数组元素的地址，要注意的是slice的第一个元素并不一定就是数组的第一个元素。长度对应slice中元素的数目；长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。内置的len和cap函数分别返回slice的长度和容量。
  
  # 多个slice之间可以共享底层的数据，并且引用的数组部分区间可能重叠。
  
  # 和数组不同的是，slice之间不能比较，因此我们不能使用==操作符来判断两个slice是否含有全部相等元素。不过标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相等（[]byte），但是对于其他类型的slice，我们必须自己展开每个元素进行比较.
  
  # 为啥slice不支持比较：
  第一个原因，一个slice的元素是间接引用的，一个slice甚至可以包含自身。虽然有很多办法处理这种情形，但是没有一个是简单有效的。
  第二个原因，因为slice的元素是间接引用的，一个固定的slice值(译注：指slice本身的值，不是元素的值)在不同的时刻可能包含不同的元素，因为底层数组的元素可能会被修改。
  
  # 一个零值的slice等于nil.一个nil值的slice并没有底层数组。一个nil值的slice的长度和容量都是0，但是也有非nil值的slice的长度和容量也是0的.
  var s []int    // len(s) == 0, s == nil
  s = nil        // len(s) == 0, s == nil
  s = []int(nil) // len(s) == 0, s == nil
  
  s = []int{}    // len(s) == 0, s != nil
  c:=make([]int,0)
  fmt.Println(c==nil)  //false
  如果你需要测试一个slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断。
  
  # 初始化切片
  make([]T, len)
  make([]T, len, cap) // same as make([]T, cap)[:len]
  分配好一个刚刚合适大小的切片：names := make([]string, 0, len(data))
  ```

- map

  ```go
  初始化：
  make(map[string]string)
  make(map[string]string,length)
  
  删除一个元素
  delete(map,key)
  
  map禁止取址，原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。
  _ = &ages["bob"] // compile error: cannot take address of map element
  
  map遍历：for k,v := range map{},遍历的顺序是随机的
  
  和slice一样，map之间也不能进行相等比较；唯一的例外是和nil进行比较。
  
  map类型的零值是nil，也就是没有引用任何哈希表。
  var ages map[string]int
  fmt.Println(ages == nil)    // "true"
  fmt.Println(len(ages) == 0) // "true"
  
  b:=map[int]int{} 						
  fmt.Println(b==nil) 				//false
  c:=make(map[int]int)
  fmt.Println(c==nil) 				//false
  但是向一个nil值的map存入元素将导致一个panic异常，在向map存数据前必须先make（）一个map。
  
  map的key必须是可比较的类型
  
  使用 cap() 获取 map 的容量。
  1.使用 make 创建 map 变量时可以指定第二个参数，不过会被忽略。
  2.cap() 函数适用于数组、数组指针、slice 和 channel，不适用于 map，可以使用 len() 返回 map 的元素个数。
  ```

- 结构体

  ```go
  如果考虑效率的话，较大的结构体通常会用指针的方式传入和返回。
  
  如果要在函数内部修改结构体成员的话，用指针传入是必须的；因为在Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。
  
  如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的。
  类型比较参考：https://juejin.im/post/6844903923166232589
  
  Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员。匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。匿名嵌入的特性，我们可以直接访问叶子属性而不需要给出完整的路径。
  type Point struct {
      X, Y int
  }
  type Circle struct {
      Point
      Radius int
  }
  ```

- 函数

  ```go
  实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实参包括引用类型，如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用被修改。
  
  可变参数：
  在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号“...”，这表示该函数会接收任意数量的该类型参数。
  定义：
  func sum(vals...int) int {
      total := 0
      for _, val := range vals {
          total += val
      }
      return total
  }
  调用方式：
  fmt.Println(sum())           // "0"
  fmt.Println(sum(3))          // "3"
  fmt.Println(sum(1, 2, 3, 4)) // "10"
  values := []int{1, 2, 3, 4}
  fmt.Println(sum(values...)) // "10"
  
  ```
  
- defer && recover

  为了方便诊断问题，runtime包允许程序员输出堆栈信息。

  ```go
  package main
  
  import (
  	"fmt"
  	"os"
  	"runtime"
  )
  
  type S struct {
  	name string
  }
  
  func f(x int) {
  	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
  	defer fmt.Printf("defer %d\n", x)
  	f(x - 1)
  }
  
  func main() {
  	defer printStack()
  	f(3)
  }
  func printStack() {
  	var buf [4096]byte
  	n := runtime.Stack(buf[:], false)
  	// 输出错误日志
  	os.Stdout.Write(buf[:n])  // fmt.Println(string(buf[:n]))
  }
  ```

- 方法的接收器 指针和非指针的区别

  ```
  1.不管你的method的receiver是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换。
  2.在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的因素，第一方面是这个对象本身是不是特别大，如果声明为非指针变量时，调用一次method会产生一次值拷贝；第二方面是如果你用指针类型作为receiver，那么你一定要注意，这种指针类型指向的始终是一块内存地址。
  ```

- 类型断言

  ```go
  data := x.(Type)
  
  data,ok := x.(Type)
  
  switch x.(Type):
  case nil:
  	//
  case int,uint:
  	//
  case string:
  	//
  default:
  	//
  ```

- chan

  ```go
  # 创建
  ch := make(chan Type)
  
  # 和其它的引用类型一样，channel的零值也是nil。
  
  # chan 基本操作
  ch <- x  // a send statement
  x = <-ch // a receive expression in an assignment statement
  <-ch     // a receive statement; result is discarded
  
  close(ch)
  Channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常。对一个已经被close过的channel进行接收操作依然可以接受到之前已经成功发送的数据；如果channel中已经没有数据的话将产生一个零值的数据。
  所有的数据已经全部发送时才需要关闭channel。
  
  # 不带缓存的Channels
  基于无缓存Channels的发送和接收操作将导致两个goroutine做一次同步操作。因为这个原因，无缓存Channels有时候也被称为同步Channels。一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接收操作，当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句。
  
  # 单方向的channel类型
  类型chan<- int表示一个只发送int的channel。相反，类型<-chan int表示一个只接收int的channel。
  对一个只接收的channel调用close将编译错误。
  
  # 带缓存的Channels
  创建 ch = make(chan string, 3)
  获取容量 cap(ch)
  获取长度 len(ch)
  
  # 导致程序死锁
  报错：fatal error: all goroutines are asleep - deadlock!
  原因：Channel满了，就阻塞写，如果Channel空了，就阻塞读
  
  # for...range 遍历所有的chan
  func main() {
  	ch := make(chan int, 100)
  	for i:=0;i<100;i++{
  		i := i
  		go func(){
  			ch <- i
  		}()
  	}
  	for c := range ch{
  		time.Sleep(time.Second)
  		fmt.Println(c)
  	}
  }
  
  # select 语句选择一组可能的send操作和receive操作去处理
  case 可以是 send 语句，也可以是receive语句，亦或者default
  package main
  import "fmt"
  func fibonacci(c, quit chan int) {
  	x, y := 0, 1
  	for {
  		select {
  			case c <- x:
  				x, y = y, x+y
  			case <-quit:
  				fmt.Println("quit")
  				return
  		}
  	}
  }
  func main() {
  	c := make(chan int)
  	quit := make(chan int)
  	go func() {
  		for i := 0; i < 10; i++ {
  			fmt.Println(<-c)
  		}
  		quit <- 0
  	}()
  	fibonacci(c, quit)
  }
  
  # time.After 和 chan 处理 timeout
  package main
  import (
  	"fmt"
  	"time"
  )
  func main() {
  	c1 := make(chan string, 1)
  	go func() {
  		time.Sleep(time.Second * 2)
  		c1 <- "result 1"
  	}()
  	select {
  		case res := <-c1:
  			fmt.Println(res)
  		case <-time.After(time.Second * 1):
  			fmt.Println("timeout 1")
  	}
  }
  
  # timer 和 ticker
  timer 多久后执行
  ticker 每隔多久去执行
  类似timer, ticker也可以通过Stop方法来停止。一旦它停止，接收者不再会从channel中接收数据了。
  
  # goroutine 退出
  （1）main退出
   (2)通过channel通知退出
    for {
  		select {
  		case <-quit:
  			fmt.Println("cancel goroutine by channel!")
  			return
  		default:
  			_, err := io.WriteString(c, "hello cancelByChannel")
  			if err != nil {
  				return
  			}
  			time.Sleep(1 * time.Second)
  		}
  	}
  (3)panic 引起退出
  (4) context
  ```

  

- Interface 是可比较的

  ```
  关于interface的概念：
  
  1.interface 包含两个字段 type 和data，要判断interface相等需要 type 和 data 都相等。
  2.interface 的零值为 nil
  3.生成 interface 有两种情况： interface 赋值给 interface， 其他对象赋值给 interface。
  
  interface 赋值给 interface：
  直接用原变量的 type 字段生成新变量的 type 字段， 原变量的 data 赋值给新变量的 data
  
  其他类型赋值给 interface:
  使用原变量的类型赋值给 interface 的 type 字段，原变量的值赋值给 interface 的 data 字段。
  
  
  ```

  