package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	功能包括：查看/新增/删除学生
*/

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id,
		name,
	}
}

var (
	allStudent map[int64]*student
)

func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号：%d  姓名：%s\n", k, v.name)
	}
}

func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&name)
	newStu := newStudent(id, name)
	allStudent[id] = newStu
}

func deleteStudent() {
	var deleteID int64
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&deleteID)
	delete(allStudent, deleteID)
}

func main() {
	allStudent = make(map[int64]*student, 50)

	for {
		// 1. 打印菜单
		fmt.Println("欢迎光临学生管理系统！")
		fmt.Println(`
1.查看所有学生
2.新增学生
3.删除学生
4.退出
	`)
		fmt.Println("请输入你要干啥：")

		// 2. 等待用户输入
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)

		// 3. 执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("非法输入！")
		}
	}

}
