package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main() {
	/*log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetPrefix("[甜甜圈日志]")
	log.Println("这是一条普通的日志")
	v := "普通"
	log.Printf("这是一条%v的日志", v)
	log.Fatalln("这是一条会触发fatal的日志")
	log.Panicln("这是一条会触发panic的日志")*/

	log.Println("这是一条普通的日志")
	log.SetPrefix("[元气少女]")
	log.Println("这是一条带有前缀并写入文件的日志")
}
