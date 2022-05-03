package main

import "fmt"

// Mover 定义一个接口类型
type Mover interface {
	Move()
}

// Dog 狗结构体类型
type Dog struct {
}

// Move 使用值接收者定义move方法实现Mover接口
func (d Dog) Move() {
	fmt.Println("狗会动")
}

// Cat 猫结构体类型
type Cat struct {
}

// Move 使用指针接收者定义Move方法实现Mover接口
func (c *Cat) Move() {
	fmt.Println("猫会动")
}

func main() {
	var x Mover
	var d1 = Dog{} // d1是Dog类型
	x = d1
	x.Move()
	// 狗会动

	var d2 = &Dog{} // d2是Dog指针类型
	x = d2
	x.Move()
	// 狗会动

	var c1 = &Cat{} // c1是*Cat类型
	x = c1          // 可以将c1当成Mover类型
	c1.Move()
	// 猫会动

	//var c2 = Cat{}
	//x = c2 // 无法将 'c2' (类型 Cat) 用作类型 Mover 类型未实现 'Mover'，因为 'Move' 方法有指针接收器
	//c2.Move()
}
