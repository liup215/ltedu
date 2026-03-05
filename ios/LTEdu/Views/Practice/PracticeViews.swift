import SwiftUI

// MARK: - PracticeView

struct PracticeView: View {
    @EnvironmentObject var practiceVM: PracticeViewModel
    @State private var showSetup = false

    var body: some View {
        NavigationStack {
            Group {
                if practiceVM.isSessionComplete {
                    PracticeResultsView()
                        .environmentObject(practiceVM)
                } else if !practiceVM.questions.isEmpty {
                    PracticeSessionView()
                        .environmentObject(practiceVM)
                } else {
                    practiceHomeView
                }
            }
            .navigationTitle("Practice")
            .sheet(isPresented: $showSetup) {
                PracticeSetupView()
                    .environmentObject(practiceVM)
            }
        }
    }

    private var practiceHomeView: some View {
        ScrollView {
            VStack(spacing: 24) {
                // Header
                VStack(spacing: 8) {
                    Image(systemName: "pencil.and.list.clipboard")
                        .font(.system(size: 50))
                        .foregroundColor(.blue)
                    Text("Practice Zone")
                        .font(.title2)
                        .fontWeight(.bold)
                    Text("Strengthen your knowledge with targeted practice")
                        .font(.subheadline)
                        .foregroundColor(.secondary)
                        .multilineTextAlignment(.center)
                }
                .padding(.top, 20)

                // Quick start options
                VStack(spacing: 12) {
                    practiceOptionCard(
                        title: "Quick Practice",
                        subtitle: "10 random questions",
                        icon: "bolt.fill",
                        color: .blue
                    ) {
                        Task { await practiceVM.startQuickPractice() }
                    }

                    practiceOptionCard(
                        title: "Custom Practice",
                        subtitle: "Choose topic & difficulty",
                        icon: "slider.horizontal.3",
                        color: .purple
                    ) {
                        showSetup = true
                    }

                    practiceOptionCard(
                        title: "Past Papers",
                        subtitle: "Practice with real exam papers",
                        icon: "doc.text.fill",
                        color: .orange
                    ) {
                        // Navigate to past papers
                    }
                }
                .padding(.horizontal)

                // Stats card
                if let stats = practiceVM.stats {
                    statsCard(stats)
                        .padding(.horizontal)
                }
            }
            .padding(.bottom, 40)
        }
        .task {
            await practiceVM.loadStats()
        }
    }

    private func practiceOptionCard(title: String, subtitle: String, icon: String, color: Color, action: @escaping () -> Void) -> some View {
        Button(action: action) {
            HStack(spacing: 16) {
                Image(systemName: icon)
                    .font(.title2)
                    .frame(width: 50, height: 50)
                    .background(color.opacity(0.1))
                    .foregroundColor(color)
                    .cornerRadius(12)

                VStack(alignment: .leading, spacing: 2) {
                    Text(title)
                        .font(.headline)
                        .foregroundColor(.primary)
                    Text(subtitle)
                        .font(.subheadline)
                        .foregroundColor(.secondary)
                }

                Spacer()

                Image(systemName: "chevron.right")
                    .foregroundColor(.secondary)
            }
            .padding()
            .background(Color(.systemBackground))
            .cornerRadius(14)
            .shadow(color: .black.opacity(0.05), radius: 6, x: 0, y: 2)
        }
    }

    private func statsCard(_ stats: AttemptStats) -> some View {
        VStack(alignment: .leading, spacing: 12) {
            Text("Your Stats")
                .font(.headline)

            HStack(spacing: 20) {
                statItem(value: "\(stats.total ?? 0)", label: "Total", color: .blue)
                statItem(value: "\(stats.correct ?? 0)", label: "Correct", color: .green)
                statItem(value: String(format: "%.0f%%", (stats.accuracy ?? 0) * 100), label: "Accuracy", color: .orange)
            }
        }
        .padding()
        .background(Color(.systemBackground))
        .cornerRadius(14)
        .shadow(color: .black.opacity(0.05), radius: 6, x: 0, y: 2)
    }

    private func statItem(value: String, label: String, color: Color) -> some View {
        VStack(spacing: 4) {
            Text(value)
                .font(.title2)
                .fontWeight(.bold)
                .foregroundColor(color)
            Text(label)
                .font(.caption)
                .foregroundColor(.secondary)
        }
        .frame(maxWidth: .infinity)
    }
}

// MARK: - PracticeSessionView

struct PracticeSessionView: View {
    @EnvironmentObject var practiceVM: PracticeViewModel
    @State private var selectedAnswer = ""
    @State private var isAnswered = false

    var body: some View {
        VStack(spacing: 0) {
            // Progress bar
            progressBar

            ScrollView {
                VStack(alignment: .leading, spacing: 20) {
                    if let question = practiceVM.currentQuestion {
                        // Question header
                        questionHeader(question)

                        // Question stem
                        if let stem = question.content?.stem {
                            Text(stem)
                                .font(.body)
                                .padding()
                                .frame(maxWidth: .infinity, alignment: .leading)
                                .background(Color(.systemGray6))
                                .cornerRadius(12)
                        }

                        // Answer input
                        answerSection(question)

                        // Result feedback
                        if isAnswered, let result = practiceVM.results[question.id] {
                            resultFeedback(result, question: question)
                        }

                        // Next button
                        if isAnswered {
                            Button(practiceVM.currentQuestionIndex < practiceVM.questions.count - 1 ? "Next Question" : "See Results") {
                                Task {
                                    await practiceVM.submitAnswer(selectedAnswer)
                                    selectedAnswer = ""
                                    isAnswered = false
                                }
                            }
                            .frame(maxWidth: .infinity)
                            .frame(height: 50)
                            .background(Color.blue)
                            .foregroundColor(.white)
                            .cornerRadius(12)
                        }
                    }
                }
                .padding()
            }
        }
        .navigationTitle("Question \(practiceVM.currentQuestionIndex + 1)/\(practiceVM.questions.count)")
        .navigationBarTitleDisplayMode(.inline)
    }

    private var progressBar: some View {
        GeometryReader { geo in
            ZStack(alignment: .leading) {
                Rectangle()
                    .fill(Color(.systemGray5))
                Rectangle()
                    .fill(Color.blue)
                    .frame(width: geo.size.width * practiceVM.progress)
                    .animation(.easeInOut, value: practiceVM.progress)
            }
        }
        .frame(height: 4)
    }

    private func questionHeader(_ question: Question) -> some View {
        HStack {
            if let type = question.questionType {
                Text(type.capitalized)
                    .font(.caption)
                    .fontWeight(.semibold)
                    .padding(.horizontal, 10)
                    .padding(.vertical, 4)
                    .background(Color.blue.opacity(0.1))
                    .foregroundColor(.blue)
                    .cornerRadius(8)
            }

            if let level = question.difficultyLevel {
                HStack(spacing: 2) {
                    ForEach(0..<5) { i in
                        Image(systemName: i < level ? "star.fill" : "star")
                            .font(.caption2)
                            .foregroundColor(.orange)
                    }
                }
            }
        }
    }

    private func answerSection(_ question: Question) -> some View {
        VStack(alignment: .leading, spacing: 10) {
            Text("Your Answer")
                .font(.subheadline)
                .fontWeight(.semibold)

            TextField("Type your answer here...", text: $selectedAnswer, axis: .vertical)
                .lineLimit(3...8)
                .padding()
                .background(Color(.systemGray6))
                .cornerRadius(10)
                .disabled(isAnswered)

            if !isAnswered && !selectedAnswer.isEmpty {
                Button("Submit Answer") {
                    isAnswered = true
                    Task {
                        await practiceVM.submitAnswer(selectedAnswer)
                        selectedAnswer = ""
                        isAnswered = false
                    }
                }
                .frame(maxWidth: .infinity)
                .frame(height: 44)
                .background(Color.green)
                .foregroundColor(.white)
                .cornerRadius(10)
            }
        }
    }

    private func resultFeedback(_ attempt: PracticeAttempt, question: Question) -> some View {
        VStack(alignment: .leading, spacing: 8) {
            HStack {
                Image(systemName: attempt.isCorrect == true ? "checkmark.circle.fill" : "xmark.circle.fill")
                    .foregroundColor(attempt.isCorrect == true ? .green : .red)
                Text(attempt.isCorrect == true ? "Correct!" : "Incorrect")
                    .fontWeight(.semibold)
                    .foregroundColor(attempt.isCorrect == true ? .green : .red)
            }

            if let analyze = question.content?.analyze {
                Text("Explanation:")
                    .font(.subheadline)
                    .fontWeight(.semibold)
                Text(analyze)
                    .font(.body)
                    .foregroundColor(.secondary)
            }
        }
        .padding()
        .background((attempt.isCorrect == true ? Color.green : Color.red).opacity(0.08))
        .cornerRadius(12)
    }
}

// MARK: - PracticeResultsView

struct PracticeResultsView: View {
    @EnvironmentObject var practiceVM: PracticeViewModel

    var body: some View {
        VStack(spacing: 24) {
            Spacer()

            // Score circle
            ZStack {
                Circle()
                    .stroke(Color(.systemGray5), lineWidth: 12)
                    .frame(width: 150, height: 150)

                Circle()
                    .trim(from: 0, to: practiceVM.questions.isEmpty ? 0 : CGFloat(practiceVM.correctCount) / CGFloat(practiceVM.questions.count))
                    .stroke(
                        practiceVM.correctCount >= practiceVM.questions.count / 2 ? Color.green : Color.orange,
                        style: StrokeStyle(lineWidth: 12, lineCap: .round)
                    )
                    .rotationEffect(.degrees(-90))
                    .frame(width: 150, height: 150)
                    .animation(.easeInOut(duration: 1), value: practiceVM.correctCount)

                VStack(spacing: 2) {
                    Text("\(practiceVM.correctCount)/\(practiceVM.questions.count)")
                        .font(.title)
                        .fontWeight(.bold)
                    Text("Correct")
                        .font(.caption)
                        .foregroundColor(.secondary)
                }
            }

            Text(resultTitle)
                .font(.title2)
                .fontWeight(.bold)

            Text(resultMessage)
                .font(.subheadline)
                .foregroundColor(.secondary)
                .multilineTextAlignment(.center)
                .padding(.horizontal, 40)

            Spacer()

            VStack(spacing: 12) {
                Button("Try Again") {
                    Task { await practiceVM.startQuickPractice() }
                }
                .frame(maxWidth: .infinity)
                .frame(height: 50)
                .background(Color.blue)
                .foregroundColor(.white)
                .cornerRadius(12)
                .padding(.horizontal, 40)

                Button("Back to Practice") {
                    practiceVM.questions = []
                    practiceVM.isSessionComplete = false
                }
                .font(.subheadline)
                .foregroundColor(.blue)
            }
            .padding(.bottom, 40)
        }
        .navigationTitle("Results")
        .navigationBarBackButtonHidden()
    }

    private var accuracy: Double {
        guard !practiceVM.questions.isEmpty else { return 0 }
        return Double(practiceVM.correctCount) / Double(practiceVM.questions.count)
    }

    private var resultTitle: String {
        switch accuracy {
        case 0.9...: return "Excellent! 🎉"
        case 0.7..<0.9: return "Great Job! 👍"
        case 0.5..<0.7: return "Keep Going! 💪"
        default: return "Keep Practicing! 📚"
        }
    }

    private var resultMessage: String {
        switch accuracy {
        case 0.9...: return "Outstanding performance! You've mastered this material."
        case 0.7..<0.9: return "Good work! Review the questions you missed to improve further."
        case 0.5..<0.7: return "You're making progress. Focus on the topics where you struggled."
        default: return "Don't give up! Review the material and try again."
        }
    }
}

// MARK: - PracticeSetupView

struct PracticeSetupView: View {
    @EnvironmentObject var practiceVM: PracticeViewModel
    @Environment(\.dismiss) var dismiss

    @State private var questionCount = 10
    @State private var selectedType: String? = nil

    let questionTypes = ["multiple_choice", "short_answer", "fill_blank"]

    var body: some View {
        NavigationStack {
            Form {
                Section("Number of Questions") {
                    Stepper("\(questionCount) questions", value: $questionCount, in: 5...50, step: 5)
                }

                Section("Question Type") {
                    Button("All Types") {
                        selectedType = nil
                    }
                    .foregroundColor(selectedType == nil ? .blue : .primary)

                    ForEach(questionTypes, id: \.self) { type in
                        Button(type.replacingOccurrences(of: "_", with: " ").capitalized) {
                            selectedType = type
                        }
                        .foregroundColor(selectedType == type ? .blue : .primary)
                    }
                }
            }
            .navigationTitle("Practice Setup")
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .navigationBarLeading) {
                    Button("Cancel") { dismiss() }
                }
                ToolbarItem(placement: .navigationBarTrailing) {
                    Button("Start") {
                        dismiss()
                        Task {
                            await practiceVM.startQuickPractice(questionType: selectedType, count: questionCount)
                        }
                    }
                }
            }
        }
    }
}
