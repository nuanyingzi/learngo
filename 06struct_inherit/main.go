package main

// 结构体的继承
// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal // 通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪叫~\n", d.name)
}

func main() {
	d1 := Dog{
		Feet: 4,
		Animal: &Animal{
			name: "旺财",
		},
	}
	d1.move() // 旺财会动
	d1.wang() // 旺财会汪汪汪叫~
}
