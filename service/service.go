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
		db.AutoMigrate(&model.Class{})
		db.AutoMigrate(&model.ClassJoinRequest{})
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

		// 初始化Repository层
		repository.InitRepositories(db)

		// baseSvr.badgerDB = badger.New(conf.Conf.Badger)
	})
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
