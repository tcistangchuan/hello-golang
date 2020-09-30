package main

import "fmt"

/*
装饰器/装饰者模式
动态地将责任附加到对象上。想要扩展功能，装饰者提供有别于继承的另一种选择。

装饰者模式 符合开放 - 关闭原则。继承 虽然属于扩展的形式之一，但是在设计中，
推崇多用组合少用继承的说法。但是好像 GO 里面没有继承的说法，只有组合。所以
用 GO 实现装饰者模式还是很方便的。装饰者的主要目的就是对基础对象不断的扩展
功能，虽然在一定程度上保持了扩展性，但是如果过度使用会出现众多小对象，会使程
序变得很复杂。
注意： 不要使用内嵌（继承），使用聚合更加灵活

适用场景：
go-cache 包

角色：
抽象构件（Component）
具体构件（ConcreteComponent）
具体装饰（ConcreteDecorator）

转自链接：https://learnku.com/articles/26820
*/

//  饮料接口
type Beverage interface {
	getDescription() string // 获取饮料描述
	cost() int              //最终价格
}

// 咖啡
type Cafe struct {
	Description string // 描述
	Cost        int    // 价格
}

func (c *Cafe) getDescription() string {
	return c.Description
}
func (c *Cafe) cost() int {
	return c.Cost
}

// 给饮料加的水
type Water struct {
	Beverage  Beverage
	WaterName string
	Cost      int
}

func (w *Water) getDescription() string {
	fmt.Printf("give %s add %s water", w.Beverage.getDescription(), w.WaterName)
	return "add water"
}
func (w *Water) cost() int {
	return w.Beverage.cost() + w.Cost
}

// 给饮料加的糖
type Sugar struct {
	Beverage  Beverage
	SugarName string
	Cost      int
}

func (s *Sugar) getDescription() string {
	fmt.Printf("give %s add %s water", s.Beverage.getDescription(), s.SugarName)
	return "add water"
}
func (s *Sugar) cost() int {
	return s.Beverage.cost() + s.Cost
}

func main() {
	// 饮料
	var b Beverage
	b = &Cafe{
		Description: "雀巢咖啡",
		Cost:        10,
	}
	// 加水
	b = &Water{
		Beverage: b,
		Cost:     1,
	}
	// 加糖
	b = &Sugar{
		Beverage: b,
		Cost:     2,
	}
	// 最终价格
	fmt.Println(b.getDescription(), "最终价格：", b.cost())
}
