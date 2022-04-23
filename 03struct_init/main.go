package main

import "fmt"

// 结构体的初始化
// 1 键值对初始化
// 2 值的列表初始化

type person struct {
	name, city string
	age        int8
}

func main() {
	// 1 键值对初始化
	p4 := person{
		name: "阿萨",
		city: "神奈川",
		age:  10,
	}
	fmt.Printf("%#v\n", p4) // main.person{name:"阿萨", city:"神奈川", age:10}

	p5 := &person{
		name: "阿达",
		age:  19,
	}
	fmt.Printf("%#v\n", p5) // &main.person{name:"阿达", city:"", age:19}

	// 2 值的列表初始化
	p6 := person{
		"小小",
		"赣州",
		20,
	}
	fmt.Printf("%#v\n", p6) // main.person{name:"小小", city:"赣州", age:20}

}
