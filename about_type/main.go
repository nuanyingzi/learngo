package main

import "fmt"

// 自定义类型和类型别名示例

// MyInt 1 自定义类型
// MyiInt 基于int类型的自定义类型
type MyInt int

// NewInt 2 类型别名
// NewInt int类型别名
type NewInt = int

func main() {
	var i MyInt
	fmt.Printf("type of i: %T value:%v\n", i, i)

	var x NewInt
	fmt.Printf("type of x: %T value:%v\n", x, x)
}
