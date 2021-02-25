package main

import (
	"fmt"
	"testing"
)

type Person struct {
	age int
}

func example02() {
	person := &Person{28}
	fmt.Printf("%p\n", person)
	defer fmt.Println(person.age)
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)
	defer func() {
		fmt.Println(person.age)
	}()
	person = &Person{29}
	fmt.Printf("%p\n", person)
}

func TestExample02(t *testing.T) {
	example02()
}
