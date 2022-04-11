package main

import (
	"fmt"
)

func main() {
	// 标准声明
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)
	// 批量声明
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)
	// 声明变量指定初始值
	var name1 string = "小王子"
	var age1 int = 18
	fmt.Println(name1, age1)
	var name2, age2 = "小公主", 17
	fmt.Println(name2, age2)
	// 类型推导，让编译器根据变量的初始值推导其类型
	var name3 = "小虾米"
	var age3 = 1
	fmt.Println(name3, age3)
	// 短变量声明
	name4 := "小李"
	fmt.Println(name4)
}

// 匿名变量
// func foo() (int, string) {
// 	return 10, "Qimi"
// }

// func main() {
// 	x, _ := foo()
// 	_, y := foo()
// 	fmt.Println("x = ", x)
// 	fmt.Println("y= ", y)
// }
