import Foundation

// MARK: - CourseService

final class CourseService {
    static let shared = CourseService()
    private let apiClient = APIClient.shared

    private init() {}

    func listCourses(pageIndex: Int = 1, pageSize: Int = APIConfig.defaultPageSize, syllabusID: Int? = nil) async throws -> PaginatedData<Course> {
        var params: [String: AnyCodable] = [
            "pageIndex": AnyCodable(pageIndex),
            "pageSize": AnyCodable(pageSize)
        ]
        if let syllabusID = syllabusID {
            params["syllabusId"] = AnyCodable(syllabusID)
        }
        return try await apiClient.post(endpoint: APIEndpoint.courseList, body: params, requiresAuth: false)
    }

    func getCourse(id: Int) async throws -> Course {
        return try await apiClient.post(endpoint: APIEndpoint.courseById, body: IDRequest(id: id), requiresAuth: false)
    }

    func listCourseVideos(courseID: Int, pageIndex: Int = 1, pageSize: Int = 50) async throws -> PaginatedData<CourseVideo> {
        let params: [String: AnyCodable] = [
            "courseId": AnyCodable(courseID),
            "pageIndex": AnyCodable(pageIndex),
            "pageSize": AnyCodable(pageSize)
        ]
        return try await apiClient.post(endpoint: APIEndpoint.courseVideoList, body: params, requiresAuth: false)
    }

    func getCourseVideo(id: Int) async throws -> CourseVideo {
        return try await apiClient.post(endpoint: APIEndpoint.courseVideoById, body: IDRequest(id: id), requiresAuth: false)
    }
}

// MARK: - QuestionService

final class QuestionService {
    static let shared = QuestionService()
    private let apiClient = APIClient.shared

    private init() {}

    func listQuestions(
        syllabusID: Int? = nil,
        chapterID: Int? = nil,
        questionType: String? = nil,
        pageIndex: Int = 1,
        pageSize: Int = APIConfig.defaultPageSize
    ) async throws -> PaginatedData<Question> {
        var params: [String: AnyCodable] = [
            "pageIndex": AnyCodable(pageIndex),
            "pageSize": AnyCodable(pageSize)
        ]
        if let syllabusID = syllabusID { params["syllabusId"] = AnyCodable(syllabusID) }
        if let chapterID = chapterID { params["chapterId"] = AnyCodable(chapterID) }
        if let questionType = questionType { params["questionType"] = AnyCodable(questionType) }

        return try await apiClient.post(endpoint: APIEndpoint.questionList, body: params, requiresAuth: false)
    }

    func getQuestion(id: Int) async throws -> Question {
        return try await apiClient.post(endpoint: APIEndpoint.questionById, body: IDRequest(id: id), requiresAuth: false)
    }

    func listPastPapers(syllabusID: Int? = nil, pageIndex: Int = 1, pageSize: Int = APIConfig.defaultPageSize) async throws -> PaginatedData<Paper> {
        var params: [String: AnyCodable] = [
            "pageIndex": AnyCodable(pageIndex),
            "pageSize": AnyCodable(pageSize)
        ]
        if let syllabusID = syllabusID { params["syllabusId"] = AnyCodable(syllabusID) }
        return try await apiClient.post(endpoint: APIEndpoint.pastPaperList, body: params, requiresAuth: false)
    }
}

// MARK: - PracticeService

final class PracticeService {
    static let shared = PracticeService()
    private let apiClient = APIClient.shared

    private init() {}

    func startQuickPractice(syllabusID: Int? = nil, chapterID: Int? = nil, questionType: String? = nil, count: Int = 10) async throws -> [Question] {
        let request = QuickPracticeRequest(
            syllabusId: syllabusID,
            chapterId: chapterID,
            questionType: questionType,
            count: count
        )
        return try await apiClient.post(endpoint: APIEndpoint.practiceQuick, body: request)
    }

    func gradePractice(questionID: Int, userAnswer: String, timeTaken: Int? = nil) async throws -> PracticeAttempt {
        let request = GradePracticeRequest(questionId: questionID, userAnswer: userAnswer, timeTaken: timeTaken)
        return try await apiClient.post(endpoint: APIEndpoint.practiceGrade, body: request)
    }

    func getAttemptStats(userID: Int? = nil, syllabusID: Int? = nil) async throws -> AttemptStats {
        var params: [String: AnyCodable] = [:]
        if let userID = userID { params["userId"] = AnyCodable(userID) }
        if let syllabusID = syllabusID { params["syllabusId"] = AnyCodable(syllabusID) }
        return try await apiClient.post(endpoint: APIEndpoint.attemptStats, body: params)
    }
}

// MARK: - KnowledgeService

final class KnowledgeService {
    static let shared = KnowledgeService()
    private let apiClient = APIClient.shared

    private init() {}

    func getProgress(syllabusID: Int? = nil) async throws -> KnowledgeProgress {
        var params: [String: AnyCodable] = [:]
        if let syllabusID = syllabusID { params["syllabusId"] = AnyCodable(syllabusID) }
        return try await apiClient.post(endpoint: APIEndpoint.knowledgeStateProgress, body: params)
    }

    func getDueReviews(syllabusID: Int? = nil) async throws -> [KnowledgeState] {
        var params: [String: AnyCodable] = [:]
        if let syllabusID = syllabusID { params["syllabusId"] = AnyCodable(syllabusID) }
        return try await apiClient.post(endpoint: APIEndpoint.knowledgeStateDueReview, body: params)
    }
}

// MARK: - ClassService

final class ClassService {
    static let shared = ClassService()
    private let apiClient = APIClient.shared

    private init() {}

    func listClasses(pageIndex: Int = 1, pageSize: Int = APIConfig.defaultPageSize) async throws -> PaginatedData<SchoolClass> {
        let params: [String: AnyCodable] = [
            "pageIndex": AnyCodable(pageIndex),
            "pageSize": AnyCodable(pageSize)
        ]
        return try await apiClient.post(endpoint: APIEndpoint.classList, body: params)
    }

    func getClass(id: Int) async throws -> SchoolClass {
        return try await apiClient.post(endpoint: APIEndpoint.classById, body: IDRequest(id: id))
    }

    func applyToClass(classID: Int) async throws {
        try await apiClient.postEmpty(endpoint: APIEndpoint.classApply, body: IDRequest(id: classID))
    }
}

// MARK: - LearningPlanService

final class LearningPlanService {
    static let shared = LearningPlanService()
    private let apiClient = APIClient.shared

    private init() {}

    func listPlans(classID: Int? = nil, pageIndex: Int = 1, pageSize: Int = APIConfig.defaultPageSize) async throws -> PaginatedData<LearningPlan> {
        var params: [String: AnyCodable] = [
            "pageIndex": AnyCodable(pageIndex),
            "pageSize": AnyCodable(pageSize)
        ]
        if let classID = classID { params["classId"] = AnyCodable(classID) }
        return try await apiClient.post(endpoint: APIEndpoint.learningPlanList, body: params)
    }

    func getPlan(id: Int) async throws -> LearningPlan {
        return try await apiClient.post(endpoint: APIEndpoint.learningPlanById, body: IDRequest(id: id))
    }

    func createPlan(classID: Int, planType: String, content: String, isPersonal: Bool = false) async throws -> LearningPlan {
        let request = CreateLearningPlanRequest(classId: classID, planType: planType, content: content, isPersonal: isPersonal)
        return try await apiClient.post(endpoint: APIEndpoint.learningPlanCreate, body: request)
    }
}

// MARK: - SyllabusService

final class SyllabusService {
    static let shared = SyllabusService()
    private let apiClient = APIClient.shared

    private init() {}

    func listSyllabuses(pageIndex: Int = 1, pageSize: Int = APIConfig.defaultPageSize) async throws -> PaginatedData<Syllabus> {
        let params: [String: AnyCodable] = [
            "pageIndex": AnyCodable(pageIndex),
            "pageSize": AnyCodable(pageSize)
        ]
        return try await apiClient.post(endpoint: APIEndpoint.syllabusList, body: params)
    }
}

// MARK: - Shared Types

struct IDRequest: Encodable {
    let id: Int
}

/// Type-erased Codable wrapper for heterogeneous dictionaries
struct AnyCodable: Codable {
    let value: Any

    init(_ value: Any) {
        self.value = value
    }

    init(from decoder: Decoder) throws {
        let container = try decoder.singleValueContainer()
        if let intVal = try? container.decode(Int.self) {
            value = intVal
        } else if let doubleVal = try? container.decode(Double.self) {
            value = doubleVal
        } else if let strVal = try? container.decode(String.self) {
            value = strVal
        } else if let boolVal = try? container.decode(Bool.self) {
            value = boolVal
        } else {
            value = NSNull()
        }
    }

    func encode(to encoder: Encoder) throws {
        var container = encoder.singleValueContainer()
        switch value {
        case let intVal as Int:
            try container.encode(intVal)
        case let doubleVal as Double:
            try container.encode(doubleVal)
        case let strVal as String:
            try container.encode(strVal)
        case let boolVal as Bool:
            try container.encode(boolVal)
        default:
            try container.encodeNil()
        }
    }
}
