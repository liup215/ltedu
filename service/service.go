package service

import (
	"edu/conf"
	"edu/lib/database/orm"
	"edu/model"
	"edu/repository"
	"fmt"
	"sync"
	// bg "github.com/dgraph-io/badger/v4"
)

var once = &sync.Once{}

func setDB() {
	once.Do(func() {
		db := orm.New(conf.Conf.Orm)

		db.AutoMigrate(&model.AdminPermission{})
		db.AutoMigrate(&model.AdminRole{})
		// admin_role_permissions many2many join table (Role <-> Permission)
		db.SetupJoinTable(&model.AdminRole{}, "Permissions", &model.AdminRolePermission{})
		db.AutoMigrate(&model.AdminRolePermission{})
		// user_roles many2many join table (User <-> Role)
		db.SetupJoinTable(&model.User{}, "Roles", &model.UserRole{})
		db.AutoMigrate(&model.UserRole{})
		db.AutoMigrate(&model.AdminLog{})
		db.AutoMigrate(&model.User{})
		db.AutoMigrate(&model.Order{})
		db.AutoMigrate(&model.AppConfig{})
		db.AutoMigrate(&model.MediaImage{})
		db.AutoMigrate(&model.Organisation{})
		db.AutoMigrate(&model.Qualification{})
		db.AutoMigrate(&model.Syllabus{})
		db.AutoMigrate(&model.PaperCode{})
		db.AutoMigrate(&model.PaperSeries{})
		db.AutoMigrate(&model.PastPaper{})
		db.AutoMigrate(&model.Chapter{})
		db.AutoMigrate(&model.VocabularyItem{})
		db.AutoMigrate(&model.VocabularySet{})
		// Register custom join table for Class.Students many2many so that the
		// user_class_relation table includes the status column.
		db.SetupJoinTable(&model.Class{}, "Students", &model.UserClassRelation{})
		db.SetupJoinTable(&model.User{}, "Classes", &model.UserClassRelation{})
		db.AutoMigrate(&model.UserClassRelation{})
		// Back-fill status for existing rows that were created before the status column existed.
		if result := db.Exec("UPDATE user_class_relation SET status = ? WHERE status IS NULL OR status = 0", model.ClassStudentStatusStudying); result.Error != nil {
			_ = result.Error // non-fatal: backfill failure does not prevent startup
		}
		db.AutoMigrate(&model.Class{})
		db.AutoMigrate(&model.ClassJoinRequest{})
		db.AutoMigrate(&model.ClassTeacherApplication{})
		db.AutoMigrate(&model.Teacher{})
		db.AutoMigrate(&model.Grade{})
		db.AutoMigrate(&model.ClassType{})
		db.AutoMigrate(&model.PerformMarkRecord{})
		db.AutoMigrate(&model.Question{})
		db.AutoMigrate(&model.MediaImage{})
		db.AutoMigrate(&model.Document{})
		db.AutoMigrate(&model.DocumentCategory{})
		db.AutoMigrate(&model.Attachment{})
		db.AutoMigrate(&model.Slide{})
		db.AutoMigrate(&model.RandomPaper{})
		db.AutoMigrate(&model.ExamPaper{})
		db.AutoMigrate(&model.QuestionRandomPapers{})

		db.AutoMigrate(&model.Course{})
		db.AutoMigrate(&model.CourseUserRecord{})
		db.AutoMigrate(&model.MediaVideo{})
		db.AutoMigrate(&model.CourseVideo{})
		db.AutoMigrate(&model.TeacherApplication{})
		db.AutoMigrate(&model.Verification{})
		db.AutoMigrate(&model.MCPToken{})

		// Syllabus Navigator models
		db.AutoMigrate(&model.Goal{})
		db.AutoMigrate(&model.KnowledgeState{})
		db.AutoMigrate(&model.Task{})
		db.AutoMigrate(&model.Attempt{})
		db.AutoMigrate(&model.TaskLog{})
		db.AutoMigrate(&model.KnowledgePoint{})

		// Learning Plan models
		db.AutoMigrate(&model.StudentLearningPlan{})
		db.AutoMigrate(&model.StudentLearningPlanVersion{})

		// Syllabus Exam Node model
		db.AutoMigrate(&model.SyllabusExamNode{})

		// Phase Plan model
		db.AutoMigrate(&model.LearningPhasePlan{})

		// 初始化Repository层
		repository.InitRepositories(db)

		// baseSvr.badgerDB = badger.New(conf.Conf.Badger)
	})
}

// SeedRBACDefaults seeds default RBAC roles and permissions. Called after full service init.
func SeedRBACDefaults() {
	if err := AdminSvr.SeedDefaultRolesAndPermissions(); err != nil {
		fmt.Println("Warning: failed to seed default roles/permissions:", err)
	}
}

var baseSvr = baseService{}

func newBaseService() baseService {
	setDB()
	return baseSvr
}

func CloseService() {
	// baseSvr.badgerDB.Close()
	fmt.Println("Badger DB closed")
}

type baseService struct {
	// db *gorm.DB
	// badgerDB *bg.DB
}
