package main

import "fmt"

func showMenu()  {
	fmt.Println("1.查看学生列表")
	fmt.Println("2.添加学生信息")
	fmt.Println("3.删除学生信息")
	fmt.Println("4.修改学生信息")
	fmt.Println("5.退出信息系统")
}

func main()  {
	for {
		showMenu()
		var choice int
		fmt.Scanln(&choice)
		switch choice {
			case 1:
				 
			}
	}
}
