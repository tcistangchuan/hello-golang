package main

import "fmt"

/*适配器模式

适配器模式将某个类的接口转换程你所期望的另一个接口的表示，
主要目的是提高兼容性，让原本因为接口不匹配不能一起工作的两
个类可以协同工作。简单来说A类的类型与B类所需的类型不同，通
过适配器类转换成B类所需的类，适配器就是在A类下的类型转换作用。
*/

/**
纸制书接口
*/
type PaperBookInterface interface {
	open()
	turnPage()
}

/**
纸制书
*/
type book struct {
}

func (*book) open() {
	fmt.Println("open book")
}
func (*book) turnPage() {
	fmt.Println("turn page")
}

/**
电子书转换纸制书的适配器
*/
type EBookAdapter struct {
	ebook EBookInterface
}

func (e *EBookAdapter) open() {
	fmt.Println("==start==")
	e.ebook.pressStart()
	fmt.Println("==end==")
}
func (e *EBookAdapter) turnPage() {
	fmt.Println("==start==")
	e.ebook.pressNext()
	fmt.Println("==end==")
}

/**
电子书接口
*/
type EBookInterface interface {
	pressStart()
	pressNext()
}

/**
电子书
*/
type EBook struct {
}

func (*EBook) pressStart() {
	fmt.Println(" open ebook")
}
func (*EBook) pressNext() {
	fmt.Println(" press next")
}

func main() {
	var eb PaperBookInterface = &EBookAdapter{
		ebook: &EBook{},
	}
	eb.open()
	fmt.Println("")
	eb.turnPage()
}
