package main

import "fmt"

type Context struct{}

// component 组件接口
type Component interface {
	// 添加组件
	Mount(c ...Component) error
	// 删除组件
	Remove(c Component) error
	// 执行组件&子组件
	Do(c *Context) error
}

// BaseComponent 基础组件
type BaseComponent struct {
	ChildComponents []Component
}

// 添加组件
func (b *BaseComponent) Mount(c ...Component) (err error) {
	b.ChildComponents = append(b.ChildComponents, c...)
	return
}

// 删除组件
func (b *BaseComponent) Remove(c Component) (err error) {
	if len(b.ChildComponents) == 0 {
		return
	}
	for k, v := range b.ChildComponents {
		if v == c {
			b.ChildComponents = append(b.ChildComponents[:k], b.ChildComponents[k+1:]...)
			return
		}
	}
	return
}

// AllChildDo 执行所有子组件的do
func (b *BaseComponent) AllChildDo(c *Context) (err error) {
	for _, v := range b.ChildComponents {
		if err = v.Do(c); err != nil {
			return
		}
	}
	return
}

type Root struct {
	BaseComponent
}

func NewRoot() *Root {
	return &Root{}
}
func (r *Root) Do(c *Context) (err error) {
	fmt.Println("root do...")
	r.AllChildDo(c)
	return
}

type Son1 struct {
	BaseComponent
}

func NewSon1() *Son1 {
	return &Son1{}
}
func (r *Son1) Do(c *Context) (err error) {
	fmt.Println("Son1 do...")
	r.AllChildDo(c)
	return
}

type Son2 struct {
	BaseComponent
}

func NewSon2() *Son2 {
	return &Son2{}
}
func (r *Son2) Do(c *Context) (err error) {
	fmt.Println("Son2 do...")
	r.AllChildDo(c)
	return
}

type Son3 struct {
	BaseComponent
}

func NewSon3() *Son3 {
	return &Son3{}
}
func (r *Son3) Do(c *Context) (err error) {
	fmt.Println("Son3 do...")
	r.AllChildDo(c)
	return
}

type GrandSon1 struct {
	BaseComponent
}

func NewGrandSon1() *GrandSon1 {
	return &GrandSon1{}
}
func (r *GrandSon1) Do(c *Context) (err error) {
	fmt.Println("GrandSon1 do...")
	r.AllChildDo(c)
	return
}

type GrandSon2 struct {
	BaseComponent
}

func NewGrandSon2() *GrandSon2 {
	return &GrandSon2{}
}
func (r *GrandSon2) Do(c *Context) (err error) {
	fmt.Println("GrandSon2 do...")
	r.AllChildDo(c)
	return
}

func main() {
	r := NewRoot()

	s1 := NewSon1()
	s2 := NewSon2()
	s3 := NewSon3()

	gs1 := NewGrandSon1()
	gs2 := NewGrandSon2()

	s1.Mount(gs1, gs2)
	s1.Remove(gs1)

	r.Mount(s1, s2, s3)

	r.Do(new(Context))
}
