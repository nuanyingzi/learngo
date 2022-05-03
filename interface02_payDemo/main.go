package main

import "fmt"

type Payer interface {
	Pay(int64)
}

type ZhiFuBao struct {
	// 支付宝
}

func Checkout(obj Payer) {
	// 支付100元
	obj.Pay(100)
}

// Pay 支付宝的支付方法
func (z *ZhiFuBao) Pay(amount int64) {
	fmt.Printf("使用支付宝付款：%.2f元。\n", float64(amount/100))
}

type WeChat struct {
	// 微信
}

func (w *WeChat) Pay(amount int64) {
	fmt.Printf("使用微信付款：%.2f元。\n", float64(amount/100))
}

func main() {
	Checkout(&ZhiFuBao{})
	Checkout(&WeChat{})
}
