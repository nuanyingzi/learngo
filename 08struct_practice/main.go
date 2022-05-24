package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("****欢迎来到学院管理系统****")
	fmt.Println("1. 添加学生")
	fmt.Println("2. 编辑学生信息")
	fmt.Println("3. 展示所有学生信息")
	fmt.Println("4. 退出系统")
}

func getInputStudent() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("****请按要求输入学员信息****")
	fmt.Println("请输入学员学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Println("请输入学员姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入学员班级：")
	fmt.Scanf("%s\n", &class)
	newStu := newStudent(id, name, class)
	return newStu
}

func main() {
	sm := newStudentMgr()
	for {
		// 1. 打印系统菜单
		showMenu()
		var input int
		fmt.Scanf("%d\n", &input)
		switch input {
		case 1: // 添加学生
			stu := getInputStudent()
			sm.addStudent(stu)
			fmt.Printf("添加成功!\n")
		case 2: // 编辑学生
			stu := getInputStudent()
			sm.modify(stu)
			fmt.Printf("编辑成功!\n")
		case 3: // 展示所有学生
			sm.showStudentsList()
		case 4: // 退出系统
			fmt.Println("正在退出")
			os.Exit(0)
		default:
			fmt.Println("你输入的指令有误，请重新输入")
		}
	}

}
