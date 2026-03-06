import SwiftUI

struct HomeView: View {
    @EnvironmentObject var authViewModel: AuthViewModel
    @EnvironmentObject var courseVM: CourseViewModel
    @EnvironmentObject var knowledgeVM: KnowledgeViewModel

    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(alignment: .leading, spacing: 20) {
                    // Greeting
                    greetingSection

                    // Progress summary
                    if let progress = knowledgeVM.progress {
                        progressSection(progress)
                    }

                    // Due reviews
                    if !knowledgeVM.dueReviews.isEmpty {
                        dueReviewSection
                    }

                    // Recent courses
                    coursesSection
                }
                .padding()
            }
            .navigationTitle("LTEdu")
            .toolbar {
                ToolbarItem(placement: .navigationBarTrailing) {
                    notificationBell
                }
            }
            .task {
                await knowledgeVM.loadProgress()
                await courseVM.loadCourses(refresh: true)
            }
            .refreshable {
                await knowledgeVM.loadProgress()
                await courseVM.loadCourses(refresh: true)
            }
        }
    }

    // MARK: - Greeting

    private var greetingSection: some View {
        VStack(alignment: .leading, spacing: 4) {
            Text(greeting)
                .font(.title2)
                .fontWeight(.semibold)
            if let user = authViewModel.currentUser {
                Text(user.displayName)
                    .font(.largeTitle)
                    .fontWeight(.bold)
            }
        }
    }

    private var greeting: String {
        let hour = Calendar.current.component(.hour, from: Date())
        switch hour {
        case 0..<12: return "Good Morning,"
        case 12..<17: return "Good Afternoon,"
        default: return "Good Evening,"
        }
    }

    // MARK: - Progress Section

    private func progressSection(_ progress: KnowledgeProgress) -> some View {
        VStack(alignment: .leading, spacing: 12) {
            Text("Learning Progress")
                .font(.headline)

            HStack(spacing: 16) {
                progressCard(
                    value: progress.mastered ?? 0,
                    total: progress.total ?? 1,
                    label: "Mastered",
                    color: .green
                )
                progressCard(
                    value: progress.learning ?? 0,
                    total: progress.total ?? 1,
                    label: "In Progress",
                    color: .orange
                )
                progressCard(
                    value: progress.notStarted ?? 0,
                    total: progress.total ?? 1,
                    label: "Not Started",
                    color: .gray
                )
            }
        }
    }

    private func progressCard(value: Int, total: Int, label: String, color: Color) -> some View {
        VStack(spacing: 8) {
            ZStack {
                Circle()
                    .stroke(color.opacity(0.2), lineWidth: 6)
                Circle()
                    .trim(from: 0, to: total > 0 ? CGFloat(value) / CGFloat(total) : 0)
                    .stroke(color, style: StrokeStyle(lineWidth: 6, lineCap: .round))
                    .rotationEffect(.degrees(-90))
                Text("\(value)")
                    .font(.headline)
                    .fontWeight(.bold)
            }
            .frame(width: 60, height: 60)

            Text(label)
                .font(.caption)
                .foregroundColor(.secondary)
        }
        .frame(maxWidth: .infinity)
        .padding()
        .background(Color(.systemBackground))
        .cornerRadius(12)
        .shadow(color: .black.opacity(0.05), radius: 4, x: 0, y: 2)
    }

    // MARK: - Due Review Section

    private var dueReviewSection: some View {
        VStack(alignment: .leading, spacing: 12) {
            HStack {
                Text("Due for Review")
                    .font(.headline)
                Spacer()
                Text("\(knowledgeVM.dueReviews.count) topics")
                    .font(.caption)
                    .foregroundColor(.secondary)
            }

            ScrollView(.horizontal, showsIndicators: false) {
                HStack(spacing: 12) {
                    ForEach(knowledgeVM.dueReviews) { review in
                        reviewChip(review)
                    }
                }
            }
        }
    }

    private func reviewChip(_ state: KnowledgeState) -> some View {
        HStack(spacing: 6) {
            Image(systemName: "clock.fill")
                .font(.caption)
                .foregroundColor(.orange)
            Text("Chapter \(state.chapterId ?? 0)")
                .font(.caption)
                .fontWeight(.medium)
        }
        .padding(.horizontal, 12)
        .padding(.vertical, 8)
        .background(Color.orange.opacity(0.1))
        .cornerRadius(20)
    }

    // MARK: - Courses Section

    private var coursesSection: some View {
        VStack(alignment: .leading, spacing: 12) {
            HStack {
                Text("Courses")
                    .font(.headline)
                Spacer()
                NavigationLink("See All") {
                    CourseListView()
                        .environmentObject(courseVM)
                }
                .font(.caption)
            }

            if courseVM.isLoading {
                ProgressView()
                    .frame(maxWidth: .infinity)
            } else {
                ForEach(courseVM.courses.prefix(3)) { course in
                    NavigationLink {
                        CourseDetailView(course: course)
                            .environmentObject(courseVM)
                    } label: {
                        CourseRowView(course: course)
                    }
                    .buttonStyle(.plain)
                }
            }
        }
    }

    // MARK: - Notification Bell

    private var notificationBell: some View {
        Button(action: {}) {
            Image(systemName: "bell")
                .imageScale(.medium)
        }
    }
}
