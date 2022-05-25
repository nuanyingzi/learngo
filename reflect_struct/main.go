package main

import (
	"fmt"
	"reflect"
)

// 结构体反射

type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}
	// 通过反射来获取结构体中所有字段信息
	t := reflect.TypeOf(stu1)
	fmt.Printf("name: %v, kind: %v\n", t.Name(), t.Kind())
	// 遍历结构体变量的所有字段
	for i := 0; i < t.NumField(); i++ {
		// 根据结构体字段的索引去取字段
		filedObj := t.Field(i)
		fmt.Printf("name:%v,type:%v,tag:%v\n", filedObj.Name, filedObj.Type, filedObj.Tag)
		fmt.Println(filedObj.Tag.Get("json"), filedObj.Tag.Get("ini"))
	}

}
