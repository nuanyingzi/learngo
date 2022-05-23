package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	ID   int
	Name string
}

// student的构造函数
func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

type class struct {
	Title    string    `json:"title"`
	Students []student `json:"student_list" db:"student" xml:"ss"`
}

func main() {
	// 创建一个班级变量c1
	c1 := class{
		Title:    "火箭101",
		Students: make([]student, 0, 20),
	}
	// 往班级中添加学生
	for i := 0; i < 10; i++ {
		temStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, temStu)
	}
	fmt.Printf("%#v\n", c1)

	// JSON序列化：Go语言中的数据 -> JSON格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marsjal failed, err:", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)

	// JSON反序列化：JSON格式的字符串 -> Go语言中的数据
	jsonStr := `{
					"Title":"火箭103",
					"Students":[
						{
							"ID":0,
							"Name":"stu00"
						},
						{
							"ID":1,
							"Name":"stu01"
						},
						{
							"ID":2,
							"Name":"stu02"
						},
						{
							"ID":3,
							"Name":"stu03"
						},
						{
							"ID":4,
							"Name":"stu04"
						},
						{
							"ID":5,
							"Name":"stu05"
						},
						{
							"ID":6,
							"Name":"stu06"
						},
						{
							"ID":7,
							"Name":"stu07"
						},
						{
							"ID":8,
							"Name":"stu08"
						},
						{
							"ID":9,
							"Name":"stu09"
						}
					]
				}`
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json unmarsjal failed, err:", err)
		return
	}
	fmt.Printf("%#v\n", c2)
}
