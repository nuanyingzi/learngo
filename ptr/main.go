package main

// 指针

import "fmt"

func modify1(x int) {
	x = 100
}
func modify2(x *int) {
	*x = 100
}

func main() {
	// 初识指针
	a := 111
	b := &a
	fmt.Printf("a:%d a的地址：%p\n", a, &a)
	fmt.Printf("b:%p b的类型：%T\n", b, b)

	// 指针传值
	c := 10
	modify1(c)
	fmt.Println(c) // 10
	modify2(&c)
	fmt.Println(c) // 100

	//var d *int
	//*d = 100
	//fmt.Println(*d)
    // 报错

	//var e map[string]int
	//e["深圳小王子"] = 100
	//fmt.Println(e)
	// 报错

	// 分配内存 new
	f := new(int)
	g := new(bool)
	fmt.Printf("f: %T\n",f)
	fmt.Printf("g: %T\n",g)
	fmt.Println(*f)
	fmt.Println(*g)

	// 分配内存 make
	var h map[string]int
	h = make(map[string]int, 10)
	h["力量"] = 10
	fmt.Println(h)

	// new与make的区别:
	// 1 二者都是用来做内存分配的。
	// 2 make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	// 3 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
}
