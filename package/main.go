package main

import "fmt"

// 包的执行顺序

var x int8 = 10

const pi = 3.14

func init() {
	fmt.Println("x:", x)
	fmt.Println("pi:", pi)
	sayHi("world")
}

func sayHi(name string) {
	fmt.Println("Hello ", name)
}

func main() {
	fmt.Println("你好，世界")
}
