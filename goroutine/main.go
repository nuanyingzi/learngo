package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello ", i)
	wg.Done() // 通知wg把计数器-1
}

func main() { // 开启一个主goroutine去执行main函数

	for i := 0; i < 10000; i++ {
		wg.Add(1)   // 计数牌+1
		go hello(i) // 开启了一个goroutine去执行hello这个函数
	}

	fmt.Println("hello main")
	wg.Wait() // 阻塞 等所有小弟都干完活之后才收兵
}
