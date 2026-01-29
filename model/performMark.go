package model

const (
	MARK_TYPE_PERFORM        = 1
	MARK_TYPE_PERFORM_NAME   = "操行分"
	MARK_TYPE_ACTIVITY       = 2
	MARK_TYPE_ACTIVITY_NAME  = "活动分"
	MARK_TYPE_ACTIVITY_LIMIT = 50
)

// type UserPerformMark struct {
// 	Model
// 	UserId     uint     `json:"userId"`
// 	User       User     `json:"user"`
// 	MarkTypeId uint     `json:"markTypeId"`
// 	MarkType   MarkType `json:"markType"`
// 	FinalMark  int      `json:"finalMark"`
// }

type PerformMarkRecord struct {
	Model
	UserId                   uint   `json:"userId"`                   // 学生ID
	UserRealname             string `json:"userRealname"`             // 学生姓名
	UserEngname              string `json:"userEngname"`              // 学生英文名
	Sex                      uint   `json:"sex"`                      // 性别
	SexLabel                 string `json:"sexLabel"`                 // 性别
	MainClassId              uint   `json:"mainClasaId"`              // 学生行政班ID
	MainClassName            string `json:"mainClassName"`            // 学生行政班名称
	GradeId                  uint   `json:"gradeId"`                  // 学生年级
	GradeName                string `json:"gradeName"`                // 学生年级名称
	TeacherType              uint   `json:"teacherType"`              // 添加人员类型
	HeadTeacherId            uint   `json:"headTeacherId"`            // 学生班主任ID
	HeadTeacherRealname      string `json:"headTeacherRealname"`      // 学生班主任姓名
	GradeLeadTeacherId       uint   `json:"gradeLeadTeacherId"`       // 学生年级组长ID
	GradeLeadTeacherRealname string `json:"gradeLeadTeacherRealname"` // 学生年级组长姓名
	TeacherTypeName          string `json:"teacherTypeName"`          // 添加人员类型名称
	TeacherId                uint   `json:"teacherId"`                // 添加人员ID
	TeacherRealname          string `json:"teacherRealname"`          // 添加人员真实姓名
	ClassId                  uint   `json:"classId"`                  // 添加班级
	ClassName                string `json:"className"`                // 添加选修班级名称
	MarkTypeId               uint   `json:"markTypeId"`               // 分数类型
	MarkTypeName             string `json:"markTypeName"`             // 分数类型名称
	PerformTypeId            uint   `json:"performTypeId"`            // 行为类别:教学、清洁、校服、考勤、寝室管理等
	PerformTypeName          string `json:"performTypeName"`          // 行为名称
	Description              string `json:"description"`              // 描述
	Mark                     int    `json:"mark"`                     // 加减分数（减分用负数）
	FinalMark                int    `json:"finalMark"`                // 记录余额
}

type PerformMarkRecordQuery struct {
	ID         uint   `json:"id"`
	UserId     uint   `json:"userId"`
	MarkTypeId uint   `json:"markTypeId"`
	StartAt    string `json:"startAt"`
	EndAt      string `json:"endAt"`
	Page
}

type PerformMarkRecordBatchAddRequest struct {
	UserIds         []uint `json:"userIds"`
	ClassId         uint   `json:"classId"`
	MarkTypeId      uint   `json:"markTypeId"`
	Mark            int    `json:"mark"`
	Description     string `json:"description"`
	TeacherType     uint   `json:"teacherType"`
	TeacherId       uint   `json:"-"`
	TeacherRealname string `json:"-"`
	PerformTypeId   uint   `json:"performTypeId"`
	PerformTypeName string `json:"performTypeName"`
}
