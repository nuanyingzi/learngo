package main

import "fmt"

func main() {
	// 初始化
	scoreMap := make(map[string]int, 8)
	fmt.Printf("%T\n", scoreMap)
	scoreMap["小张"] = 18
	scoreMap["小李"] = 20
	fmt.Println(scoreMap)

	// 2 初始化时填充元素
	userInfo := map[string]string{
		"username": "小王子",
		"password": "123456",
	}
	fmt.Println(userInfo)

	// 3 判断某个键是否存在
	// value, ok := map[key]
	v, ok := scoreMap["小张"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	// 4 map的遍历
	rangeMap := map[string]int{
		"张三": 90,
		"李四": 100,
		"伊人": 18,
	}
	for k, v := range rangeMap {
		fmt.Println(k, v)
	}
	for k := range rangeMap {
		fmt.Println(k)
	}
	for _, v := range rangeMap {
		fmt.Println(v)
	}
}
