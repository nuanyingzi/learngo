package main

import (
	"fmt"
	"os"
)

// 学员信息管理系统

// 需求
// 1. 添加学员信息
// 2. 编辑学员信息
// 3. 展示所有学员信息

func showMenu() {
	fmt.Println("****欢迎来到学员信息管理系统****")
	fmt.Println("1. 添加学员")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 展示所有学员信息")
	fmt.Println("4. 退出系统")
}

// 获取用户输入的的学员信息
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("****请按要求输入学员信息****")
	fmt.Println("请输入学员的学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Println("请输入学员的姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入学员的班级：")
	fmt.Scanf("%s\n", &class)
	stu := newStudent(id, name, class) // 调用studnet的构造函数造一个学生
	return stu
}

func main() {
	sm := newStudentMgr()
	for {
		// 1. 打印系统菜单
		showMenu()
		// 2. 等待用户选择要执行的选项
		var input int
		fmt.Printf("请输入您要操作的序号：")
		fmt.Scanf("%d\n", &input)

		// 3. 执行用户选择的动作
		switch input {
		case 1: // 添加学员
			stu := getInput()
			sm.addStudent(stu)
		case 2: // 编辑学员
			stu := getInput()
			sm.modifyStudent(stu)
		case 3: // 展示所有学员
			sm.showStudent()
		case 4: // 退出
			os.Exit(0)
		default: // error
			fmt.Println("您输入的指令有误，请重新输入")
		}
	}

}
