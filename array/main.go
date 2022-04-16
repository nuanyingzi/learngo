package main

import "fmt"

// 数组相关内容

func main() {
	// var a [3]int
	// var b [4]int
	// fmt.Println(a)
	// fmt.Println(b)
	// // [0 0 0]
	// // [0 0 0 0]

	// // 数组的初始化
	// // 1 定义时使用初始值列表的方式初始化
	// var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	// fmt.Println(cityArray)
	// // 2 编译器推导数组的长度
	// var boolArray = [...]bool{true, false, false, true, false}
	// fmt.Println(boolArray)
	// // 3 指定索引值
	// A := [...]int{1: 3, 5: 2}
	// fmt.Println(A)
	// fmt.Printf("type of A: %T\n", A)
	// // 4 遍历数组
	// ergodicArray := [...]string{"苹果", "西瓜", "雪梨", "脐橙"}
	// // 4.1 for遍历
	// for i := 0; i < len(ergodicArray); i++ {
	// 	fmt.Println(ergodicArray[i])
	// }
	// // 4.2 for range遍历
	// for index, value := range ergodicArray {
	// 	fmt.Println(index, value)
	// }

	// // 5 二维数组
	// twoArray := [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	// fmt.Println(twoArray)
	// fmt.Println(twoArray[2][1])
	// fmt.Printf("二维数组：%T\n", twoArray)

	// // 6 二维数组的遍历
	// for _, v1 := range twoArray {
	// 	for _, v2 := range v1 {
	// 		fmt.Printf("%s\t", v2)
	// 	}
	// 	fmt.Println("\t")
	// }

	// // 7 数组是值类型，赋值或者传参会复制整个数组，因此改变副本的值，不会改变本身的值
	// n := [3]int{1, 2, 3}
	// m := [3][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// 	{5, 6},
	// }
	// modifyArray(n)
	// fmt.Println(n)
	// modifyArray2(m)
	// fmt.Println(m)
	// // [1 2 3]
	// // [[1 2] [3 4] [5 6]]

	// 题1 求数组[1, 3, 5, 7, 8]所有元素的和
	var sum int
	sumArray := [...]int{1, 3, 5, 7, 8}
	for _, value := range sumArray {
		sum += value
	}
	fmt.Println(sum)

	// 题2 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	operateArray := [...]int{1, 2, 0, 8, 10, 19, 20, 10, 8, 6, 5}
	tagret := 11
	for i := 0; i < len(operateArray); i++ {
		for j := 0; j <= i; j++ {
			if operateArray[i]+operateArray[j] == tagret {
				fmt.Printf("(%d %d)\t", i, j)
			}
		}
	}
}

// func modifyArray(x [3]int) {
// 	x[0] = 100
// }

// func modifyArray2(x [3][2]int) {
// 	x[2][1] = 1000
// }
