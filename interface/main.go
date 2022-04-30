package main

import "fmt"

// 接口

type Sayer interface {
	Say()
}

func MakeHungry(s Sayer) {
	s.Say()
}

type Cat struct {
}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog struct {
}

func (d Dog) Say() {
	fmt.Println("汪汪汪~")
}

type Sheep struct {
}

func (s Sheep) Say() {
	fmt.Println("咩咩咩~")
}

func main() {
	c1 := Cat{}
	c1.Say()
	d1 := Dog{}
	d1.Say()
	s1 := Sheep{}
	s1.Say()
}
