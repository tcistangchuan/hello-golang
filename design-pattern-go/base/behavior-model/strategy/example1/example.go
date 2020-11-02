package main

import "fmt"

type Context struct {
}

type PaymentInterface interface {
	Pay(ctx *Context) error
	Refund(ctx *Context) error
}

type WXPay struct {

}
func NewWXPay() *WXPay {
	return &WXPay{}
}
func (W WXPay) Pay(ctx *Context) error {
	fmt.Println("wx pay success")
	return nil
}
func (W WXPay) Refund(ctx *Context) error {
	fmt.Println("wx refund success")
	return nil
}

type AliPay struct {

}

func NewAliPay() *AliPay {
	return &AliPay{}
}

func (a AliPay) Pay(ctx *Context) error {
	fmt.Println("ali pay success")
	return nil
}
func (a AliPay) Refund(ctx *Context) error {
	fmt.Println("ali refund success")
	return nil
}

const(
	// 微信支付
	ConstWxPay = iota
	// 支付宝支付
	ConstAliPay
)

func main(){
	ctx := &Context{}

	// 支付方式
	var instance PaymentInterface
	PayType := 0
	switch PayType {
	case ConstWxPay:
		instance = NewWXPay()
	case ConstAliPay:
		instance = NewAliPay()
	default:
		panic("无效的支付方式")
	}
	instance.Pay(ctx)
}





