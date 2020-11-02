package main

import "fmt"

type context struct {

}

// 策略接口
type strategy interface {
	Handler()error
}

// 具体的策略
type strategyA struct {

}

func (s strategyA) Handler() error {
	fmt.Println("do strategyA")
	return nil
}

type strategyB struct {

}

func (s strategyB) Handler() error {
	fmt.Println("do strategyB")
	return nil
}

// 策略的执行者
type client struct {
	strategy strategy
}
func (c *client) set(s strategy)*client{
	c.strategy = s
	return c
}
func (c *client) Do()error{
	_ = c.strategy.Handler()
	return nil
}

func main(){
	s1 := strategyA{}
	s2 := strategyB{}

	c := client{}
	_ = c.set(s1).Do()
	_ = c.set(s2).Do()
}



