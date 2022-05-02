package main

import "fmt"

// 嵌套结构体

type Address struct {
	Provice, City string
}
type Person struct {
	Name, Gender string
	Age          int8
	Address      // 嵌套另外一个结构体
}

func main() {
	p1 := Person{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
		Address: Address{
			Provice: "江西",
			City:    "赣州",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Name, p1.Gender, p1.Age, p1.Address)
	fmt.Println(p1.Address.Provice) // 通过嵌套的匿名结构体字段访问其内部的字段
	fmt.Println(p1.Provice)         // 直接访问匿名结构体中的字段
}
