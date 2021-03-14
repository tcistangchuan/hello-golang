package main

import "fmt"

func printAll(vals interface{}) {//1
	fmt.Println(vals)
}

func main() {
	names := []string{"stanley", "david", "oscar"}
	printAll(names)
}
