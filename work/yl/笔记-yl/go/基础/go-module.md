- 常用命令

  ```
  命令	作用
  go mod init	生成 go.mod 文件
  go mod download	下载 go.mod 文件中指明的所有依赖
  go mod tidy	整理现有的依赖
  go mod graph	查看现有的依赖结构(每个包依赖哪些包的关系)
  go mod edit	编辑 go.mod 文件
  go mod vendor	导出项目所有的依赖到vendor目录
  go mod verify	校验一个模块是否被篡改过
  go mod why	查看为什么需要依赖某模块（go mod why -m xx 可以查看为什么要间接依赖某个包）
  ```

  

- 环境变量 GO111MODULE

  ```
  这个环境变量主要是 Go modules 的开关，主要有以下参数：
  - auto：只在项目包含了 go.mod 文件时启用 Go modules，在 Go 1.13 中仍然是默认值，详见 ：golang.org/issue/31857。
  - on：无脑启用 Go modules，推荐设置，未来版本中的默认值，让 GOPATH 从此成为历史。
  - off：禁用 Go modules。
  
  开启GO111MODULE,mac下设置gopath环境变量
  1.Bash
  编辑~/.bash_profile文件，添加以下代码
  export GOROOT=/usr/local/Cellar/go/1.10.3/libexec
  export GOPATH=/Users/chenxingyi/work/go
  export GOBIN=
  export PATH=$PATH:${GOPATH//://bin:}/bin
  保存，然后执行
  source ~/.bash_profile
  
  2.Zsh
  编辑~/.zshrc文件，添加以下代码
  export GOROOT=/usr/local/Cellar/go/1.10.3/libexec
  export GOPATH=/Users/chenxingyi/work/go
  export GOBIN=
  export PATH=$PATH:${GOPATH//://bin:}/bin
  保存，然后执行
  source ~/.zshrc
  ```

- 环境变量 GOPROXY

  ```
  这个环境变量主要是用于设置 Go 模块代理，主要如下：
  
  它的值是一个以英文逗号 “,” 分割的 Go module proxy 列表（稍后讲解）
  作用：用于使 Go 在后续拉取模块版本时能够脱离传统的 VCS 方式从镜像站点快速拉取。它拥有一个默认值，但很可惜 proxy.golang.org 在中国无法访问，故而建议使用 goproxy.cn 作为替代。
  设置为 “off” ：禁止 Go 在后续操作中使用任 何 Go module proxy。
  刚刚在上面，我们可以发现值列表中有 “direct” ，它又有什么作用呢？
  
  其实值列表中的 “direct” 为特殊指示符，用于指示 Go 回源到模块版本的源地址去抓取 (比如 GitHub 等)，当值列表中上一个 Go module proxy 返回 404 或 410 错误时，Go 自动尝试列表中的下一个，遇见 “direct” 时回源，遇见 EOF 时终止并抛出类似 “invalid version: unknown revision...” 的错误。 
  ```

- GONOPROXY/GONOSUMDB/GOPRIVATE

  ```
  这三个环境变量都是用在当前项目依赖了私有模块，也就是依赖了由 GOPROXY 指定的 Go module proxy 或由 GOSUMDB 指定 Go checksum database 无法访问到的模块时的场景，他们具有如下特性：
  
  GOPROXY 是无权访问到任何人的私有模块的，所以你放心，安全性没问题。
  GOPROXY 除了设置模块代理的地址以外，还需要增加 “direct” 特殊标识才可以成功拉取私有库。
  
  它们三个的值都是一个以英文逗号 “,” 分割的模块路径前缀，匹配规则同 path.Match。
  其中 GOPRIVATE 较为特殊，它的值将作为 GONOPROXY 和 GONOSUMDB 的默认值，所以建议的最佳姿势是只是用 GOPRIVATE。
  在使用上来讲，比如 GOPRIVATE=*.corp.example.com 表示所有模块路径以 corp.example.com 的下一级域名 (如 team1.corp.example.com) 为前缀的模块版本都将不经过 Go module proxy 和 Go checksum database，需要注意的是不包括 corp.example.com 本身。Global Caching
  ```

- Global Caching

  ```
  这个主要是针对 Go modules 的全局缓存数据说明，如下：
  
  同一个模块版本的数据只缓存一份，所有其他模块共享使用。
  目前所有模块版本数据均缓存在 $GOPATH/pkg/mod和  $GOPATH/pkg/sum 下，未来或将移至 $GOCACHE/mod和 $GOCACHE/sum 下( 可能会在当 $GOPATH 被淘汰后)。
  可以使用 go clean-modcache 清理所有已缓存的模块版本数据。
  另外在 Go1.11 之后 GOCACHE 已经不允许设置为 off 了，我想着这也是为了模块数据缓存移动位置做准备，因此大家应该尽快做好适配。 
  ```

- 快速迁移项目至 Go Modules

  ```
   1.升级到 Go 1.13。
   2.让 GOPATH 从你的脑海中完全消失，早一步踏入未来。
   修改 GOBIN 路径（可选）。
   打开 Go modules 的开关。
   设置 GOPROXY。
  3.按照你喜欢的目录结构重新组织你的所有项目。
  4.在你项目的根目录下执行 go mod init<OPTIONAL_MODULE_PATH> 以生成 go.mod 文件。
  ```

- 更新现有全部模块

  ```
  go get -u 只会更新主要模块
  go get -u all 更新所有模块
  
  go get golang.org/x/text@latest	拉取最新的版本，若存在tag，则优先使用。
  go get golang.org/x/text@master	拉取 master 分支的最新 commit。
  go get golang.org/x/text@v0.3.2	拉取 tag 为 v0.3.2 的 commit。
  go get golang.org/x/text@342b2e	拉取 hash 为 342b231 的 commit，最终会被转换为 v0.3.2。
  ```

- 查看go.mod文件

  ```
  
  module github.com/eddycjy/module-repo
  
  go 1.13
  
  require (
      example.com/apple v0.1.2
      example.com/banana v1.2.3
      example.com/banana/v2 v2.3.4
      example.com/pear // indirect
      example.com/strawberry // incompatible
  )
  
  exclude example.com/banana v1.2.4
  replace example.com/apple v0.1.2 => example.com/fried v0.1.0 
  replace example.com/banana => example.com/fish
  ```

  ```
  - module：用于定义当前项目的模块路径。
  - go：用于标识当前模块的 Go 语言版本，值为初始化模块时的版本，目前来看还只是个标识作用。
  - require：用于设置一个特定的模块版本。
  - exclude：用于从使用中排除一个特定的模块版本。
  - replace：用于将一个模块版本替换为另外一个模块版本。
  
  另外你会发现 `example.com/pear` 的后面会有一个 indirect 标识，indirect 标识表示该模块为间接依赖，也就是在当前应用程序中的 import 语句中，并没有发现这个模块的明确引用，有可能是你先手动 `go get` 拉取下来的，也有可能是你所依赖的模块所依赖的，情况有好几种。
  ```

