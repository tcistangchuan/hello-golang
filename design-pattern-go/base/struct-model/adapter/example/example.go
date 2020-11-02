package main

import "fmt"

// 纸质书接口
type ABook interface{
	open()
	turnPage()
}

// 电子书接口
type BBook interface {
	BOpen()
	BTurnPage()
}

// 电子书对象
type BObj struct {
}
func (B *BObj) BOpen() {
	fmt.Println("BBook open")
}

func (B *BObj) BTurnPage() {
	fmt.Println("BBook turnPage")
}

// 电子书-》纸质书的接口
type adapter struct {
	BBook BBook
}

func (a *adapter) open() {
	a.BBook.BOpen()
	fmt.Println(">>>>>>>>>>>>>>>>")
}

func (a *adapter) turnPage() {
	a.BBook.BTurnPage()
	fmt.Println("<<<<<<<<<<<<<<<<")

}

func main()  {
	// 电子书
	var b BBook = &BObj{}

	// 适配器(转换为纸质书)
	var a  ABook = &adapter{
		b,
	}
	a.open()
	a.turnPage()
}





