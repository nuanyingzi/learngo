package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// bufio按行读取示例

func main() {
	file, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读取结束")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err: ", err)
		}
		fmt.Println(line)
	}
}
