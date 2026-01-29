package repository

import (
"gorm.io/gorm"
)

// DB 全局数据库连接
var DB *gorm.DB

// 全局Repository实例
var (
UserRepo        IUserRepository
CourseRepo      ICourseRepository
PaperSeriesRepo IPaperSeriesRepository
PaperCodeRepo   IPaperCodeRepository
AdminRoleRepo   IAdminRoleRepository
AdminPermRepo   IAdminPermissionRepository
AttachmentRepo  IAttachmentRepository
AppConfigRepo   IAppConfigRepository
DocumentRepo    IDocumentRepository
ExamPaperRepo   IExamPaperRepository
PastPaperRepo   IPastPaperRepository
RandomPaperRepo IRandomPaperRepository
CourseVideoRepo ICourseVideoRepository
DocumentCategoryRepo IDocumentCategoryRepository
OrganisationRepo IOrganisationRepository
QualificationRepo IQualificationRepository
SyllabusRepo ISyllabusRepository
ChapterRepo IChapterRepository
QuestionRepo IQuestionRepository
GradeRepo IGradeRepository
SlideRepo ISlideRepository
MediaVideoRepo IMediaVideoRepository
MediaImageRepo IMediaImageRepository
VocabularySetRepo IVocabularySetRepository
VocabularyItemRepo IVocabularyItemRepository
VerificationRepo IVerificationRepository
MCPTokenRepo IMCPTokenRepository
)

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
SyllabusRepo = NewSyllabusRepository(db)
ChapterRepo = NewChapterRepository(db)
QuestionRepo = NewQuestionRepository(db)
GradeRepo = NewGradeRepository(db)
SlideRepo = NewSlideRepository(db)
MediaVideoRepo = NewMediaVideoRepository(db)
MediaImageRepo = NewMediaImageRepository(db)
VocabularySetRepo = NewVocabularySetRepository(db)
VocabularyItemRepo = NewVocabularyItemRepository(db)
VerificationRepo = NewVerificationRepository(db)
MCPTokenRepo = NewMCPTokenRepository(db)
// ... 其他repository按需添加
}
