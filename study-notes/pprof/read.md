参考链接：
https://www.liwenzhou.com/posts/Go/performance_optimisation/
https://juejin.im/post/6844903992720359432

# Go语言内置了获取程序的运行数据的工具，包括以下两个标准库：
(1)runtime/pprof：采集工具型应用运行数据进行分析
(2)net/http/pprof：采集服务型应用运行时数据进行分析
pprof开启后，每隔一段时间（10ms）就会收集下当前的堆栈信息，
获取各各函数占用的CPU以及内存资源；最后通过对这些采样数据进行分析，
形成一个性能分析报告。
注意，我们只应该在性能测试的时候才在代码中引入pprof。配合压测工具测性能。


# http生成报告：
1.开启pprof 引入 net/hhtp/pprof
2.可以通过命令收集cpu: go tool pprof http://localhost:6060/debug/pprof/profile
3.可以通过命令收集内存: go tool pprof   http://localhost:6060/debug/pprof/heap

# 或者runtime 生成报告
go tool pprof cpu.pprof


# 分析报告：
1.进入交互式模式之后，比较常用的有 top、list、web 等命令
top: 返回字段
    flat
    本函数的执行耗时

    flat%
    flat 占 CPU 总时间的比例。程序总耗时 16.22s, Eat 的 16.19s 占了 99.82%
    
    sum%
    前面每一行的 flat 占比总和
    
    cum
    累计量。指该函数加上该函数调用的函数总耗时
    
    cum%
    cum 占 CPU 总时间的比例

list: list + 函数名 ：直接定位到了相关长耗时的代码处：

web: 直接输入web，通过svg图的方式查看程序中详细的CPU占用情况。
想要查看图形化的界面首先需要安装graphviz图形化工具。
brew install graphviz

每个方框由两个标签组成：在 cpu profile 中，一个是方法运行的时间占比，
一个是它在采样的堆栈中出现的时间占比（前者是 flat 时间，后者则是 cumulate 时间占比)；
框越大，代表耗时越多或是内存分配越多。

quit: 离开交互式界面
