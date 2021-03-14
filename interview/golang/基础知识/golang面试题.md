- 限制协程执行数量的最简单方法

  ```
  参考：https://blog.csdn.net/xuduorui/article/details/53898180
  
    用带缓冲的channel。协程开始的时候往channel中写一个数据，结束时读一个数据，当channel中数据数量超过缓存数，就会阻塞。
  ```

- 如何控制函数的超时

  ```
  参考:https://cloud.tencent.com/developer/article/1603998
  
    1.用context.WithTimeout()和select。WithTimeout方法第二个参数可以设置超时时间，如果超时会自动触发cancel()函数 ,Done通道负责监听context什么时候完成。
    2.channel+select+time.After。先设置一个channel标示，开启一个协程，处理完任务就往channel中写数据，用select在主函数里面监听channel的读操作，同时监听time.After的读操作。
    3.
  ```

- golang中的自增操作n++是原子操作吗

  ```
  实际上 n++ 操作就是CPU先从内存取回n的值，然后加1，再放回内存，显然是存在并发问题的。
  比如：
  func TestNpp(t *testing.T) {
  	n := 0
  	for i := 0; i < 20; i++ {
  		go func() {
  			n++
  		}()
  	}
  
  	fmt.Println(n)
  }
  结果是随机的
  ```

- 除了锁可以控制同步外，还有什么方法？

  ```
  可以channel实现。
  ```

- channel的实现

  ```
  Channel 是一个用于同步和通信的有锁队列，使用互斥锁解决程序中可能存在的线程竞争问题是很常见的，我们能很容易地实现有锁队列。
  Channel 分成了以下三种类型：
    同步 Channel — 不需要缓冲区，发送方会直接将数据交给（Handoff）接收方；
    异步 Channel — 基于环形缓存的传统生产者消费者模型；
    chan struct{} 类型的异步 Channel — struct{} 类型不占用内存空间，不需要实现缓冲区和直接发送（Handoff）的语义；
  ```

  