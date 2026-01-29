package model

type Grade struct {
	Model
	Name               string  `json:"name"`
	GradeLeadTeacherId uint    `json:"gradeLeadTeacherId"`
	GradeLeadTeacher   Teacher `json:"gradeLeadTeacher"`
}

type GradeQuery struct {
	ID             uint `json:"id"`
	GradeTeacherId uint `json:"gradeTeacherId"`
	Page
}

type GradeCreateEditRequest struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	GradeLeadTeacherId uint   `json:"gradeLeadTeacherId"`
}

type ClassType struct {
	Model
	Name   string `json:"name"`
	IsMain int    `json:"isMain"`
}

type ClassTypeQuery struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	IsMain int    `json:"isMain"`
	Page   Page
}

type ClassTypeCreateEditRequest struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	IsMain int    `json:"isMain"`
}

type Class struct {
	Model
	Name                   string    `json:"name"`
	GradeId                uint      `json:"gradeId"`
	Grade                  Grade     `json:"grade"`
	SubjectTeacherId       uint      `json:"subjectTeacherId" gorm:"index"`
	SubjectTeacher         Teacher   `json:"subjectTeacher"`
	ClassTypeId            uint      `json:"classTypeId"`
	ClassType              ClassType `json:"classType"`
	HeadTeacherId          uint      `json:"headTeacherId"`
	HeadTeacher            Teacher   `json:"headTeacher"`
	DeputyHeadTeacherId    uint      `json:"deputyHeadTeacherId" `
	DeputyHeadTeacher      Teacher   `json:"deputyHeadTeacher"`
	IntershipHeadTeacherId uint      `json:"intershipHeadTeacherId"`
	IntershipHeadTeacher   Teacher   `json:"intershipHeadTeacher"`
	Students               []*User   `json:"students" gorm:"many2many:user_class_relation;"`
}

type ClassCreateEditRequest struct {
	ID                     uint   `json:"id"`
	Name                   string `json:"name"`
	GradeId                uint   `json:"gradeId"`
	SubjectTeacherId       uint   `json:"subjectTeacherId" gorm:"index"`
	ClassTypeId            uint   `json:"classTypeId"`
	HeadTeacherId          uint   `json:"headTeacherId"`
	DeputyHeadTeacherId    uint   `json:"deputyHeadTeacherId" `
	IntershipHeadTeacherId uint   `json:"intershipHeadTeacherId"`
}

const (
	CLASS_TEACHER_TYPE_SUBJECT                = 1
	CLASS_TEACHER_TYPE_HEAD_TEACHER           = 2
	CLASS_TEACHER_TYPE_DEPUTY_HEAD_TEACHER    = 3
	CLASS_TEACHER_TYPE_INTERSHIP_HEAD_TEACHER = 4
	CLASS_TEACHER_TYPE_GRADE_LEAD_TEACHER     = 5
	CLASS_TEACHER_TYPE_ADMIN                  = 6
)

type ClassQuery struct {
	ID                     uint `json:"id"`
	TeacherId              uint `json:"-"`
	TeacherType            uint `json:"teacherType"` // 1. 学科老师, 2. 班主任 3. 副班 4. 实习班主任
	GradeId                uint `json:"gradeId"`
	SubjectTeacherId       uint `json:"subjectTeacherId"`
	ClassTypeId            uint `json:"classTypeId"`
	HeadTeacherId          uint `json:"headTeacherId"`
	DeputyHeadTeacherId    uint `json:"deputyHeadTeacherId" `
	IntershipHeadTeacherId uint `json:"intershipHeadTeacherId"`
	Page
}
