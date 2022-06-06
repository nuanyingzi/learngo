package main

import (
	"fmt"
	"time"
)

func main() {

	/*now := time.Now()
	fmt.Printf("current time:%v\n", now)
	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 时
	minute := now.Minute() // 分
	second := now.Second() // 秒
	fmt.Println(year, month, day, hour, minute, second)

	timestamp := now.Unix()      // 秒级时间戳
	timeMilli := now.UnixMilli() // 毫秒级时间戳
	timeMicro := now.UnixMicro() // 微秒级时间戳
	timeNano := now.UnixNano()   // 纳秒级时间戳
	fmt.Println(timestamp, timeMilli, timeMicro, timeNano)

	// 将时间戳转换为具体的时间格式
	t := time.Unix(1654527617+3600, 0)
	fmt.Println(t)

	// 时间间隔
	n := 5
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("over")*/

	/*now := time.Now()
	later := now.Add(time.Hour)
	fmt.Println(later)
	later2 := now.Sub(later)
	fmt.Println(later2)*/

	/*// 定时器
	for tmp := range time.Tick(time.Second) {
		fmt.Println(tmp)
	}*/

	// 时间格式化： Y-m-d H:M:S
	//             2006-01-02 15:04:05
	/*now := time.Now()
	ret1 := now.Format("2006-01-02 15:04:05")
	fmt.Println(ret1)
	ret2 := now.Format("2006-01-02 03:04:05 PM")
	fmt.Println(ret2)*/

	// 解析字符串类型的时间
	timeStr := "2022/10/31 08:11:11"
	// 1 拿到时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2 根据时区去解析一个字符串格式的时间
	timeObj, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil {
		fmt.Printf("parse timeStr failed:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	timeObj2, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Printf("parse timeStr2 failed:%v\n", err)
		return
	}
	fmt.Println(timeObj2)

}
