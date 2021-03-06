package demo

import f "fmt"

// 包级别标识符的可见性

// num 定义一个全局整型变量
// 首字母小写,对外不可见(只能在包内使用)
var num = 100

// Mode 定义一个常量
// 首字母大写,对外可见(可在其他包使用)
const Mode = 1

// person 定义一个代表人的结构体
// 首字母小写,对外不可见(只能在包内使用)
type person struct {
	name string
	Age  int
}

// Add 返回两个整数和的函数
// 首字母大写,对外可见(可在其他包中使用)
func Add(x, y int) int {
	return x + y
}

// sayHi 打招呼的函数
// 首字母小写,对外不可见(只能在包内使用)
func sayHi() {
	myName := "tao" // 函数局部变量,只能在当前函数内使用
	f.Println("hello, ", myName)
}
