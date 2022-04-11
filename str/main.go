package main

import (
	"fmt"
	"strings"
)

// 字符串常见操作

func main() {
	// 求字符串的长度
	s1 := "hello"
	fmt.Println(len(s1)) // 5
	s2 := "hello南山"
	fmt.Println(len(s2)) //中文在字符串中占用三个字节 // 11

	// 拼接字符串
	fmt.Println(s1 + s2) // // hellohello南山
	s3 := fmt.Sprintf("%s - %s", s1, s2)
	fmt.Println(s3) // hello - hello南山

	// 字符串的分割
	s4 := "how do you do"
	fmt.Println(strings.Split(s4, " "))        // [how do you do]
	fmt.Printf("%T\n", strings.Split(s4, " ")) // []string

	// 判断是否包含
	fmt.Println(strings.Contains(s4, "do")) // true

	// 判断前缀
	fmt.Println(strings.HasPrefix(s4, "how")) // true

	// 判断子串的位置
	fmt.Println(strings.Index(s4, "do")) // 4

	// 子串最后出现的位置
	fmt.Println(strings.LastIndex(s4, "do")) // 11

	// join
	s5 := []string{"how", "do", "you", "do"}
	fmt.Println(s5)                    // [how do you do]
	fmt.Println(strings.Join(s5, "+")) // how+do+you+do
}
