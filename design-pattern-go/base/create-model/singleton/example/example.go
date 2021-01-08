package main

import (
	"fmt"
	"stuty-go/design-pattern-go/base/create-model/singleton"
)

type T struct {
	name string
}
func main(){
	//a := &T{}
	//a.name ="123"
	//fmt.Println(a)
	a := singleton.GetInstance2()
	(*a).Age =12
	(*a).Name ="xx"
	fmt.Printf("%p", a)
}
