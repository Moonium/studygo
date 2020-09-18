package main

import "fmt"

// 学生管理系统的方法实现

type student struct {
	id   int64
	name string
}

type studentMgr struct {
	allStudent map[int64]student
}

func (s studentMgr) showStudent() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", stu.id, stu.name)
	}
}

func (s studentMgr) addStudent() {
	var (
		stuID   int64
		stuName string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	newStu := student{
		stuID,
		stuName,
	}
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加成功！")
}

func (s studentMgr) editStudent() {
	var stuID int64
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人！")
		return
	}
	fmt.Printf("你要修改的学生信息如下：\n学号：%d 姓名：%s\n", stuObj.id, stuObj.name)
	fmt.Print("请输入学生的新名字：")
	var newName string
	fmt.Scanln(&newName)
	stuObj.name = newName
	s.allStudent[stuID] = stuObj
	fmt.Println("修改成功！")
}

func (s studentMgr) deleteStudent() {
	var stuID int64
	fmt.Print("请输入要删除学生的学号：")
	fmt.Scanln(&stuID)
	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	delete(s.allStudent, stuID)
	fmt.Println("删除成功！")
}
