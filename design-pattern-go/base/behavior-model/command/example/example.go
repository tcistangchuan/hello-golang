package main

import (
	"fmt"
)

// 接收者
type Receiver interface {
	Execute() error
}

// 命令
type Command interface {
	Call() error
}

type ReceiverA struct{}

func (r *ReceiverA) Execute() error {
	fmt.Println("receiver A doing")
	return nil
}

type ReceiverB struct{}

func (r *ReceiverB) Execute() error {
	fmt.Println("receiver B doing")
	return nil
}

type CommonA struct {
	Receiver
}

func (c *CommonA) Call() error {
	return c.Receiver.Execute()
}

type CommonB struct {
	Receiver
}

func (c *CommonB) Call() error {
	return c.Receiver.Execute()
}

// 调用者
type Invoker interface {
	AddCommand(commands ...Command)
	Handler()
}

type Invoker1 struct {
	CommandList []Command
}

func (i *Invoker1) AddCommand(commands ...Command) {
	i.CommandList = append(i.CommandList, commands...)
}
func (i *Invoker1) Handler() {
	for _, v := range i.CommandList {
		_ = v.Call()
	}
}

func main() {
	// 给命令绑定执行的接收者
	c1 := &CommonA{new(ReceiverA)}
	c2 := &CommonB{new(ReceiverB)}

	// 调用者
	i := new(Invoker1)
	// 添加 command
	// i.AddCommand(c1, c2)
	i.AddCommand([]Command{c1, c2}...)
	// 执行命令
	i.Handler()
}
