package main

import "log"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetPrefix("[甜甜圈日志]")
	log.Println("这是一条普通的日志")
	v := "普通"
	log.Printf("这是一条%v的日志", v)
	log.Fatalln("这是一条会触发fatal的日志")
	log.Panicln("这是一条会触发panic的日志")
}
