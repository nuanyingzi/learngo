package main

import "fmt"

// 空接口
// 接口中没有定义任何需要实现的方法时，该接口就是一个空接口
// 任意类型都实现了空接口 --》 空接口变量可以存储任意值

// 空接口一般不需要提前定义
type xxx interface {
}

// 空接口的应用
// 1. 空接口类型作为函数的参数
// 2. 空接口可以作为map的value

func main() {
	var x interface{}
	//x = "hello"
	//x = 100
	x = false
	//fmt.Println(x)

	//var m = make(map[string]interface{}, 15)
	//m["name"] = "娜扎"
	//m["age"] = 19
	//m["hobby"] = []string{"唱", "跳", "rap"}
	//fmt.Println(m)

	ret, ok := x.(string) // 类型断言 猜的类型不对时，ok=false，ret=string类型的零值
	if !ok {
		fmt.Println("不是字符串类型")
	} else {
		fmt.Println("是字符串类型", ret)
	}
	// 使用switch进行类型断言
	switch v := x.(type) {
	case string:
		fmt.Printf("是字符串类型，value：%v\n", v)
	case bool:
		fmt.Printf("是布尔类型，value：%v\n", v)
	case int:
		fmt.Printf("是int类型，value：%v\n", v)
	default:
		fmt.Printf("猜不到了，value：%v\n", v)
	}
}
