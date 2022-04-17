package main

import "fmt"

// 函数

func sayHello() {
	fmt.Println("Hello 丽丽")
}

// 定义一个接受string类型的name参数
func sayHello2(name string) {
	fmt.Println("hello, ", name)
}

// 定义接收多个参数的函数并且有一个返回值
func intSum(a int, b int) int {
	ret := a + b
	return ret
}
func intSum2(a int, b int) (ret int) {
	ret = a + b
	return
}

// 函数接收可变参数
func intSum3(a ...int) {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

func intSum4(a ...int) int {
	ret := 0
	for _, arg := range a {
		ret += arg
	}
	return ret
}

// 固定参数和可变参数同时出现时，可变参数要放在最后
func intSum5(a int, b ...int) int {
	ret := a + len(b)
	return ret
}

// Golang中参数类型简写
func intSum6(a, b int) {

}

// 定义具有多个返回值的函数
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	// 函数调用
	sayHello()
	sayHello2("丽丽")
	r := intSum(1, 9)
	fmt.Println(r)
	intSum3(10, 29)
	fmt.Println(intSum4(1, 9, 89, 10))
	fmt.Println(intSum5(1, 9, 89, 10))
	fmt.Println(calc(1, 9))
}
