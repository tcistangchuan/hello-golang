

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
  1.用互斥锁让多个协程可以串行执行
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
  
  2.通过channel实现的广播模式，实现对多个协程的通知，比如通知结束。
  
  
  3.通过sync包中的WaitGroup实现同时执行多个协程，当最慢的协程执行完，才结束。
  	开启多个协程同时处理请求，全部请求之后返回主函数。
  
  4.多个协程执行，其中一个协程执行完就返回。用带缓冲的channel
  
  ```
  
  

