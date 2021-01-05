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

//添加学生
func (s StudentManage)addStudent()  {

}

