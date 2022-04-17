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

// 闭包实例1 添加文件后缀
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()

	func() {
		fmt.Println("立即执行匿名函数")
	}()
	// 闭包 = 函数 + 外层变量的引用
	r := a("力量") // r此时就是一个闭包
	r()          //相当于执行了a函数内部的匿名函数

	txtFunc := makeSuffixFunc(".txt")
	jpgFunc := makeSuffixFunc(".jpg")

	fmt.Println(txtFunc("Test"))
	fmt.Println(jpgFunc("Test"))

}
