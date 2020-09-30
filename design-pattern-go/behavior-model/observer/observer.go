package main

import "fmt"

/**
观察者模式
观察者模式有时也被称作发布/订阅模式，该模式用于为对象实现发布/订阅功能：
一旦主体对象状态发生改变，与之关联的观察者对象会收到通知，并进行相应操作。
角色
抽象观察者
具体观察者

抽象被观察者/对象 [可选]
具体被观察者/对象
*/

/**
被观察者对象
*/
type subject struct {
	context   string
	observers []*observer
}

func newSub(ctx string) *subject {
	return &subject{
		context: ctx,
	}
}
func (s *subject) attach(o *observer) {
	s.observers = append(s.observers, o)
}
func (s *subject) notify() {
	for _, k := range s.observers {
		k.update(s)
	}
}
func (s *subject) updateCtx(ctx string) {
	s.context = ctx
	s.notify()
}

/**
观察者接口
*/
type observerInterface interface {
	update(*subject)
}

/**
具体观察者
*/
type observer struct {
	name string
}

func newOb(str string) *observer {
	return &observer{
		name: str,
	}
}
func (o *observer) update(s *subject) {
	fmt.Printf("%s receive %s\n", o.name, s.context)
}

func main() {
	s := newSub("去吃饭了")
	s.attach(newOb("张三"))
	s.attach(newOb("李四"))
	s.attach(newOb("王二麻子"))
	s.notify()
	fmt.Println("")
	s.updateCtx("临时有事，先开会")
}
