package main

import (
	"fmt"
)

/*策略模式：
	策略模式定义了算法族，分别封装起来，让他们之间可以相互替换。该模式让算法独立于使用它的客户而独立变化。

	应用场景：
	多个类只区别在表现行为不同，可以使用策略模式，在运行时动态选择具体要执行的行为。
	需要在不同情况下使用不同的策略(算法)，或者策略还可能在未来用其它方式来实现。
	对客户隐藏具体策略(算法)的实现细节，彼此完全独立。

	列子：
		我乘公交出去玩
		现在改变主意，我乘地铁出去玩
		现在改变主意，我自己开着车出去玩

链接：https://segmentfault.com/a/1190000007523487*/
/**
交通工具接口
*/
type travel interface {
	goBy()
}

/**
交通工具
*/
type common struct {
	flag int // 换乘标示
}

/**
公交车
*/
type bus struct {
	common
}

func (b *bus) goBy() {
	if b.flag == 1 {
		fmt.Println("change the type of travel to bus")
	} else {
		fmt.Println("start travel ...")
	}
	b.flag = 1
	fmt.Println("go travel by bus")
}

/**
摩托车
*/
type metro struct {
	common
}

func (m *metro) goBy() {
	if m.flag == 1 {
		fmt.Println("change the type of travel to metro")
	} else {
		fmt.Println("start travel ...")
	}
	m.flag = 1
	fmt.Println("go travel by metro")
}

/**
客户（知道如何选择策略）
*/
type mime struct {
	travel
}

func main() {
	m := &mime{
		travel: &bus{
			common: common{
				flag: 0,
			},
		},
	}
	m.goBy()
	m.travel = &metro{
		common: common{
			flag: m.travel.(*bus).flag,
		},
	}
	m.goBy()
	m.travel = &bus{
		common: common{
			flag: m.travel.(*metro).flag,
		},
	}
	m.goBy()
}
