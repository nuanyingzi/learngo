package main

import "fmt"

// 运算符

func main() {
	// 1 算术运算符
	a := 10
	b := 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++
	fmt.Println(a) // 11
	a--
	fmt.Println(a) // 10

	// 2 关系运算符
	fmt.Println(10 > 2)
	fmt.Println(10 != 2)
	fmt.Println(10 < 2)

	// 3 逻辑运算符
	fmt.Println(10 > 5 && 1 == 1) // true
	fmt.Println(!(10 > 5))        // false
	fmt.Println(1 > 5 || 1 == 1)  // true

	// 4 位运算符
	c := 1              // 001
	d := 5              // 101
	fmt.Println(c & d)  // 001 => 1
	fmt.Println(c | d)  // 101 => 5
	fmt.Println(c ^ d)  // 100 => 4
	fmt.Println(1 << 2) // 100 => 4
	fmt.Println(4 >> 2) // 001 => 1

	// 5 赋值运算符
	var e int
	e = 5
	e += 5
	fmt.Println(e) // 10
}
