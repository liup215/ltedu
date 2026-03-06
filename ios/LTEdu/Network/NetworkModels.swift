import Foundation

// MARK: - Generic API Response Wrapper

struct APIResponse<T: Decodable>: Decodable {
    let code: Int
    let message: String
    let data: T?

    var isSuccess: Bool { code == 0 }
}

struct APIResponseEmpty: Decodable {
    let code: Int
    let message: String

    var isSuccess: Bool { code == 0 }
}

// MARK: - Paginated Response

struct PaginatedData<T: Decodable>: Decodable {
    let list: [T]
    let total: Int
}

// MARK: - Pagination Request

struct PaginationRequest: Encodable {
    let pageIndex: Int
    let pageSize: Int

    init(pageIndex: Int = 1, pageSize: Int = APIConfig.defaultPageSize) {
        self.pageIndex = pageIndex
        self.pageSize = pageSize
    }
}

// MARK: - Auth Models

struct CaptchaResponse: Decodable {
    let captchaId: String
    let base64: String
}

struct LoginRequest: Encodable {
    let username: String
    let password: String
    let captchaId: String
    let captchaValue: String
}

struct LoginResponse: Decodable {
    let token: String
    let expire: String
}

struct RegisterRequest: Encodable {
    let username: String
    let email: String
    let mobile: String?
    let password: String
    let passwordConfirm: String
    let verificationCode: String
}

struct SendVerificationRequest: Encodable {
    let email: String
}

struct ChangePasswordRequest: Encodable {
    let oldPassword: String
    let newPassword: String
}

// MARK: - User Models

struct User: Codable, Identifiable {
    let id: Int
    let username: String
    let email: String?
    let nickname: String?
    let realname: String?
    let engname: String?
    let mobile: String?
    let avatar: String?
    let sex: Int?
    let status: Int?
    let isAdmin: Bool?
    let isTeacher: Bool?
    let teacherApplyStatus: Int?
    let vipExpireAt: String?
    let createdAt: String?
    let updatedAt: String?

    var displayName: String {
        nickname ?? realname ?? username
    }

    var isVip: Bool {
        guard let expireStr = vipExpireAt,
              let expireDate = ISO8601DateFormatter().date(from: expireStr) else { return false }
        return expireDate > Date()
    }
}

struct UpdateAccountRequest: Encodable {
    let realname: String?
    let nickname: String?
    let engname: String?
    let sex: Int?
    let mobile: String?
}

// MARK: - Course Models

struct Course: Codable, Identifiable {
    let id: Int
    let title: String
    let slug: String?
    let thumb: String?
    let charge: Int?
    let isFree: Int?
    let isShow: Int?
    let shortDescription: String?
    let originalDesc: String?
    let renderDesc: String?
    let syllabusId: Int?
    let userCount: Int?
    let publishedAt: String?

    var isFreeAccess: Bool { isFree == 1 || charge == 0 }
}

struct CourseVideo: Codable, Identifiable {
    let id: Int
    let title: String
    let courseId: Int?
    let duration: Int?
    let videoUrl: String?
    let isFree: Int?
    let sortOrder: Int?
}

// MARK: - Class Models

struct SchoolClass: Codable, Identifiable {
    let id: Int
    let name: String
    let classTypeId: Int?
    let gradeId: Int?
    let createdAt: String?
    let updatedAt: String?
}

struct ClassJoinRequest: Codable, Identifiable {
    let id: Int
    let classId: Int
    let userId: Int
    let status: Int?
    let createdAt: String?
    let user: User?
}

// MARK: - Learning Plan Models

struct LearningPlan: Codable, Identifiable {
    let id: Int
    let classId: Int?
    let userId: Int?
    let planType: String?
    let content: String?
    let version: Int?
    let createdBy: Int?
    let isPersonal: Bool?
    let createdAt: String?
    let updatedAt: String?
}

struct CreateLearningPlanRequest: Encodable {
    let classId: Int
    let planType: String
    let content: String
    let isPersonal: Bool?
}

struct EditLearningPlanRequest: Encodable {
    let id: Int
    let content: String?
    let planType: String?
}

// MARK: - Question & Paper Models

struct Question: Codable, Identifiable {
    let id: Int
    let questionType: String?
    let difficultyLevel: Int?
    let syllabusId: Int?
    let chapterId: Int?
    let content: QuestionContent?
    let createdAt: String?
}

struct QuestionContent: Codable {
    let stem: String?
    let analyze: String?
    let answer: String?
}

struct Paper: Codable, Identifiable {
    let id: Int
    let name: String?
    let paperCode: String?
    let syllabusId: Int?
    let examNodeId: Int?
    let year: Int?
    let region: String?
    let questionCount: Int?
}

// MARK: - Practice Models

struct PracticeAttempt: Codable, Identifiable {
    let id: Int
    let userId: Int?
    let questionId: Int?
    let paperId: Int?
    let isCorrect: Bool?
    let userAnswer: String?
    let timeTaken: Int?
    let createdAt: String?
}

struct QuickPracticeRequest: Encodable {
    let syllabusId: Int?
    let chapterId: Int?
    let questionType: String?
    let count: Int?
}

struct GradePracticeRequest: Encodable {
    let questionId: Int
    let userAnswer: String
    let timeTaken: Int?
}

struct AttemptStats: Codable {
    let total: Int?
    let correct: Int?
    let incorrect: Int?
    let accuracy: Double?
    let totalTime: Int?
}

// MARK: - Knowledge State Models

struct KnowledgeState: Codable, Identifiable {
    let id: Int
    let userId: Int?
    let chapterId: Int?
    let masteryLevel: Double?
    let reviewCount: Int?
    let nextReviewAt: String?
    let updatedAt: String?
}

struct KnowledgeProgress: Codable {
    let mastered: Int?
    let learning: Int?
    let notStarted: Int?
    let total: Int?
}

// MARK: - Syllabus & Exam Node Models

struct Syllabus: Codable, Identifiable {
    let id: Int
    let name: String
    let description: String?
    let subject: String?
    let grade: String?
}

struct ExamNode: Codable, Identifiable {
    let id: Int
    let name: String
    let syllabusId: Int?
    let parentId: Int?
    let sortOrder: Int?
    let description: String?
}

// MARK: - Analytics Models

struct ClassAnalyticsSummary: Codable {
    let classId: Int?
    let averageScore: Double?
    let totalStudents: Int?
    let activeStudents: Int?
    let totalAttempts: Int?
}

// MARK: - AI Chat Models

struct ChatMessage: Identifiable, Codable {
    let id: UUID
    let role: ChatRole
    let content: String
    let timestamp: Date

    enum ChatRole: String, Codable {
        case user
        case assistant
        case system
    }

    init(id: UUID = UUID(), role: ChatRole, content: String, timestamp: Date = Date()) {
        self.id = id
        self.role = role
        self.content = content
        self.timestamp = timestamp
    }
}

struct ChatSession: Identifiable, Codable {
    let id: UUID
    var title: String
    var messages: [ChatMessage]
    let createdAt: Date
    var updatedAt: Date

    init(id: UUID = UUID(), title: String = "New Chat", messages: [ChatMessage] = []) {
        self.id = id
        self.title = title
        self.messages = messages
        self.createdAt = Date()
        self.updatedAt = Date()
    }
}
