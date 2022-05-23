package main

import (
	"fmt"
	"sync"
)

// 声明全局等待组变量
var wg sync.WaitGroup

func hello() {
	fmt.Println("Hello")
	wg.Done() // 告知当前goroutine完成
}

func countHello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello", i)
}

func main() {
	//wg.Add(1) // 登记一个goroutine
	//go hello()
	//fmt.Println("你好")
	//wg.Wait() // 阻塞等待登记的goroutine完成
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go countHello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
