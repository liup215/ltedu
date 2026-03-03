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
		// Migrate chapter-exam_node relationship from many2many join table to FK on Chapter.
		// This is idempotent: if exam_node_chapters doesn't exist, the Exec is a no-op error.
		if result := db.Exec(`
			UPDATE chapters c
			JOIN (
				SELECT chapter_id, MIN(syllabus_exam_node_id) AS exam_node_id
				FROM exam_node_chapters
				GROUP BY chapter_id
			) t ON c.id = t.chapter_id
			SET c.exam_node_id = t.exam_node_id
			WHERE c.exam_node_id = 0 OR c.exam_node_id IS NULL
		`); result.Error != nil {
			_ = result.Error // non-fatal: table may not exist (fresh install) or already migrated
		}
		if result := db.Exec("DROP TABLE IF EXISTS exam_node_chapters"); result.Error != nil {
			_ = result.Error // non-fatal: backfill failure does not prevent startup
		}

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
	// Migrate legacy is_admin=true users to the RBAC "admin" role (one-time idempotent migration)
	migrateAdminFlagToRBAC()
}

// migrateAdminFlagToRBAC finds users with the old is_admin=true DB column and assigns
// them the "admin" RBAC role if they don't already have it. Idempotent.
func migrateAdminFlagToRBAC() {
	adminRole, err := repository.AdminRoleRepo.FindBySlug("admin")
	if err != nil || adminRole == nil {
		fmt.Println("Warning: could not find 'admin' role for migration")
		return
	}
	db := repository.GetDB()
	if db == nil {
		return
	}
	// Check if the legacy is_admin column still exists before querying it
	if !db.Migrator().HasColumn(&struct{ IsAdmin bool `gorm:"column:is_admin"` }{}, "is_admin") {
		return // Column already removed, migration not needed
	}
	// Query users that still have is_admin = true in the DB column
	var legacyAdminIDs []uint
	if err := db.Raw("SELECT id FROM users WHERE is_admin = true").Scan(&legacyAdminIDs).Error; err != nil {
		fmt.Println("Warning: failed to query legacy admin users:", err)
		return
	}
	for _, userID := range legacyAdminIDs {
		has, err := repository.UserRoleRepo.HasRole(userID, adminRole.ID)
		if err != nil || has {
			continue
		}
		if err := repository.UserRoleRepo.AssignRole(userID, adminRole.ID); err != nil {
			fmt.Printf("Warning: failed to migrate user %d to admin role: %v\n", userID, err)
		}
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
