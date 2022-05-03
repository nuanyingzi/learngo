package main

import "fmt"

// 接口

type Sayer interface {
	Say()
}

// MakeHungry 饿肚子了
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

// MakeCatHungry 猫饿了会喵喵喵~
func MakeCatHungry(c Cat) {
	c.Say()
}

// MakeSheepHungry 羊饿了会咩咩咩~
func MakeSheepHungry(s Sheep) {
	s.Say()
}

func main() {
	//c1 := Cat{}
	//c1.Say()
	//d1 := Dog{}
	//d1.Say()
	//s1 := Sheep{}
	//s1.Say()
	var c Cat
	MakeHungry(c)
	var d Dog
	MakeHungry(d)

	var x Sayer
	a := Cat{}
	b := Dog{}
	x = a
	x.Say()
	x = b
	x.Say()
}
