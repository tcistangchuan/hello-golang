

- ```
  Go 通过 channel 实现 CSP 通信模型，主要用于 goroutine 之间的消息传递和事件通知。
  ```

  
  
- 并行和并发

  ```
  并行是指多个时间同一时间之行，并发是在短时间间隔执行，实际还是串行。
  
  并发编程是指在一台处理器上“同时”处理多个任务。并发是在同一实体上的多个事件。多个事件在同一时间间隔发生。并发编程的目标是充分的利用处理器的每一个核，以达到最高的处理性能。
  ```

- Golang 中常用的并发控制模型有三种:

   ```
  1.用互斥锁sync.Mutex
  func main() {
      var mu sync.Mutex
  
      go func(){
          fmt.Println("你好, 世界")
          mu.Lock()
      }()
  
      mu.Unlock()
  }
  
  2.通过channel生产者消费者模型实现并发控制
  // 生产者: 生成 factor 整数倍的序列
  func Producer(factor int, out chan<- int) {
      for i := 0; ; i++ {
          out <- i*factor
      }
  }
  
  // 消费者
  func Consumer(in <-chan int) {
      for v := range in {
          fmt.Println(v)
      }
  }
  func main() {
      ch := make(chan int, 64) // 成果队列
  
      go Producer(3, ch) // 生成 3 的倍数的序列
      go Producer(5, ch) // 生成 5 的倍数的序列
      go Consumer(ch)    // 消费 生成的队列
  
      // 运行一定时间后退出
      time.Sleep(5 * time.Second)
  }
  
  3.通过sync包中的WaitGroup实现并发控制
  
  ```

  

