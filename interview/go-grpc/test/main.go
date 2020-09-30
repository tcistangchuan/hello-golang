package main

import (
	"fmt"
	"study/go/interview/go-grpc/test/ctx"
)

func main() {
	a := ctx.Tc{Name: "tang"}
	fmt.Println(a)
	fmt.Println(ctx.Log)
	ctx.SetLog(11)
	fmt.Println(ctx.Log)
}
