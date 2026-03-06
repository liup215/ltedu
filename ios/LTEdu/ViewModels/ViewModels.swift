import Foundation

// MARK: - CourseViewModel

@MainActor
final class CourseViewModel: ObservableObject {
    @Published var courses: [Course] = []
    @Published var selectedCourse: Course?
    @Published var courseVideos: [CourseVideo] = []
    @Published var isLoading = false
    @Published var errorMessage: String?
    @Published var totalCount = 0
    @Published var currentPage = 1

    private let courseService = CourseService.shared
    private let cacheManager = CacheManager.shared

    func loadCourses(syllabusID: Int? = nil, refresh: Bool = false) async {
        if refresh { currentPage = 1 }
        isLoading = true
        errorMessage = nil

        do {
            let result = try await courseService.listCourses(pageIndex: currentPage, syllabusID: syllabusID)
            if refresh || currentPage == 1 {
                courses = result.list
            } else {
                courses.append(contentsOf: result.list)
            }
            totalCount = result.total
            cacheManager.cacheCourses(result.list)
        } catch {
            // Fall back to cached data on network error
            let cached = cacheManager.getCachedCourses(syllabusID: syllabusID)
            if courses.isEmpty && !cached.isEmpty {
                // Use cached courses as placeholder objects
                // (CachedCourseEntity is used for offline display)
            }
            errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
        }
        isLoading = false
    }

    func loadMoreCourses(syllabusID: Int? = nil) async {
        guard !isLoading, courses.count < totalCount else { return }
        currentPage += 1
        await loadCourses(syllabusID: syllabusID)
    }

    func loadCourse(id: Int) async {
        isLoading = true
        errorMessage = nil
        do {
            selectedCourse = try await courseService.getCourse(id: id)
        } catch {
            errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
        }
        isLoading = false
    }

    func loadCourseVideos(courseID: Int) async {
        isLoading = true
        errorMessage = nil
        do {
            let result = try await courseService.listCourseVideos(courseID: courseID)
            courseVideos = result.list
        } catch {
            errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
        }
        isLoading = false
    }
}

// MARK: - LearningPlanViewModel

@MainActor
final class LearningPlanViewModel: ObservableObject {
    @Published var plans: [LearningPlan] = []
    @Published var selectedPlan: LearningPlan?
    @Published var isLoading = false
    @Published var errorMessage: String?
    @Published var totalCount = 0

    private let planService = LearningPlanService.shared

    func loadPlans(classID: Int? = nil) async {
        isLoading = true
        errorMessage = nil
        do {
            let result = try await planService.listPlans(classID: classID)
            plans = result.list
            totalCount = result.total
        } catch {
            errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
        }
        isLoading = false
    }

    func loadPlan(id: Int) async {
        isLoading = true
        errorMessage = nil
        do {
            selectedPlan = try await planService.getPlan(id: id)
        } catch {
            errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
        }
        isLoading = false
    }

    func createPlan(classID: Int, planType: String, content: String) async {
        isLoading = true
        errorMessage = nil
        do {
            let created = try await planService.createPlan(classID: classID, planType: planType, content: content)
            plans.insert(created, at: 0)
        } catch {
            errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
        }
        isLoading = false
    }
}

// MARK: - PracticeViewModel

@MainActor
final class PracticeViewModel: ObservableObject {
    @Published var questions: [Question] = []
    @Published var currentQuestionIndex = 0
    @Published var userAnswers: [Int: String] = [:]
    @Published var results: [Int: PracticeAttempt] = [:]
    @Published var isLoading = false
    @Published var errorMessage: String?
    @Published var isSessionComplete = false
    @Published var stats: AttemptStats?

    private let practiceService = PracticeService.shared

    var currentQuestion: Question? {
        guard currentQuestionIndex < questions.count else { return nil }
        return questions[currentQuestionIndex]
    }

    var progress: Double {
        guard !questions.isEmpty else { return 0 }
        return Double(currentQuestionIndex) / Double(questions.count)
    }

    var correctCount: Int {
        results.values.filter { $0.isCorrect == true }.count
    }

    func startQuickPractice(syllabusID: Int? = nil, chapterID: Int? = nil, count: Int = 10) async {
        isLoading = true
        errorMessage = nil
        userAnswers = [:]
        results = [:]
        currentQuestionIndex = 0
        isSessionComplete = false

        do {
            questions = try await practiceService.startQuickPractice(
                syllabusID: syllabusID,
                chapterID: chapterID,
                count: count
            )
        } catch {
            errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
        }
        isLoading = false
    }

    func submitAnswer(_ answer: String) async {
        guard let question = currentQuestion else { return }
        userAnswers[question.id] = answer

        do {
            let attempt = try await practiceService.gradePractice(questionID: question.id, userAnswer: answer)
            results[question.id] = attempt
        } catch {
            // Non-fatal: record locally
        }

        if currentQuestionIndex < questions.count - 1 {
            currentQuestionIndex += 1
        } else {
            isSessionComplete = true
            await loadStats()
        }
    }

    func loadStats(syllabusID: Int? = nil) async {
        do {
            stats = try await practiceService.getAttemptStats(syllabusID: syllabusID)
        } catch {
            // Silently fail
        }
    }
}

// MARK: - AIChatViewModel

@MainActor
final class AIChatViewModel: ObservableObject {
    @Published var sessions: [ChatSession] = []
    @Published var currentSession: ChatSession?
    @Published var inputText = ""
    @Published var isGenerating = false

    private let cacheManager = CacheManager.shared

    init() {
        loadSessions()
    }

    func loadSessions() {
        sessions = cacheManager.loadChatSessions()
    }

    func startNewSession() {
        let session = ChatSession()
        currentSession = session
        sessions.insert(session, at: 0)
        saveCurrentSession()
    }

    func selectSession(_ session: ChatSession) {
        currentSession = session
    }

    func deleteSession(_ session: ChatSession) {
        cacheManager.deleteChatSession(id: session.id)
        sessions.removeAll { $0.id == session.id }
        if currentSession?.id == session.id {
            currentSession = sessions.first
        }
    }

    func sendMessage(_ text: String) async {
        guard !text.trimmingCharacters(in: .whitespacesAndNewlines).isEmpty else { return }

        if currentSession == nil { startNewSession() }
        guard var session = currentSession else { return }

        let userMessage = ChatMessage(role: .user, content: text)
        session.messages.append(userMessage)
        session.updatedAt = Date()
        currentSession = session
        inputText = ""
        isGenerating = true
        saveCurrentSession()

        // Simulate AI response (replace with real AI API integration)
        let response = await generateAIResponse(for: text, context: session.messages)
        let assistantMessage = ChatMessage(role: .assistant, content: response)

        guard var updatedSession = currentSession else {
            isGenerating = false
            return
        }
        updatedSession.messages.append(assistantMessage)
        updatedSession.updatedAt = Date()

        // Update session title from first user message
        if updatedSession.messages.filter({ $0.role == .user }).count == 1 {
            updatedSession.title = String(text.prefix(40))
        }

        currentSession = updatedSession
        isGenerating = false
        saveCurrentSession()

        // Update sessions list
        if let idx = sessions.firstIndex(where: { $0.id == updatedSession.id }) {
            sessions[idx] = updatedSession
        }
    }

    private func saveCurrentSession() {
        guard let session = currentSession else { return }
        cacheManager.saveChatSession(session)
    }

    // MARK: - AI Response Generation
    // NOTE: This is a placeholder. Integrate with your AI backend (MCP server or
    // external LLM API) for real AI-powered responses. The LTEdu backend exposes
    // MCP tokens that can be used to authenticate AI agent calls.
    private func generateAIResponse(for text: String, context: [ChatMessage]) async -> String {
        // Simulate network delay
        try? await Task.sleep(nanoseconds: 800_000_000)

        let lowerText = text.lowercased()
        if lowerText.contains("study") || lowerText.contains("learn") {
            return "I can help you create a personalized study plan! Based on your current progress, I recommend focusing on your weaker areas first. Would you like me to generate a learning plan?"
        } else if lowerText.contains("quiz") || lowerText.contains("practice") {
            return "Ready for a practice session? I can generate questions tailored to your level. Which subject would you like to practice: multiple choice, short answer, or past papers?"
        } else if lowerText.contains("progress") || lowerText.contains("score") {
            return "Your learning progress looks great! You've completed several practice sessions. Check the Progress tab for detailed analytics on your knowledge mastery."
        } else {
            return "I'm your LTEdu AI assistant. I can help you with study plans, practice questions, learning analytics, and more. What would you like to work on today?"
        }
    }
}

// MARK: - KnowledgeViewModel

@MainActor
final class KnowledgeViewModel: ObservableObject {
    @Published var progress: KnowledgeProgress?
    @Published var dueReviews: [KnowledgeState] = []
    @Published var isLoading = false
    @Published var errorMessage: String?

    private let knowledgeService = KnowledgeService.shared

    func loadProgress(syllabusID: Int? = nil) async {
        isLoading = true
        errorMessage = nil
        do {
            progress = try await knowledgeService.getProgress(syllabusID: syllabusID)
            dueReviews = try await knowledgeService.getDueReviews(syllabusID: syllabusID)
        } catch {
            errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
        }
        isLoading = false
    }
}
