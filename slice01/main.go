package main

import "fmt"

// 切片

func main() {
	// var a []string
	// var b []int
	// var c = []bool{false, true}
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// // 1 基于数组获得切片
	// n := [5]int{1, 99, 100, 28, 18}
	// m := n[1:4]
	// fmt.Println(m)
	// fmt.Printf("%T\n", m)
	// // 2 切片再次切片
	// k := m[0:len(m)]
	// fmt.Println(k)
	// fmt.Printf("%T\n", k)
	// // 3 make函数构造切片
	// d := make([]int, 5, 10)
	// fmt.Println(d)
	// fmt.Printf("%T\n", d)
	// // 通过len()函数获取切片的长度
	// fmt.Println(len(d))
	// // 通过cap()函数获取切片的容量
	// fmt.Println(cap(d))

	// // 4 nil的判断
	// var i []int     // 声明int类型切片
	// var j = []int{} // 声明并且初始化
	// p := make([]int, 0)
	// fmt.Println(i, len(i), cap(i))
	// if i == nil {
	// 	fmt.Println("i是一个nil")
	// }
	// fmt.Println(j, len(j), cap(j))
	// if j == nil {
	// 	fmt.Println("j是一个nil")
	// }
	// fmt.Println(p, len(p), cap(p))
	// if p == nil {
	// 	fmt.Println("p是一个nil")
	// }

	// // 5 切片的赋值拷贝
	// a := make([]int, 3) //[0, 0, 0]
	// b := a
	// b[0] = 100
	// fmt.Println(a)
	// fmt.Println(b)

	// // 6 切片的遍历
	// a := []int{1, 2, 3, 4, 5}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(i, a[i])
	// }
	// fmt.Println()
	// for index, value := range a {
	// 	fmt.Println(index, value)
	// }

	// 7 切片的扩容
	// var a []int // 此时它没有申请内存
	// a = append(a, 10)
	// fmt.Println(a)
	// for i := 0; i < 10; i++ {
	// 	a = append(a, i)
	// 	fmt.Printf("%v len:%d cap:%d ptr:%p\n", a, len(a), cap(a), a)
	// }

	// a = append(a, 1, 2, 3, 4, 5, 6)
	// fmt.Println(a)
	// b := []int{12, 13, 14, 15}
	// a = append(a, b...)
	// fmt.Println(a)

	// 8 copy函数
	// a := []int{1, 2, 3, 4, 5}
	// b := make([]int, 5, 5)
	// copy(b, a)
	// b[4] = 100
	// fmt.Println(a)
	// fmt.Println(b)

	// 9 切片删除元素
	a := []string{"北京", "上海", "广州", "深圳"}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)
}
