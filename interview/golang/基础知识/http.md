- Get和Post区别

  ```
  Get是不安全的，因为在传输过程，数据被放在请求的URL中；。
  Get传送的数据量较小，这主要是因为受URL长度限制；Post传送的数据量较大，一般被默认为不受限制。
  Get执行效率却比Post方法好。Get是form提交的默认方法。
  ```
  
- Http请求的完全过程

  ```
  1.浏览器根据域名解析IP地址（DNS）,并查DNS缓存
  2.浏览器与WEB服务器建立一个TCP连接
  3.浏览器给WEB服务器发送一个HTTP请求（GET/POST）：一个HTTP请求报文由 请求行（request line）、请求头部（headers）、空行（blank line）和请求数据（request body）4个部分组成。
  4.服务端响应HTTP响应报文，报文由 状态行（status line）、响应头部（headers）、空行（blank line）和响应数据（response body）4个部分组成。
  5.浏览器解析渲染
  ```

  

- 三次握手四次挥手

  参考：https://blog.csdn.net/qzcsu/article/details/72861891

- TCP三次握手：

![image-20210303194019740](/Users/chuantang/go/src/study/go/interview/golang/基础知识/三次握手.png)

- TCP连接的释放（四次挥手）

  ![image-20210303203506351](/Users/chuantang/go/src/study/go/interview/golang/基础知识/四次挥手.png)

- 三次握手和四次挥手高频问题

  https://juejin.cn/post/6844903958624878606#heading-2



- http和https的区别,http2的优点

  https://juejin.cn/post/6844903489596833800