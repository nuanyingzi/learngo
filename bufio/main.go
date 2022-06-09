package main

import (
	"fmt"
	"io/ioutil"
)

// bufio 按行读取示例

// ioutil.ReadFile读取整个文件

func main() {
	/*file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("Open file failed", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // 这里是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("Read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}*/

	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("Open file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}
