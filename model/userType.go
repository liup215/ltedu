package model

type UserType struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

const (
	USER_TYPE_STUDENT = iota + 1
	USER_TYPE_TEACHER
	USER_TYPE_ADMIN
)

var UserTypeMap = map[int]string{
	USER_TYPE_STUDENT: "学生",
	USER_TYPE_TEACHER: "老师",
	USER_TYPE_ADMIN:   "管理员",
}
