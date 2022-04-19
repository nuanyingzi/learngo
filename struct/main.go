package main
// 结构体
import "fmt"

type person struct {
	name, city string
	age int8
}

func main() {
	// 1 基本实例化
	var p1 person
	p1.name = "小王子"
	p1.city = "赣州"
	p1.age = 18
	fmt.Printf("p1: %v\n", p1) // p1: {小王子 赣州 18}
	fmt.Printf("p1: %#v\n", p1) // p1: main.person{name:"小王子", city:"赣州", age:18}

	// 2 匿名结构体
	var user struct{Name string; Age int}
	user.Name = "小偶像"
	user.Age = 19
	fmt.Printf("user: %#v\n", user) // user: struct { Name string; Age int }{Name:"小偶像", Age:19}

	// 3 创建指针类型结构体
	var p2 = new(person)
	fmt.Printf("%T\n", p2) // *main.person
	fmt.Printf("p2 = %#v\n", p2) // p2 = &main.person{name:"", city:"", age:0}
	p2.name = "小公主"
	p2.city = "上海"
	p2.age = 43
	fmt.Printf("p2=%#v\n", p2) // p2=&main.person{name:"小公主", city:"上海", age:43}
}
