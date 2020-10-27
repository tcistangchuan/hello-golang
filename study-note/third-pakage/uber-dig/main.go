package main

import "fmt"

type Config struct {
	Prefix string
}

func main() {
	a := 1
	switch a {
	case 1:
		fallthrough
	case 2:
		fmt.Println("xx:", a)
		fallthrough
	default:
		fmt.Println("default")

	}
}
