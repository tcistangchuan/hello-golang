

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
  1.用互斥锁sync.Mutex //能控制并发对共享量的读写
  func main() {
    var mu sync.Mutex
  	mu.Lock()
  	go func(){
  		time.Sleep(time.Second)
  		fmt.Println("coding")
  		defer mu.Unlock()
  
  	}()
  	fmt.Println("start")
  	mu.Lock()
  	fmt.Println("end")
  	defer  mu.Unlock()
  }
  
  2.通过channel生产者消费者模型实现并发控制
  func main() {
  	ch := make(chan int)
  	go func(){
  		fmt.Println("生产者")
  		ch<- 1
  	}()
  
  	fmt.Println("消费者")
  	time.Sleep(2*time.Second)
  
  }
  
  3.通过sync包中的WaitGroup实现并发控制
  
  4.contex 能结束掉超时的协程
  
  ```
  
  

