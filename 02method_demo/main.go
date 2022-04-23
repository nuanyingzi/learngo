package main

import "fmt"

// 为任意类型添加方法

type MyInt int

// 基于内置的基本类型自定义一个方法
func (m MyInt) sayHi() {
	fmt.Println("Hi!")
}

func main() {
	var m1 MyInt
	m1 = 100
	m1.sayHi()
}
