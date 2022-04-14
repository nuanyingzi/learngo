package main

import "fmt"

// 数组相关内容

func main() {
	var a [3]int
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)
	// [0 0 0]
	// [0 0 0 0]

	// 数组的初始化
	// 1 定义时使用初始值列表的方式初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(cityArray)
	// 2 编译器推导数组的长度
	var boolArray = [...]bool{true, false, false, true, false}
	fmt.Println(boolArray)
}
