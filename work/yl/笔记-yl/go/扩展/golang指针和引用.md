#指针和引用：

参考地址：https://blog.csdn.net/hel12he/article/details/80590519

​					https://juejin.im/post/5d8f05b55188254a371f1a77







- #### 变量名是程序员给地址取的外号

![image-20200527111255917](/Users/chuantang/Library/Application Support/typora-user-images/image-20200527111255917.png)

```
x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2          // equivalent to x = 2
fmt.Println(x)  // "2"
```

- #### 指针变量

  其实就是这个变量里边存放的是一个变量的地址，通过这个地址，机器可以找到对应变量的值，例如：int * pa = &a，就表示变量 **pa** 抽屉里放的是 **a** 的地址，它的类型是：int*，继续看图：

  ![image-20200527111502660](/Users/chuantang/Library/Application Support/typora-user-images/image-20200527111502660.png)

**这里需要重要说明的是：指针pa与a的关系是：a抽屉里边放的是变量值10，pa放的是变量的地址：0x00000001**



- #### 引用是变量的别名，那么它的地址应该与变量一致



# 小结

- 变量由三分部分构成：变量名、变量值、变量地址；
- 变量名实际上只是给程序员看的，编译后的代码中并不存在变量名；
- 指针变量就是一个变量存储了另外一个变量的地址，系统也会为他分配内存空间来存储这个地址；
- 引用实际是变量的别名，他跟变量有相同的地址。