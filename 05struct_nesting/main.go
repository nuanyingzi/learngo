package main

import "fmt"

// Person 结构体的匿名字段
type Person struct {
	string
	int
}

// 嵌套结构体

// Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

// Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

// User 用户结构体
type User struct {
	Name    string
	Age     int8
	Address // 匿名字段
	Email
}

func main() {
	p1 := Person{
		"小公",
		18,
	}
	fmt.Printf("%#v\n", p1)        // main.Person{string:"小公", int:18}
	fmt.Println(p1.string, p1.int) // 小公 18

	user1 := User{
		Name: "先生",
		Age:  18,
		Address: Address{
			Province: "江孜",
			City:     "庐陵",
		},
	}
	fmt.Printf("user=%#v\n", user1) // user=main.User{Name:"先生", Age:18, Address:main.Address{Province:"江孜", City:"庐陵"}}

	var user2 User
	user2.Name = "小种"
	user2.Age = 20
	user2.Address.Province = "河北"    // 匿名字段默认使用类型名作为字段名
	user2.City = "常州"                // 匿名字段可以省略, 当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找
	fmt.Printf("user2=%#v\n", user2) // user2=main.User{Name:"小种", Age:20, Address:main.Address{Province:"河北", City:"常州"}}

	var user3 User
	user3.Name = "嘻嘻"
	user3.Age = 8
	//user3.CreateTime = "2020年6月12号" // 报错：ambiguous selector user3.CreateTime
	user3.Address.CreateTime = "2020年6月12号"
	user3.Email.CreateTime = "2020年6月12号" // 嵌套结构体的字段名冲突时，需指定哪个结构体的字段
	fmt.Printf("user3=%#v\n", user3)
	// user3=main.User{Name:"嘻嘻", Age:8, Address:main.Address{Province:"", City:"", CreateTime:"2020年6月12号"}, Email:main.Email{Account:"", CreateTime:"2020年6月12号"}}

}
