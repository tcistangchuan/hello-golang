package main

import (
	"errors"
	"fmt"
)

type Context struct {

}
type Handler interface {
	Do(c Context)error
	SetNextHandler(h Handler)Handler
	Run(c Context)error
}

type Next struct {
	NextHandler Handler
}
func(n *Next)SetNextHandler(h Handler)Handler{
	n.NextHandler = h
	return h
}
func (n *Next)Run(c Context)error{
	if n.NextHandler != nil {
		if err := n.NextHandler.Do(c);err!=nil{
			fmt.Println("err:", err)
			return nil
		}
		return n.NextHandler.Run(c)
	}
	fmt.Println("err2")
	return nil
}

type Example1 struct {
	Next
}
func(e Example1)Do(c Context)error{
	fmt.Println("执行example1")
	return nil
}

type Example2 struct {
	Next
}
func(e *Example2)Do(c Context)error{
	fmt.Println("执行example2")
	return errors.New("error")
}

type Example3 struct {
	Next
}
func(e *Example3)Do(c Context)error{
	fmt.Println("执行example3")
	return nil
}

type NullHandler struct {
	Next
}
func (n *NullHandler)Do(c Context)error{
	return nil
}

func main(){
	c := Context{}
	n := &NullHandler{}
	_ =n.SetNextHandler(new(Example1)).SetNextHandler(new(Example2)).SetNextHandler(new(Example3))
	if err := n.Run(c);err!=nil{
		fmt.Println(err)
		return
	}
	return
}