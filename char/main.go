package main

import "fmt"

// 字符

func main() {
	// byte uint8的别名 ASCII码
	// rune int32的别名
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)                 // 99 99
	fmt.Printf("c1:%T c2:%T\n", c1, c2) // c1:uint8 c2:int32

	s := "hello深圳"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}
	// h
	// e
	// l
	// l
	// o
	// æ
	// ·
	// ±
	// å

	// ³

	for _, r := range s {
		fmt.Printf("%c\n", r)
	}
	// h
	// e
	// l
	// l
	// o
	// 深
	// 圳
}
