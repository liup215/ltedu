import Foundation

/// Base URL for the LTEdu API. Override via the `LTEDU_API_BASE_URL` key in Info.plist
/// or the `LTEduAPIBaseURL` environment variable.
enum APIConfig {
    static var baseURL: String {
        if let envURL = ProcessInfo.processInfo.environment["LTEduAPIBaseURL"] {
            return envURL
        }
        if let plistURL = Bundle.main.object(forInfoDictionaryKey: "LTEduAPIBaseURL") as? String,
           !plistURL.isEmpty {
            return plistURL
        }
        return "http://localhost:80/api"
    }

    static var apiV1: String { "\(baseURL)/v1" }

    /// JWT token expiry duration in seconds (30 days)
    static let tokenExpiryDuration: TimeInterval = 30 * 24 * 60 * 60

    /// Default page size for paginated requests
    static let defaultPageSize = 20

    /// Timeout interval for network requests (seconds)
    static let requestTimeout: TimeInterval = 30
}

// MARK: - API Endpoints

enum APIEndpoint {
    // MARK: - Auth (no token required)
    static let captcha = "/captcha"
    static let login = "/login"
    static let register = "/register"
    static let sendVerificationCode = "/verification/send-code"

    // MARK: - User (token required)
    static let currentUser = "/user"
    static let updateAccount = "/account/update"
    static let changePassword = "/change-password"
    static let applyTeacher = "/teacher/apply"

    // MARK: - Courses (public)
    static let courseList = "/course/list"
    static let courseById = "/course/byId"
    static let courseVideoList = "/courseVideo/list"
    static let courseVideoById = "/courseVideo/byId"

    // MARK: - Papers & Questions (public)
    static let pastPaperList = "/paper/past/list"
    static let pastPaperById = "/paper/past/getById"
    static let randomPaperList = "/paper/random/list"
    static let questionList = "/question/list"
    static let questionById = "/question/byId"

    // MARK: - Classes (token required)
    static let classList = "/school/class/list"
    static let classById = "/school/class/byId"
    static let classStudentList = "/school/class/studentList"
    static let classApply = "/school/class/apply"
    static let classJoinRequestList = "/school/class/joinRequest/list"
    static let classJoinRequestApprove = "/school/class/joinRequest/approve"

    // MARK: - Learning Plans (token required)
    static let learningPlanList = "/learning-plan/list"
    static let learningPlanById = "/learning-plan/byId"
    static let learningPlanCreate = "/learning-plan/create"
    static let learningPlanEdit = "/learning-plan/edit"
    static let learningPlanGenerateTemplate = "/learning-plan/generateTemplate"

    // MARK: - Knowledge State (token required)
    static let knowledgeStateByChapter = "/knowledge-state/byChapter"
    static let knowledgeStateProgress = "/knowledge-state/progress"
    static let knowledgeStateDueReview = "/knowledge-state/due-review"

    // MARK: - Practice & Attempts (token required)
    static let practiceQuick = "/practice/quick"
    static let practicePaper = "/practice/paper"
    static let practiceGrade = "/practice/grade"
    static let attemptCreate = "/attempt/create"
    static let attemptStats = "/attempt/stats"

    // MARK: - Analytics (token required)
    static let analyticsClassSummary = "/analytics/class/summary"
    static let analyticsStudentList = "/analytics/student/list"
    static let analyticsHeatmap = "/analytics/heatmap"
    static let analyticsTrends = "/analytics/trends"
    static let analyticsEarlyWarning = "/analytics/earlyWarning"

    // MARK: - Syllabus (token required)
    static let syllabusList = "/syllabus/list"
    static let syllabusById = "/syllabus/byId"
    static let examNodeList = "/exam-node/list"
    static let examNodeById = "/exam-node/byId"
}
