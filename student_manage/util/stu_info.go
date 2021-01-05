package util


//学生
type StudentInfo struct {
	id int64
	name string
}

//全部学生信息
var (
	allStudentInfo map[int64]StudentInfo
)

//管理者
type StudentManage struct {
	manage map[int64]StudentInfo
}


//查看学生信息
func (s StudentManage) lookStudentInfo(){

}

//添加学生信息
func (s StudentManage) addStudentInfo()  {

}

//删除学生信息
func (s StudentManage)  deleteStudentInfo() {

}

//编辑学生信息
func (s StudentManage) editStudentInfo(){

}
