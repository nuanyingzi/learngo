package main

import (
	"fmt"
	"io"
	"os"
)

// copyFile 实现copy文件功能
func copyFile(dstName, srcName string) (written int64, err error) {
	// 以只读方式打开文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s file failed, err: %v\n", srcName, err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open %s file failed, err: %v\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) // 调用io.Copy()拷贝内容
}

func main() {
	_, err := copyFile("dst.txt", "src.txt")
	if err != nil {
		fmt.Println("copy file failed, err: ", err)
		return
	}
	fmt.Println("copy done")
}
