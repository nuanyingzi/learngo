package main

import "fmt"

// 函数进阶

// 定义全局变量
var num int = 10

// 定义函数
func testGlobal() {
	num := 100
	// fmt.Println("全局变量 ", num)
	fmt.Println("变量num ", num)
}

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	testGlobal()

	// 变量i此时只在for循环的语句块中生效
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	r1 := calc(10, 20, add)
	fmt.Println(r1)
	r2 := calc(10, 20, sub)
	fmt.Println(r2)
}
