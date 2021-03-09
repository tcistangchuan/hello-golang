- GMP原理

  参考链接：https://learnku.com/articles/41728
  
  ```
  G: goroutine对象,包含了goroutine的栈，上下文，状态等信息。
M: 指的是线程，goroutine是在线程上执行的。
  P: 指的是处理器，又叫协程调度器。每一个M都要绑定一个P,由P来调度本地的goroutine队列中的G到绑定的M上去运行。
  
  抢占机调度证了不会有一个G长时间的运行导致其他G无法运行的情况发生。一个g运行到10ms,p就会解除g和m，将g放回全局g队列里。
  
  ```
  
  

