参考文档地址：https://blog.csdn.net/ffzhihua/article/details/84783802



- 下载golint

  ```shell
  mkdir -p $GOPATH/src/golang.org/x/
  cd $GOPATH/src/golang.org/x/
  git clone https://github.com/golang/lint.git
  git clone https://github.com/golang/tools.git
  
  到目录$GOPATH/src/golang.org/x/lint/golint中运行
  go install
  
  ```

- 打开setting对话框

  Tools->External Tools

  配置参数例子：

  Program   /Users/chuantang/go/bin/golint （golint执行程序的绝对地址）
  Arguments   $FilePath$										（用这个例子里的就行）
  Working directory   /Users/chuantang/go/bin （配置了第一个参数会自动生成，$ProjectFileDir$）		