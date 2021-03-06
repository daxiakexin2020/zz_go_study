package main

import (
	"fmt"
	"student_manage/util"
)

func showMenu() {
	fmt.Println("1.查看学生列表")
	fmt.Println("2.添加学生信息")
	fmt.Println("3.删除学生信息")
	fmt.Println("4.修改学生信息")
	fmt.Println("5.退出信息系统")
}

func main() {

	studentManage := util.StudentManage{

	}

	for {
		if util.IfLoop==false{
			break
		}
		showMenu()
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			studentManage.LookStudentInfo()
			break
		case 2:
			studentManage.AddStudentInfo()
			break

		case 3:
			studentManage.DeleteStudentInfo()
			break
		case 4:
			studentManage.EditStudentInfo()
			break
		case 5:
			studentManage.ExitSystem()
			break
		}
	}
}
