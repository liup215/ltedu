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
	Name         string  `json:"name"`
	InviteCode   string  `json:"inviteCode" gorm:"uniqueIndex;size:32"`
	AdminUserId  uint    `json:"adminUserId" gorm:"index"`
	AdminUser    *User   `json:"adminUser,omitempty" gorm:"foreignKey:AdminUserId"`
	Students     []*User `json:"students" gorm:"many2many:user_class_relation;"`
}

const (
	ClassJoinStatusPending  = 0
	ClassJoinStatusApproved = 1
	ClassJoinStatusRejected = 2
)

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
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ClassQuery struct {
	ID          uint   `json:"id"`
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
