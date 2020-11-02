package main

import "fmt"

type scales interface {
	// 获取当前的重量
	getWeight()int
}

// 人 基础类
type person struct {
	weight int
}
func (p *person)getWeight()int {
	return p.weight
}

// 内衣 装饰器类
type underwear struct {
	weight int
	other scales
}
func (p *underwear)getWeight()int {
	return p.weight+p.other.getWeight()
}

// 外套 装饰器类
type coat struct {
	weight int
	other scales
}
func (p *coat)getWeight()int {
	return p.weight+p.other.getWeight()
}

func main(){
	// 人
	p := &person{
		weight: 100,
	}
	fmt.Println("人净重：",p.getWeight())

	// +内衣
	p2 := &underwear{
		weight: 10,
		other: p,
	}
	fmt.Println("加上内衣净重：",p2.getWeight())

	// +外套
	p3 := &coat{
		weight: 50,
		other:  p2,
	}
	fmt.Println("加上外套净重：",p3.getWeight())

}