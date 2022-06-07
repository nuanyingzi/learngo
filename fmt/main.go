package main

import (
	"bufio"
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

	// Sprint系列函数会把传入的数据生成并返回一个字符串
	s1 := fmt.Sprint("赣州小公主")
	name = "深圳小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("深圳小狗")
	fmt.Println(s1, s2, s3)

	// Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误
	//err := fmt.Errorf("这是一个错误")

	// 整型占位符
	n := 65
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	//err := fmt.Errorf("这是一个错误")

	/*var (
		studentName string
		score       int
		married     bool
	)
	fmt.Scan(&studentName, &score, &married)
	fmt.Scanf("1:%s 2:%d 3:%t", &studentName, &score, &married)
	fmt.Scanln(&studentName, &score, &married)
	fmt.Printf("扫描结果： studentName:%s score:%d married:%t \n", studentName, score, married)*/

	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	fmt.Printf("%#v\n", text)
}
