```
Go101地址：https://gfw.go101.org/article/container.html	
```

- #### 切片的操作

  ```go
  #复制：
  假设s是被谈到的切片、T是它的元素类型、t0是类型T的零值字面量
  sClone=append(s[:0:0],s...)
  对于长度较大（并且元素值的直接部分不含有指针）的切片，上面这种方法比下面这种方法效率更高。
  sClone := make([]T, len(s))
  copy(sClone, s)
  
  #删除：
  // 第一种方法（保持剩余元素的次序）：
  s = append(s[:from], s[to:]...)
  // 第二种方法（保持剩余元素的次序）：
  s = s[:from + copy(s[from:], s[to:])]
  
  ```

  

- #### 切片: 元素位置上的修改不会改变内存地址，增加超过cap容量地址会改变

  ```go
  func main() {
  	a1:=make([]int,3,10)
  	b1 := a1
  	fmt.Printf("%p \n", a1)
  	fmt.Printf("%p \n", b1)
  	fmt.Println()
  	b1[1] = 22
  	fmt.Printf("%p \n", a1)
  	fmt.Printf("%p \n", b1)
  	fmt.Println()
  	b1 = append(b1, 4)
  	fmt.Printf("%p \n", a1)
  	fmt.Printf("%p \n", b1)
  	fmt.Println()
    //输出：
    //0xc00008c000
    //0xc00008c000
    //
    //0xc00008c000
    //0xc00008c000
    //
    //0xc00008c000
    //0xc00008c000
  }
  
  
  func main(){
  	a1 := []int{1, 2, 3}
  	//a1:=make([]int,3,10)
  	b1 := a1
  	fmt.Printf("%p \n", a1)
  	fmt.Printf("%p \n", b1)
  	fmt.Println()
  	b1[1] = 22
  	fmt.Printf("%p \n", a1)
  	fmt.Printf("%p \n", b1)
  	fmt.Println()
  	b1 = append(b1, 4)
  	fmt.Printf("%p \n", a1)
  	fmt.Printf("%p \n", b1)
  	fmt.Println()
  	
  	//输出
  	//0xc000016180
  	//0xc000016180
  	//
  	//0xc000016180
  	//0xc000016180
  	//
  	//0xc000016180
  	//0xc0000180f0
  } 
  ```

- #### 数组或者切片派生切片（取子切片）

  ```go
  1.派生出来的切片的元素和基础切片（或者数组）的元素位于同一个内存片段上。或者说，派生出来的切片和基础切片（或者数组）将共享一些元素
  例子：
  func main(){
  	a1:=[]int{1,2,3,4,5}
  	fmt.Println(a1)
  	fmt.Printf("%p \n",a1)
  	b1:=a1[0:2]
  	fmt.Println(b1)
  	fmt.Printf("%p \n",b1)
  	//输出
  	//[1 2 3 4 5]
  	//0xc0000180f0
  	//[1 2]
  	//0xc0000180f0
  }
  
  2.Go中有两种取子切片的语法形式
  baseContainer[low : high]       // 双下标形式
  baseContainer[low : high : cap(baseContainer)] // 三下标形式
  // 双下标形式
  0 <= low <= high <= cap(baseContainer)
  // 三下标形式
  0 <= low <= high <= max <= cap(baseContainer)
  
  #如果下标low为零，则它可被省略。此条规则同时适用于双下标形式和三下标形式。
  #如果下标high等于len(baseContainer)，则它可被省略。此条规则同时只适用于双下标形式。
  #三下标形式中的下标max在任何情况下都不可被省略。
  
  ```

  

- #### map :无论内部元素如何改变都不会改变整个map的内存地址

  ```go
  package main
  
  import "fmt"
  
  func main(){
  	a1:=map[int]int{1:1,2:2,3:3}
  	b1:=a1
  	fmt.Printf("%p \n",a1)
  	fmt.Printf("%p \n",b1)
  	fmt.Println()
  	b1[1]=22 //修改
  	fmt.Printf("%p \n",a1)
  	fmt.Printf("%p \n",b1)
  	fmt.Println()
  	delete(b1,1) //删除
  	b1[4]=44 //新增
  	fmt.Printf("%p \n",a1)
  	fmt.Printf("%p \n",b1)
  	fmt.Println(a1)
  	fmt.Println(b1)
  }
  //输出：
  //0xc000066180
  //0xc000066180
  //
  //0xc000066180
  //0xc000066180
  //
  //0xc000066180
  //0xc000066180
  //map[2:2 3:3 4:44]
  //map[2:2 3:3 4:44]
  ```

- #### for ... range 奇怪的现象

  ```go
  #for 和 for range的区别：
  1.for range 是值拷贝，所以不要试图直接去修改value,即使要修改也要用array[i]的方式。
  2.for range 在遍历大数据时，效率不高，需要拷贝数据。
  3.遍历map顺序是随机的。
  4.遍历过程中，一个还没有被遍历到的条目被删除了，则此条目保证不会被遍历出来;一个新的条目被添加入此映射，则此条目并不保证将在此遍历过程中被遍历出来。
  
  # 遍历索引和元素的 range 循环时，Go 语言会额外创建一个新的v变量存储切片中的元素，循环中使用的这个变量 v 会在每一次迭代被重新赋值而覆盖，在赋值时也发生了拷贝。
  func main() {
  	arr := []int{1, 2, 3}
  	newArr := []*int{}
  	for  _,v:= range arr {
  		newArr=append(newArr,&v)
  	}
  	for _,v := range newArr{
  		fmt.Println(*v)
  	}
  	//输出
  	//3
  	//3
  	//3
  }
  
  # 用访问索引的方式
  func main() {
  	arr := []int{1, 2, 3}
  	newArr := []*int{}
  	for  k,_:= range arr {
  		newArr=append(newArr,&arr[k])
  	}
  	for _,v := range newArr{
  		fmt.Println(*v)
  	}
  	//输出
  	//1
  	//2
  	//3
  }
  ```

- #### 把数据指针当数组用

  ```go
  var pa *[5]int // == nil
  fmt.Println(len(pa), cap(pa)) // 5 5
  
  例子1:
  func main() {
  	var p [5]int // nil
  	fmt.Println(p)
  
  	for i,_ := range p { // okay
  		fmt.Println(i)
  	}
  
  	for i, n := range p { // panic
  		fmt.Println(i, n)
  	}
  	//[0 0 0 0 0]
  	//
  	//0
  	//1
  	//2
  	//3
  	//4
  	//
  	//0 0
  	//1 0
  	//2 0
  	//3 0
  	//4 0
  }
  
  例子2:
  func main() {
  	var p *[5]int // nil
  	fmt.Println(p)
  
  	for i,_ := range p { // okay
  		fmt.Println(i)
  	}
  
  	//for i, n := range p { // panic
  	//	fmt.Println(i, n)
  	//}
  	//[0 0 0 0 0]
  	//
  	//0
  	//1
  	//2
  	//3
  	//4
  }
  ```

- 