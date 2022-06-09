package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*file, err := os.OpenFile("xxx.txt", os.O_CREATE|os.O_TRUNC|os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "祭祀阿斯弗红沙发哈桑理解和罚款收到回复"
	file.Write([]byte(str)) // 写入字节切片数据
	file.WriteString(str)   // 直接写入字符串数据*/

	str := "啥快递"
	err := ioutil.WriteFile("./xxxx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
