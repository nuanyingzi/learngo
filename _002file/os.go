package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 只读方式打开当前目录的os.go文件
	file, err := os.Open("./os.go")
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	// 关闭文件
	defer file.Close()
	/*// 使用Read方法读取文件
	var tmp = make([]byte, 1024)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
	}
	if err != nil {
		fmt.Println("read file failed, err: ", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))*/

	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 1024)

	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读取结束")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err: ", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}
