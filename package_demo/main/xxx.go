package main

import (
	"fmt"
	c "learngo/package_demo/calc"
)

// Go语言中不允许导入包而不使用
// Go语言中不允许循环引用包

// 多用来做一些初始化的操作
func init() {
	fmt.Println("xxx.init()")
}

// 当你的代码都放在GOPATH目录下的话
// 我们导入包的路径要从GOPATH/src后面继续写起
func main() {
	fmt.Println("Hello")
	ret := c.Add(1, 2)
	fmt.Println(ret)
	fmt.Println(c.Name)
}
