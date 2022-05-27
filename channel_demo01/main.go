package main

import "fmt"

func main() {
	/*var ch1 chan int // chan是引用类型，需要初始化之后才能使用
	ch1 = make(chan int, 1)*/

	ch1 := make(chan int, 1)
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
	close(ch1)
}
