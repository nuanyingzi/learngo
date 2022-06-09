package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	/*// 以只读方式打开文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("Open file failed, err: ", err)
		return
	}
	// 关闭文件
	defer file.Close()
	// 使用Read方法读取文件
	var tmp = make([]byte, 512)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))*/

	// 循环读取
	// 以只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("Open file failed", err)
	}
	// 关闭文件
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("Read file failed", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}
