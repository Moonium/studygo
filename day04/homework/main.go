package main

import (
	"fmt"
	"os"
)

var smr studentMgr

// 菜单函数
func showMenu() {
	fmt.Println("-----Welcome to SMS!-----")
	fmt.Println(`
	1. 查看所有学生
	2. 添加学生
	3. 修改学生
	4. 删除学生
	5. 退出
	`)
}

func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		fmt.Print("请输入要进行的操作：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择的选项是：%d\n", choice)
		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("非法输入！")
		}
	}
}
