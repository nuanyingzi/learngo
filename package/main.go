package main

import f "fmt"

// 包的执行顺序

var x int8 = 10

const pi = 3.14

func init() {
	f.Println("x:", x)
	f.Println("pi:", pi)
	sayHi("world")
}

func sayHi(name string) {
	f.Println("Hello ", name)
}

func main() {
	f.Println("你好，世界")
}
