package main

// 一个类型实现多个接口,多种类型实现同一接口

import "fmt"

// Sayer 接口
type Sayer interface {
	Say()
}

// Mover 接口
type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

// Say 实现Sayer接口
func (d *Dog) Say() {
	fmt.Printf("%s会汪汪叫\n", d.Name)
}

// Move 实现Mover接口
func (d *Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}

type Car struct {
	Brand string
}

func (c *Car) Move() {
	fmt.Printf("%s速度70迈\n", c.Brand)
}

// 一个接口的所有方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct {
}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer // 嵌入甩干器
}

// Wash 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func main() {
	var d = &Dog{Name: "旺财"}
	var s Sayer = d
	var m Mover = d

	s.Say() // 对Sayer类型调用Say方法
	// 旺财会汪汪叫
	m.Move() // 对Mover类型调用Move方法
	// 旺财会动

	var obj Mover
	obj = &Dog{Name: "大黄"}
	obj.Move()
	// 大黄会动

	obj = &Car{Brand: "宝马"}
	obj.Move()
	// 宝马速度70迈

}
