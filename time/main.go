package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Printf("current time:%v\n", now)
	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 时
	minute := now.Minute() // 分
	second := now.Second() // 秒
	fmt.Println(year, month, day, hour, minute, second)
}
