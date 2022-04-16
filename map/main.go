package main

// map映射

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	// 光声明map类型 但是没有初始化 a就是初始值nil
	var a map[string]int
	fmt.Println(a == nil)
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

	// 5 删除键值对
	delete(rangeMap, "李四")
	for k := range rangeMap {
		fmt.Println(k)
	}

	// 6 按照某个固定顺序遍历map
	var studentMap = make(map[string]int, 100)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) // 0~99的随机整数
		studentMap[key] = value
	}
	for k, v := range studentMap {
		fmt.Println(k, v)
	}
	// 6.1 先取出所有的key存放到切片中
	var keys = make([]string, 0, 100)
	for k := range studentMap {
		keys = append(keys, k)
	}
	// 6.2 对key做排序
	sort.Strings(keys)
	// 6.3 按照排序后的key对studentMap排序
	for _, key := range keys {
		fmt.Println(key, studentMap[key])
	}

	// 7 map类型的切片
	var mapSlice = make([]map[string]int, 8, 8) // 只完成了切片的初始化
	// [nil nil nil nil nil nil nil nil ]

	mapSlice[0] = make(map[string]int, 8) // 完成了map的初始化

	// 8 值为切片的map
	var sliceMap = make(map[string][]int, 8) // 只完成了map的初始化
	value, ok := sliceMap["中国"]
	if ok {
		fmt.Println(value)
	} else {
		//sliceMap没有这个键
		sliceMap["中国"] = make([]int, 5)
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
		sliceMap["中国"][2] = 300
	}
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}

	// 题1 统计一个字符串中每个单词出现的次数
	// "how do you do"
	s := "how do you do"
	words := strings.Split(s, " ")
	fmt.Println(words)
	var wordCount = make(map[string]int, 10)
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			wordCount[word] = v + 1
		} else {
			wordCount[word] = 1
		}
	}
	fmt.Println(wordCount)
	// map[do:2 how:1 you:1]
}
