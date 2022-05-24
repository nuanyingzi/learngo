package main

import "fmt"

type student struct {
	id    int // 学号唯一，唯一标识符
	name  string
	class string
}

// student的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id,
		name,
		class,
	}
}

// 学员管理的类型
type studentMgr struct {
	allStudents []*student
}

func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 1 添加学生
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2 编辑学生信息
func (s *studentMgr) modify(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id { // 传入的学号相等时，就找到了该学生
			s.allStudents[i] = newStu // 根据切片的索引直接赋值新学生
			return
		}
	}
	fmt.Printf("你输入的学生信息有误，没有找到学号为%d的学生", newStu.id)
}

// 3 展示所有学生信息
func (s *studentMgr) showStudentsList() {
	for _, v := range s.allStudents {
		fmt.Printf("学号： %d 姓名：%s 班级：%s\n", v.id, v.name, v.class)
	}
}
