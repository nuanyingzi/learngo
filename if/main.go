package main

import "fmt"

// if判断

func main() {
	// 1 基本写法
	// score := 65
	// if score >= 90 {
	// 	fmt.Println("A")
	// } else if score > 75 {
	// 	fmt.Println("B")
	// } else {
	// 	fmt.Println("C")
	// }
	// 2 if判断的特殊写法
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
