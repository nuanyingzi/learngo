package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&UserInfo{}) // 自动迁移

	u1 := UserInfo{1, "小米", "男", "敲键盘"}
	u2 := UserInfo{2, "赫赫", "女", "吃冰棒"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)

	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.First(uu)
}
