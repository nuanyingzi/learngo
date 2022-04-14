package main

import "fmt"

// foe循环

func main() {
	// 1 基本写法
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 2 省略初始语句，但是必须保留初始语句后面的分号
	// var j = 0
	// for ; j < 10; j++ {
	// 	fmt.Println(j)
	// }

	// 3 省略初始语句与结束语句
	// var k = 10
	// for k > 0 {
	// 	fmt.Println(k)
	// 	k--
	// }

	// 4 无限循环
	// for {
	// 	fmt.Println("hello nihao")
	// }

	// 5 break跳出循环
	// for m := 0; m < 10; m++ {
	// 	fmt.Println(m)
	// 	if m == 3 {
	// 		break
	// 	}
	// }

	// 6 continue 继续下一次循环
	// for n := 0; n < 10; n++ {
	// 	if n == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	// 7 goto语句
	// 	for o := 0; o < 10; o++ {
	// 		for p := 0; p < o; p++ {
	// 			if p == 6 {
	// 				goto breakTag
	// 			}
	// 			fmt.Println(o, p)
	// 		}
	// 	}
	// breakTag:
	// 	fmt.Println("结束循环")

	// 打印99乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println("\t")
	}
}
