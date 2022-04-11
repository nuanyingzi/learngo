package main

import "fmt"

// 常量
// const Pi = 3.14
// const e = 2.7

// const (
// 	n1 = 1
// 	n2
// 	n3
// 	n4
// )

// const (
// 	n1 = iota // 0
// 	n2        // 1
// 	n3		  // 2
// 	n4 		  // 3
// )

// const (
// 	n1 = iota // 0
// 	n2 		  // 1
// 	_		  //
// 	n4		  // 3
// )

// const (
// 	n1 = iota // 0
// 	n2 = iota // 1
// 	n3 = 100  // 100
// 	n4 = iota //3
// )
// const n5 = iota // 0

// const (
// 	_  = iota
// 	KB = 1 << (10 * iota) // 1024
// 	MB = 1 << (10 * iota) // 1048576
// 	GB = 1 << (10 * iota) // 1073741824
// 	TB = 1 << (10 * iota) // 1099511627776
// 	PB = 1 << (10 * iota) // 1125899906842624
// )

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

func main() {
	// fmt.Println(Pi, e)
	fmt.Println(a, b, c, d, e, f)
}
