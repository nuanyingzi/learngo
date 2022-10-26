package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 读
	/*content, err := ioutil.ReadFile("./ioutil.go")
	if err != nil {
		fmt.Println("read file failed, err: ", err)
		return
	}
	fmt.Println(string(content))*/

	// 写
	str := "This is a kindful man, could you gentle to youself please?"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed", err)
		return
	}
}
