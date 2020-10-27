package main

import "fmt"

type Context struct {
}

type Handler interface {
	Do(ctx Context) error
	SetNextHandler(h Handler) Handler
	Run(ctx Context) error
}

type Next struct {
	NextHandler Handler
}

func (n *Next) SetNextHandler(h Handler) Handler {
	n.NextHandler = h
	return h
}

func (n *Next) Run(ctx Context) (err error) {
	if n.NextHandler != nil {
		// 执行下一个handler的do
		if err = n.NextHandler.Do(ctx); err != nil {
			return
		}
		return (n.NextHandler).Run(ctx)
	}
	return
}

// 实例1
type Example1 struct {
	// 这里不能用 *Next 会导致nil指针panic
	Next
}

func (e1 Example1) Do(ctx Context) error {
	fmt.Println("执行实例一步骤")
	return nil
}

// 实例2
type Example2 struct {
	Next
}

func (e2 Example2) Do(ctx Context) error {
	fmt.Println("执行实例二步骤")
	return nil
}

// 实例3
type Example3 struct {
	Next
}

func (e3 Example3) Do(ctx Context) error {
	fmt.Println("执行实例三步骤")
	return nil
}

// 啥都不干 只做载体
type NullHandler struct {
	Next
}

func (n *NullHandler) Do(ctx Context) error {
	return nil
}


func main() {
	var ctx Context
	h := new(NullHandler)
	h.SetNextHandler(new(Example1)).SetNextHandler(new(Example2)).SetNextHandler(new(Example3))
	if err := h.Run(ctx); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("ok")


}
