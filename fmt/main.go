package main

import (
	"fmt"
	"os"
)

func main() {
	//Print 直接输出内容
	fmt.Print("在终端打印信息\n")
	//Println 会在输出内容的结尾添加一个换行符
	fmt.Println("在此换行")
	//Printf 支持格式化输出字符串
	name := "tao"
	fmt.Printf("我是%v\n", name)

	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name = "赣州小公主"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写入信息:%s", name)
}
