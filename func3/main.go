package main

import (
	"fmt"
	"strings"
)

// 匿名函数和闭包
// 定义一个函数它的返回值是一个函数
// 把函数作为返回值
func a(name string) func() {
	return func() {
		fmt.Println("hello world, my ", name)
	}
}

// 匿名函数作为函数的参数
func do(action func()) {
	action()
}

// 匿名函数作为函数的返回值 即 闭包
func getAction() func() {
	return func() {
		fmt.Println("匿名函数作为函数的返回值")
	}
}

// 闭包实例1 添加文件后缀
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//func calc(base int) (func(int) int, func(int) int) {
//	add := func(i int) int {
//		base += i
//		return base
//	}
//
//	sub := func(i int) int {
//		base -= i
//		return base
//	}
//	return add, sub
//}

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin(coins int, users []string, distribution map[string]int) (int, map[string]int) {
	for _, user := range users {
		for _, word := range user {
			switch word{
			case 'e','E':
				distribution[user] += 1
				coins -= 1
			case 'i', 'I':
				distribution[user] += 2
				coins -= 2
			case 'o', 'O':
				distribution[user] += 3
				coins -= 3
			case 'u', 'U':
				distribution[user] += 4
				coins -= 4
			default:
				break
			}
		}
	}
	return coins, distribution
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()

	b := func() {
		fmt.Println("立即执行匿名函数")
	}
	// 闭包 = 函数 + 外层变量的引用
	r := a("力量") // r此时就是一个闭包
	r()          //相当于执行了a函数内部的匿名函数

	txtFunc := makeSuffixFunc(".txt")
	jpgFunc := makeSuffixFunc(".jpg")

	fmt.Println(txtFunc("Test"))
	fmt.Println(jpgFunc("Test"))

	do(b)

	actionName := getAction()
	actionName()

	f1, f2 := calc(10)
	fmt.Println(f1(3), f2(4))

	x, y := dispatchCoin(coins,users,distribution)
	fmt.Println("还剩",x,"个金币\n")
	fmt.Println("分配方式为：")
	for v, k := range y {
		fmt.Println(v, k,"个")
	}
}
