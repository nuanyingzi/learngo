package main

// 结构体指针

import "fmt"

type person struct {
	name, city string
	age        int8
}

func main() {
	// 1 创建指针类型结构体
	var p2 = new(person)
	fmt.Printf("%T\n", p2)       // *main.person
	fmt.Printf("p2 = %#v\n", p2) // p2 = &main.person{name:"", city:"", age:0}
	p2.name = "小公主"
	p2.city = "上海"
	p2.age = 43
	fmt.Printf("p2=%#v\n", p2) // p2=&main.person{name:"小公主", city:"上海", age:43}

	// 2 取结构体的地址进行实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)  // *main.person
	fmt.Printf("%#v\n", p3) // &main.person{name:"", city:"", age:0}
	p3.name = "西欧"
	p3.city = "星星"
	p3.age = 20
	fmt.Printf("%#v", p3) // &main.person{name:"西欧", city:"星星", age:20}
}
