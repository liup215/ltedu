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

const (
	// ClassTypeTeaching 教学班 - 用于管理教学，用户可加入多个
	ClassTypeTeaching = 1
	// ClassTypeAdministrative 行政班 - 用于管理日常行为，每个用户只能属于一个
	ClassTypeAdministrative = 2
)

type Class struct {
	Model
 	Name        string    `json:"name"`
	ClassType   int       `json:"classType" gorm:"default:1"` // 1: 教学班, 2: 行政班
	InviteCode  string    `json:"inviteCode" gorm:"uniqueIndex;size:32"`
	AdminUserId uint      `json:"adminUserId" gorm:"index"`
	AdminUser   *User     `json:"adminUser,omitempty" gorm:"foreignKey:AdminUserId"`
	Students    []*User   `json:"students" gorm:"many2many:user_class_relation;"`
	Teachers    []*User   `json:"teachers,omitempty" gorm:"many2many:class_teacher_relation;"`
	SyllabusId  *uint     `json:"syllabusId,omitempty" gorm:"index"` // 教学班绑定的syllabus（仅教学班使用）
	Syllabus    *Syllabus `json:"syllabus,omitempty" gorm:"foreignKey:SyllabusId"`
}

const (
	ClassJoinStatusPending  = 0
	ClassJoinStatusApproved = 1
	ClassJoinStatusRejected = 2
)

// ClassStudentStatus constants for the student's enrollment status within a class.
const (
	ClassStudentStatusStudying    = 1 // 在读
	ClassStudentStatusGraduated   = 2 // 结业
	ClassStudentStatusTransferred = 3 // 转走
	ClassStudentStatusDropped     = 4 // 弃科
)

// UserClassRelation is the explicit join-table model for the Class ↔ User many2many
// association. It extends the default join table with a student status field.
type UserClassRelation struct {
	ClassId uint `gorm:"primaryKey" json:"classId"`
	UserId  uint `gorm:"primaryKey" json:"userId"`
	Status  int  `gorm:"default:1"  json:"status"` // 1:在读, 2:结业, 3:转走, 4:弃科
}

func (UserClassRelation) TableName() string {
	return "user_class_relation"
}

// ClassStudentView combines a User with their status inside a specific class.
type ClassStudentView struct {
	User
	StudentStatus int `json:"studentStatus" gorm:"column:student_status"`
}

type UpdateStudentStatusRequest struct {
	ClassId uint `json:"classId" binding:"required"`
	UserId  uint `json:"userId" binding:"required"`
	Status  int  `json:"status" binding:"required"`
}

type ClassJoinRequest struct {
	Model
	ClassId uint   `json:"classId" gorm:"index"`
	Class   *Class `json:"class,omitempty" gorm:"foreignKey:ClassId"`
	UserId  uint   `json:"userId" gorm:"index"`
	User    *User  `json:"user,omitempty" gorm:"foreignKey:UserId"`
	Status  int    `json:"status" gorm:"default:0"` // 0: pending, 1: approved, 2: rejected
	Message string `json:"message" gorm:"size:500"`
}

type ClassCreateEditRequest struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ClassType int    `json:"classType"` // 1: 教学班 (default), 2: 行政班
}

type ClassQuery struct {
	ID          uint   `json:"id"`
	ClassType   int    `json:"classType"` // 0 = all, 1 = 教学班, 2 = 行政班
	AdminUserId uint   `json:"adminUserId"`
	InviteCode  string `json:"inviteCode"`
	Page
}

type ClassJoinRequestQuery struct {
	ClassId uint `json:"classId"`
	UserId  uint `json:"userId"`
	Status  *int `json:"status"`
	Page
}

const (
	ClassTeacherAppStatusPending  = 0
	ClassTeacherAppStatusApproved = 1
	ClassTeacherAppStatusRejected = 2
)

type ClassTeacherApplication struct {
	Model
	ClassId uint   `json:"classId" gorm:"index"`
	Class   *Class `json:"class,omitempty" gorm:"foreignKey:ClassId"`
	UserId  uint   `json:"userId" gorm:"index"`
	User    *User  `json:"user,omitempty" gorm:"foreignKey:UserId"`
	Status  int    `json:"status" gorm:"default:0"` // 0: pending, 1: approved, 2: rejected
	Message string `json:"message" gorm:"size:500"`
}

type ClassTeacherApplicationQuery struct {
	ClassId uint `json:"classId"`
	UserId  uint `json:"userId"`
	Status  *int `json:"status"`
	Page
}
