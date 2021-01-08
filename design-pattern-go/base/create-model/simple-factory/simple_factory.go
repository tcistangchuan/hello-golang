package simple_factory

import (
	"fmt"
)

//创建工厂结构体:
type Factory struct {
}

//创建产品接口，这里为了方便，只写了一个方法，请根据自己的需要扩展
type Product interface {
	create()
}

//创建两个产品：产品1和产品2，它们实现了产品接口:
// 产品1，实现产品接口
type Product1 struct {
}

func (p1 Product1) create() {
	fmt.Println("this is product 1")
}

// 产品2，实现产品接口
type Product2 struct {
}

func (p1 Product2) create() {
	fmt.Println("this is product 2")
}

//为工厂结构体添加一个方法用于生产产品（实例化对象）:
func (f Factory) Generate(name string) Product {
	switch name {
	case "product1":
		return Product1{}
	case "product2":
		return Product2{}
	default:
		return nil
	}
}

func main(){
	// 创建一个工厂类，在应用中可以将这个工厂类实例作为一个全局变量
	factory := new(Factory)

	// 在工厂类中传入不同的参数，获取不同的实例
	p1 := factory.Generate("product1")
	p1.create() // output:   this is product 1

	p2 := factory.Generate("product2")
	p2.create() // output:   this is product 2

}