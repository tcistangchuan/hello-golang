package design_pattern_go

/*
	1.
	责任链模式：将多个对象串成一条链路，将请求透串下去，直到每个对象都处理完成，或者其中一个终止。
	实现方式：链表，数组。
	应用：gin框架的context.Next()方法

	type Context struct {
    // ...

    // handlers 是一个包含执行函数的数组
    // type HandlersChain []HandlerFunc
	handlers HandlersChain
    // index 表示当前执行到哪个位置了
	index    int8

    // ...
	}

	// Next 会按照顺序将一个个中间件执行完毕
	// 并且 Next 也可以在中间件中进行调用，达到请求前以及请求后的处理
	// Next should be used only inside middleware.
	// It executes the pending handlers in the chain inside the calling handler.
	// See example in GitHub.
	func (c *Context) Next() {
		c.index++
		for c.index < int8(len(c.handlers)) {
			c.handlers[c.index](c)
			c.index++
		}
	}

	2.
*/
