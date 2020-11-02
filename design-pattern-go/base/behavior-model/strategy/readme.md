代码建模
http://tigerb.cn/go-patterns/#/?id=%e4%bb%80%e4%b9%88%e6%98%af%e3%80%8c%e7%ad%96%e7%95%a5%e6%a8%a1%e5%bc%8f%e3%80%8d%ef%bc%9f

「策略模式」的核心是接口：(详见 example1)
例如：
PaymentInterface
Pay(ctx *Context) error 当前支付方式的支付逻辑
Refund(ctx *Context) error 当前支付方式的退款逻辑

参考步骤：(详见 example2)
1.策略的接口
2.策略对象
3.策略的执行者