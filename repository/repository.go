package repository

import (
	"gorm.io/gorm"
)

// DB 全局数据库连接
var DB *gorm.DB

// 全局Repository实例
var (
	UserRepo             IUserRepository
	CourseRepo           ICourseRepository
	PaperSeriesRepo      IPaperSeriesRepository
	PaperCodeRepo        IPaperCodeRepository
	AdminRoleRepo        IAdminRoleRepository
	AdminPermRepo        IAdminPermissionRepository
	UserRoleRepo         IUserRoleRepository
	AttachmentRepo       IAttachmentRepository
	AppConfigRepo        IAppConfigRepository
	DocumentRepo         IDocumentRepository
	ExamPaperRepo        IExamPaperRepository
	PastPaperRepo        IPastPaperRepository
	RandomPaperRepo      IRandomPaperRepository
	CourseVideoRepo      ICourseVideoRepository
	DocumentCategoryRepo IDocumentCategoryRepository
	OrganisationRepo     IOrganisationRepository
	QualificationRepo    IQualificationRepository
	SyllabusRepo         ISyllabusRepository
	ChapterRepo          IChapterRepository
	QuestionRepo         IQuestionRepository
	GradeRepo            IGradeRepository
	SlideRepo            ISlideRepository
	MediaVideoRepo       IMediaVideoRepository
	MediaImageRepo       IMediaImageRepository
	VocabularySetRepo    IVocabularySetRepository
	VocabularyItemRepo   IVocabularyItemRepository
	VerificationRepo     IVerificationRepository
	MCPTokenRepo         IMCPTokenRepository
	GoalRepo             IGoalRepository
	KnowledgeStateRepo   IKnowledgeStateRepository
	TaskRepo             ITaskRepository
	AttemptRepo          IAttemptRepository
	TaskLogRepo          ITaskLogRepository
	KnowledgePointRepo   IKnowledgePointRepository
	ClassRepo                    IClassRepository
	ClassJoinRequestRepo         IClassJoinRequestRepository
	StudentLearningPlanRepo      IStudentLearningPlanRepository
	ExamNodeRepo                 IExamNodeRepository
	PhasePlanRepo                IPhasePlanRepository
	ClassTeacherApplicationRepo  IClassTeacherApplicationRepository
	ConversationRepo             IConversationRepository
	ConversationSessionRepo      IConversationSessionRepository
	ConversationMessageRepo      IConversationMessageRepository
	NLUFeedbackRepo              INLUFeedbackRepository
	AuditLogRepo                 IAuditLogRepository
	FeedbackRepo                 IFeedbackRepository
	BlogPostRepo                  IBlogPostRepository
)

func GetTableName(db *gorm.DB, model interface{}) string {
	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(model); err != nil || stmt.Schema == nil {
		return ""
	}
	return stmt.Schema.Table
}

// GetDB returns the global database connection.
func GetDB() *gorm.DB {
	return DB
}

// InitRepositories 初始化所有仓储
func InitRepositories(db *gorm.DB) {
	DB = db

	// 初始化各个Repository
	UserRepo = NewUserRepository(db)
	CourseRepo = NewCourseRepository(db)
	PaperSeriesRepo = NewPaperSeriesRepository(db)
	PaperCodeRepo = NewPaperCodeRepository(db)
	AdminRoleRepo = NewAdminRoleRepository(db)
	AdminPermRepo = NewAdminPermissionRepository(db)
	UserRoleRepo = NewUserRoleRepository(db)
	AttachmentRepo = NewAttachmentRepository(db)
	AppConfigRepo = NewAppConfigRepository(db)
	DocumentRepo = NewDocumentRepository(db)
	ExamPaperRepo = NewExamPaperRepository(db)
	PastPaperRepo = NewPastPaperRepository(db)
	RandomPaperRepo = NewRandomPaperRepository(db)
	CourseVideoRepo = NewCourseVideoRepository(db)
	DocumentCategoryRepo = NewDocumentCategoryRepository(db)
	OrganisationRepo = NewOrganisationRepository(db)
	QualificationRepo = NewQualificationRepository(db)
	SyllabusRepo = NewCachedSyllabusRepository(NewSyllabusRepository(db))
	ChapterRepo = NewCachedChapterRepository(NewChapterRepository(db))
	QuestionRepo = NewQuestionRepository(db)
	GradeRepo = NewGradeRepository(db)
	SlideRepo = NewSlideRepository(db)
	MediaVideoRepo = NewMediaVideoRepository(db)
	MediaImageRepo = NewMediaImageRepository(db)
	VocabularySetRepo = NewVocabularySetRepository(db)
	VocabularyItemRepo = NewVocabularyItemRepository(db)
	VerificationRepo = NewVerificationRepository(db)
	MCPTokenRepo = NewMCPTokenRepository(db)
	GoalRepo = NewGoalRepository(db)
	KnowledgeStateRepo = NewKnowledgeStateRepository(db, ChapterRepo)
	TaskRepo = NewTaskRepository(db)
	AttemptRepo = NewAttemptRepository(db)
	TaskLogRepo = NewTaskLogRepository(db)
	KnowledgePointRepo = NewKnowledgePointRepository(db)
	ClassRepo = NewClassRepository(db)
	ClassJoinRequestRepo = NewClassJoinRequestRepository(db)
	StudentLearningPlanRepo = NewStudentLearningPlanRepository(db)
	ExamNodeRepo = NewExamNodeRepository(db)
	PhasePlanRepo = NewPhasePlanRepository(db)
	ClassTeacherApplicationRepo = NewClassTeacherApplicationRepository(db)
	ConversationRepo = NewConversationRepository(db)
	ConversationSessionRepo = NewConversationSessionRepository(db)
	ConversationMessageRepo = NewConversationMessageRepository(db)
	NLUFeedbackRepo = NewNLUFeedbackRepository(db)
	AuditLogRepo = NewAuditLogRepository(db)
	FeedbackRepo = NewFeedbackRepository(db)
	BlogPostRepo = NewBlogPostRepository(db)
	// ... 其他repository按需添加
}
