package util

import (
	"fmt"
)

//学生
type StudentInfo struct {
	id int64
	name string
}

//全部学生信息
var (
	allStudentInfo map[int64]StudentInfo
	IfLoop bool=true
)

//管理者
type StudentManage struct {
	manage map[int64]StudentInfo
}


//查看学生信息
func (s StudentManage) LookStudentInfo(){
	for id,info:= range allStudentInfo{
	  fmt.Println(id,info)
	}
}

//添加学生信息
func (s StudentManage) AddStudentInfo()  {
	var id int64
	var name string
    fmt.Println("请输入学生id")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名")
	fmt.Scanln(&name)

	//map 没有初始化，首先初始化
	if allStudentInfo==nil{
		allStudentInfo = make(map[int64]StudentInfo)
	}
	allStudentInfo[id] = StudentInfo{
		id:id,
		name: name,
	}
}

//删除学生信息
func (s StudentManage)  DeleteStudentInfo() {
    var id int64
    fmt.Println("请输入学生编号")
    fmt.Scanln(&id)

    if _,ok:=allStudentInfo[id];ok{
		delete(allStudentInfo,id)
		fmt.Println("删除学生成功")
	}else{
		fmt.Println("不存在此编号的学生")
	}
}

//编辑学生信息
func (s StudentManage) EditStudentInfo(){
    var id int64
    var name string
    fmt.Println("请输人学生编号")
    fmt.Scanln(&id)

	if _,ok:=allStudentInfo[id];!ok{
		fmt.Println("不存在此编号的学生!!!!!")
		return
	}

    fmt.Println("请输入学生名字")
    fmt.Scanln(&name)

	allStudentInfo[id] = StudentInfo{
		id:id,
		name:name,
	}
}

//退出信息系统
func (s StudentManage) ExitSystem(){
    IfLoop=false
}
