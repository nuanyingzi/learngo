package main

// 基本数据类型

import (
	"fmt"
	"math"
)

func main() {
	// 十进制打印为二进制
	n := 10
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	// 八进制
	m := 075
	fmt.Printf("%d\n", m)
	fmt.Printf("%o\n", m)
	// 十六进制
	f := 0xff
	fmt.Println(f)
	fmt.Printf("%x\n", f)

	// uint8
	var age uint8 // 0~255
	age = 255
	fmt.Println(age)

	// 浮点数
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// 布尔值
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)

	// 字符串
	s1 := "Hello shenzhen"
	s2 := "你好 深圳"
	fmt.Println(s1)
	fmt.Println(s2)

	// 打印windows平台下的一个路径 c:\code\go.exe
	fmt.Println("c:\\code\\go.exe")

	fmt.Println("\t制表符\n换行符")

	s3 := `
多行字符串
在两个反引号之间的内容
会原样输出
\t
\n
`
	fmt.Println(s3)
}
