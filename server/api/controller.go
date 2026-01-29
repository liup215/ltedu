package backend

import (
	"edu/conf"
	"edu/lib/net/http/middleware/auth"
	v1 "edu/server/api/v1"
	"edu/server/mcp"
	"fmt"
	"os"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func NewHandler() *Handler {
	authMiddleware, err := auth.New(conf.Conf.Auth)
	if err != nil {
		fmt.Println("Auth模块初始化失败！")
		os.Exit(0)
	}
	authMiddleware.Authenticator = v1.AuthCtrl.Authenticator
	authMiddleware.Authorizator = v1.AuthCtrl.Authorizator
	authMiddleware.PayloadFunc = v1.AuthCtrl.PayloadFunc
	authMiddleware.LoginResponse = v1.AuthCtrl.LoginResponse
	return &Handler{
		authMiddleware: authMiddleware,
	}
}

type Handler struct {
	authMiddleware *jwt.GinJWTMiddleware
}

func (h *Handler) Route(r *gin.RouterGroup) {
	h.noAuthRout(r)
	h.authRout(r)
}

func (h *Handler) noAuthRout(r *gin.RouterGroup) {
	r.POST("/v1/login", h.authMiddleware.LoginHandler)
	r.POST("/v1/captcha", v1.CaptchaCtrl.GetImage)

	r.POST("/v1/verification/send-code", v1.VerificationCtrl.SendCode)

	r.GET("/v1/view/page/:hash/*page", v1.AttachmentCtrl.ViewDocumentPages)
	r.GET("/v1/view/cover/:hash", v1.AttachmentCtrl.ViewDocumentCover)
	r.GET("/v1/download/:jwt", v1.AttachmentCtrl.DownloadDocument)

	r.POST("/v1/register", v1.AuthCtrl.Register) // Add registration route

	// MCP endpoint (no auth, uses token in query parameter)
	// Stateless HTTP transport (JSON-RPC 2.0 over HTTP)
	r.POST("/mcp", mcp.HandleHTTP)

	r.POST("/v1/mediaImage/list", v1.MediaImageCtrl.SelectImageList)
	r.POST("/v1/mediaImage/byId", v1.MediaImageCtrl.SelectImageById)

	r.POST("/v1/mediaVideo/list", v1.MediaVideoCtrl.SelectVideoList)
	r.POST("/v1/mediaVideo/byId", v1.MediaVideoCtrl.SelectVideoById)

	r.POST("/v1/organisation/list", v1.QualificationCtrl.SelectOrganisationList)
	r.POST("/v1/organisation/byId", v1.QualificationCtrl.SelectOrganisationById)
	r.POST("/v1/organisation/all", v1.QualificationCtrl.SelectOrganisationAll)

	r.POST("/v1/qualification/list", v1.QualificationCtrl.SelectQualificationList)
	r.POST("/v1/qualification/byId", v1.QualificationCtrl.SelectQualificationById)
	r.POST("/v1/qualification/all", v1.QualificationCtrl.SelectQualificationAll)

	r.POST("/v1/syllabus/list", v1.QualificationCtrl.SelectSyllabusList)
	r.POST("/v1/syllabus/byId", v1.QualificationCtrl.SelectSyllabusById)
	r.POST("/v1/syllabus/all", v1.QualificationCtrl.SelectSyllabusAll)

	r.POST("/v1/chapter/tree", v1.QualificationCtrl.GetChapterTree)
	r.POST("/v1/chapter/byId", v1.QualificationCtrl.SelectChapterById)

	r.POST("/v1/pastPaper/series/getById", v1.PaperCtrl.SelectSeriesById)
	r.POST("/v1/pastPaper/series/list", v1.PaperCtrl.SelectSeriesList)
	r.POST("/v1/pastPaper/series/all", v1.PaperCtrl.SelectSeriesAll)

	r.POST("/v1/pastPaper/code/getById", v1.PaperCtrl.SelectCodeById)
	r.POST("/v1/pastPaper/code/list", v1.PaperCtrl.SelectCodeList)
	r.POST("/v1/pastPaper/code/all", v1.PaperCtrl.SelectCodeAll)

	r.POST("/v1/paper/past/getById", v1.PaperCtrl.SelectPastPaperById)
	r.POST("/v1/paper/past/list", v1.PaperCtrl.SelectPastPaperList)
	r.POST("/v1/paper/past/all", v1.PaperCtrl.SelectPastPaperAll)

	r.POST("/v1/paper/random/getById", v1.PaperCtrl.SelectRandomPaperById)
	r.POST("/v1/paper/random/list", v1.PaperCtrl.SelectRandomPaperList)

	r.POST("/v1/question/list", v1.QuestionCtrl.SelectQuestionList)
	r.POST("/v1/question/byId", v1.QuestionCtrl.SelectQuestionById)
	r.POST("/v1/question/all", v1.QuestionCtrl.SelectQuestionAll)

	r.POST("/v1/vocabularySet/list", v1.VocabularyCtrl.SelectVocabularySetList)
	r.POST("/v1/vocabularySet/byId", v1.VocabularyCtrl.SelectVocabularySetById)

	r.POST("/v1/documentCategory/list", v1.DocumentCategoryCtrl.SelectCategoryList)
	r.POST("/v1/documentCategory/byId", v1.DocumentCategoryCtrl.SelectCategoryById)
	r.POST("/v1/documentCategory/all", v1.DocumentCategoryCtrl.SelectCategoryAll)

	r.POST("/v1/document/list", v1.DocumentCtrl.SelectDocumentList)
	r.POST("/v1/document/byId", v1.DocumentCtrl.SelectDocumentById)
	r.POST("/v1/document/all", v1.DocumentCtrl.SelectDocumentAll)

	r.POST("/v1/slide/link", v1.SlideCtrl.ViewSlide)
	r.POST("/v1/slide/getById", v1.SlideCtrl.SelectSlideById)
	r.POST("/v1/slide/list", v1.SlideCtrl.SelectSlideList)
	r.POST("/v1/slide/all", v1.SlideCtrl.SelectSlideAll)

	r.POST("/v1/course/list", v1.CourseCtrl.SelectCourseList)
	r.POST("/v1/course/byId", v1.CourseCtrl.SelectCourseById)

	r.POST("/v1/courseVideo/list", v1.CourseVideoCtrl.SelectCourseVideoList)
	r.POST("/v1/courseVideo/byId", v1.CourseVideoCtrl.SelectCourseVideoById)

	r.POST("/v1/practice/quick", v1.PracticeCtrl.QuickPractice)
	r.POST("/v1/practice/paper", v1.PracticeCtrl.PaperPractice)

}

func (h *Handler) authRout(r *gin.RouterGroup) {
	r.Use(h.authMiddleware.MiddlewareFunc())
	r.GET("/v1/user", v1.UserCtrl.User)

	r.GET("/v1/addons", v1.AddonCtrl.Index)
	r.GET("/v1/dashboard", v1.DashboardCtrl.Index)

	r.GET("/v1/statistic/userRegister", v1.StatisticCtrl.UserRegister)
	r.GET("/v1/statistic/orderCreated", v1.StatisticCtrl.OrderCreated)
	r.GET("/v1/statistic/orderPaidCount", v1.StatisticCtrl.OrderPaidCount)
	r.GET("/v1/statistic/orderPaidSum", v1.StatisticCtrl.OrderPaidSum)

	r.GET("/v1/syssetting/imageUpload", v1.SettingCtrl.GetImageUploadConfig)
	r.POST("/v1/syssetting/imageUpload", v1.SettingCtrl.SaveImageUploadConfig)
	r.GET("/v1/syssetting/videoUpload", v1.SettingCtrl.GetVideoUploadConfig)
	r.POST("/v1/syssetting/videoUpload", v1.SettingCtrl.SaveVideoUploadConfig)
	r.GET("/v1/syssetting/webSite", v1.SettingCtrl.GetWebSiteConfig)
	r.POST("/v1/syssetting/webSite", v1.SettingCtrl.SaveWebSiteConfig)

	r.GET("/v1/mediaImage/disk", v1.MediaImageCtrl.UploadDisk)
	r.POST("/v1/mediaImage/uploadToDisk", v1.MediaImageCtrl.UploadImageToDisk)
	r.POST("/v1/mediaImage/token/qiniu", v1.MediaImageCtrl.QiniuUploadToken)
	r.POST("/v1/mediaImage/create", v1.MediaImageCtrl.CreateImage)

	r.GET("/v1/mediaVideo/disk", v1.MediaVideoCtrl.UploadDisk)
	r.POST("/v1/mediaVideo/uploadToDisk", v1.MediaVideoCtrl.UploadVideoToDisk)
	r.POST("/v1/mediaVideo/token/qiniu", v1.MediaVideoCtrl.QiniuUploadToken)
	r.POST("/v1/mediaVideo/create", v1.MediaVideoCtrl.CreateVideo)

	r.POST("/v1/organisation/create", v1.QualificationCtrl.CreateOrganisation)
	r.POST("/v1/organisation/edit", v1.QualificationCtrl.EditOrganisation)
	r.POST("/v1/organisation/delete", v1.QualificationCtrl.DeleteOrganisation)

	r.POST("/v1/qualification/create", v1.QualificationCtrl.CreateQualification)
	r.POST("/v1/qualification/edit", v1.QualificationCtrl.EditQualification)
	r.POST("/v1/qualification/delete", v1.QualificationCtrl.EditQualification)

	r.POST("/v1/syllabus/create", v1.QualificationCtrl.CreateSyllabus)
	r.POST("/v1/syllabus/edit", v1.QualificationCtrl.EditSyllabus)
	r.POST("/v1/syllabus/delete", v1.QualificationCtrl.DeleteSyllabus)

	r.POST("/v1/paper/past/upload", v1.PaperCtrl.UploadPastPaper)
	r.POST("/v1/paper/past/create", v1.PaperCtrl.CreatePastPaper)
	r.POST("/v1/paper/past/edit", v1.PaperCtrl.EditPastPaper)

	// r.POST("/v1/paper/past/question/update", v1.PaperCtrl.UpdatePastPaperQuestion)
	// r.POST("/v1/paper/past/question/add", v1.PaperCtrl.AddPastPaperQuestion)
	// r.POST("/v1/paper/past/question/delete", v1.PaperCtrl.DeletePastPaperQuestion)

	r.POST("/v1/paper/random/create", v1.PaperCtrl.CreateRandomPaper)

	r.POST("/v1/paper/exam/byId", v1.PaperCtrl.SelectExamPaperById)
	r.POST("/v1/paper/exam/list", v1.PaperCtrl.SelectExamPaperList)
	r.POST("/v1/paper/exam/all", v1.PaperCtrl.SelectExamPaperAll)

	r.POST("/v1/paper/exam/upload", v1.PaperCtrl.UploadExamPaper)
	r.POST("/v1/paper/exam/question/update", v1.PaperCtrl.UpdateExamPaperQuestion)
	r.POST("/v1/paper/exam/create", v1.PaperCtrl.CreateExamPaper)
	r.POST("/v1/paper/exam/edit", v1.PaperCtrl.EditExamPaper)

	r.POST("/v1/question/create", v1.QuestionCtrl.CreateQuestion)
	r.POST("/v1/question/edit", v1.QuestionCtrl.EditQuestion)
	r.POST("/v1/question/delete", v1.QuestionCtrl.DeleteQuestion)
	r.POST("/v1/question/addChapter", v1.QuestionCtrl.AddQuestionChapter)
	r.POST("/v1/question/deleteChapter", v1.QuestionCtrl.DeleteQuestionChapter)

	r.POST("/v1/vocabularySet/create", v1.VocabularyCtrl.CreateVocabularySet)
	r.POST("/v1/vocabularySet/delete", v1.VocabularyCtrl.DeleteVocabularySet)
	r.POST("/v1/vocabularySet/edit", v1.VocabularyCtrl.EditVocabularySet)

	r.POST("/v1/vocabularyItem/insert", v1.VocabularyCtrl.InsertVocabularyItem)
	r.POST("/v1/vocabularyItem/update", v1.VocabularyCtrl.UpdateVocabularyItem)
	r.POST("/v1/vocabularyItem/delete", v1.VocabularyCtrl.DeleteVocabularyItem)

	r.POST("/v1/documentCategory/create", v1.DocumentCategoryCtrl.CreateCategory)
	r.POST("/v1/documentCategory/edit", v1.DocumentCategoryCtrl.EditCategory)
	r.POST("/v1/documentCategory/delete", v1.DocumentCategoryCtrl.DeleteCategory)

	r.POST("/v1/document/create", v1.DocumentCtrl.CreateDocument)
	r.POST("/v1/document/edit", v1.DocumentCtrl.EditDocument)
	r.POST("/v1/document/delete", v1.DocumentCtrl.DeleteDocument)
	r.POST("/v1/document/download", v1.DocumentCtrl.DownloadDocument)
	r.POST("/v1/upload/document", v1.AttachmentCtrl.UploadDocument)

	r.POST("/v1/slide/create", v1.SlideCtrl.CreateSlide)
	r.POST("/v1/slide/edit", v1.SlideCtrl.EditSlide)
	r.POST("/v1/slide/delete", v1.SlideCtrl.DeleteSlide)

	r.POST("/v1/course/create", v1.CourseCtrl.CreateCourse)
	r.POST("/v1/course/edit", v1.CourseCtrl.EditCourse)
	r.POST("/v1/course/delete", v1.CourseCtrl.DeleteCourse)

	r.POST("/v1/courseVideo/create", v1.CourseVideoCtrl.CreateCourseVideo)
	r.POST("/v1/courseVideo/edit", v1.CourseVideoCtrl.EditCourseVideo)
	r.POST("/v1/courseVideo/delete", v1.CourseVideoCtrl.DeleteCourseVideo)

	// Teacher Application routes
	r.POST("/v1/user/teacher/apply", v1.TeacherApplicationCtrl.Apply)                   // Submit application
	r.POST("/v1/user/teacher/application", v1.TeacherApplicationCtrl.GetByUser)         // Get own application
	r.POST("/v1/admin/teacher-applications/list", v1.TeacherApplicationCtrl.List)       // List applications (admin)
	r.POST("/v1/admin/teacher-applications/detail", v1.TeacherApplicationCtrl.Get)      // Get application details (admin)
	r.POST("/v1/admin/teacher-applications/approve", v1.TeacherApplicationCtrl.Approve) // Approve application (admin)
	r.POST("/v1/admin/teacher-applications/reject", v1.TeacherApplicationCtrl.Reject)   // Reject application (admin)

	r.POST("/v1/practice/grade", v1.PracticeCtrl.GradePractice)

	// 修改密码
	r.POST("/v1/change-password", v1.AuthCtrl.ChangePassword)

	// MCP Token管理
	r.POST("/v1/mcp/token/create", v1.MCPTokenCtrl.CreateToken)
	r.POST("/v1/mcp/token/list", v1.MCPTokenCtrl.ListTokens)
	r.POST("/v1/mcp/token/delete", v1.MCPTokenCtrl.DeleteToken)
	r.POST("/v1/mcp/token/deactivate", v1.MCPTokenCtrl.DeactivateToken)
	r.POST("/v1/mcp/token/activate", v1.MCPTokenCtrl.ActivateToken)

	// MCP Token管理 - 管理员端点
	r.POST("/v1/mcp/token/admin/list", v1.MCPTokenCtrl.AdminListTokens)
	r.POST("/v1/mcp/token/admin/delete", v1.MCPTokenCtrl.AdminDeleteToken)
	r.POST("/v1/mcp/token/admin/deactivate", v1.MCPTokenCtrl.AdminDeactivateToken)
	r.POST("/v1/mcp/token/admin/activate", v1.MCPTokenCtrl.AdminActivateToken)

	r.POST("/v1/chapter/create", v1.QualificationCtrl.CreateChapter)
	r.POST("/v1/chapter/edit", v1.QualificationCtrl.EditChapter)
	r.POST("/v1/chapter/delete", v1.QualificationCtrl.DeleteChapter)

	r.POST("/v1/user/list", v1.UserCtrl.SelectUserList)
	r.POST("/v1/user/byId", v1.UserCtrl.SelectUserById)
	r.POST("/v1/user/all", v1.UserCtrl.SelectUserAll)
	r.POST("/v1/user/create", v1.UserCtrl.CreateUser)
	r.POST("/v1/user/edit", v1.UserCtrl.EditUser)
	r.POST("/v1/user/delete", v1.UserCtrl.DeleteUser)
	r.POST("/v1/user/setAdmin", v1.UserCtrl.SetAdmin)
	r.POST("/v1/user/removeAdmin", v1.UserCtrl.RemoveAdmin)
	r.POST("/v1/user/vip", v1.UserCtrl.GrantVipMonth)

	r.POST("/v1/school/grade/list", v1.SchoolCtrl.SelectGradeList)
	r.POST("/v1/school/grade/byId", v1.SchoolCtrl.SelectGradeById)
	r.POST("/v1/school/grade/all", v1.SchoolCtrl.SelectGradeAll)
	r.POST("/v1/school/grade/create", v1.SchoolCtrl.CreateGrade)
	r.POST("/v1/school/grade/edit", v1.SchoolCtrl.EditGrade)
	r.POST("/v1/school/grade/delete", v1.SchoolCtrl.DeleteGrade)

	r.POST("/v1/school/classType/list", v1.SchoolCtrl.SelectClassTypeList)
	r.POST("/v1/school/classType/byId", v1.SchoolCtrl.SelectClassTypeById)
	r.POST("/v1/school/classType/all", v1.SchoolCtrl.SelectClassTypeAll)
	r.POST("/v1/school/classType/create", v1.SchoolCtrl.CreateClassType)
	r.POST("/v1/school/classType/edit", v1.SchoolCtrl.EditClassType)
	r.POST("/v1/school/classType/delete", v1.SchoolCtrl.DeleteClassType)

	r.POST("/v1/school/class/list", v1.SchoolCtrl.SelectClassList)
	r.POST("/v1/school/class/byId", v1.SchoolCtrl.SelectClassById)
	r.POST("/v1/school/class/all", v1.SchoolCtrl.SelectClassAll)
	r.POST("/v1/school/class/create", v1.SchoolCtrl.CreateClass)
	r.POST("/v1/school/class/edit", v1.SchoolCtrl.EditClass)
	r.POST("/v1/school/class/delete", v1.SchoolCtrl.DeleteClass)
	r.POST("/v1/school/class/studentList", v1.SchoolCtrl.GetStudentsByClassId)
	r.POST("/v1/school/class/addStudent", v1.SchoolCtrl.AddStudentToClass)
	r.POST("/v1/school/class/deleteStudent", v1.SchoolCtrl.DeleteStudentFromClass)

	r.POST("/v1/pastPaper/series/create", v1.PaperCtrl.CreateSeries)
	r.POST("/v1/pastPaper/series/edit", v1.PaperCtrl.EditSeries)

	r.POST("/v1/pastPaper/code/create", v1.PaperCtrl.CreateCode)
	r.POST("/v1/pastPaper/code/edit", v1.PaperCtrl.EditCode)
}
