package main

import "fmt"

// channel

func retc(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	// 无缓冲的通道
	ch := make(chan int)
	go retc(ch) // 创建一个goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")

	// 有缓冲的通道
	ch2 := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch2 <- 10
	fmt.Println("发送成功")
}
