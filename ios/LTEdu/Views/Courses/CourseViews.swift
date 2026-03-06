import SwiftUI

// MARK: - CourseListView

struct CourseListView: View {
    @EnvironmentObject var courseVM: CourseViewModel
    @State private var searchText = ""

    var filteredCourses: [Course] {
        guard !searchText.isEmpty else { return courseVM.courses }
        return courseVM.courses.filter { $0.title.localizedCaseInsensitiveContains(searchText) }
    }

    var body: some View {
        NavigationStack {
            Group {
                if courseVM.isLoading && courseVM.courses.isEmpty {
                    ProgressView("Loading courses...")
                } else if courseVM.courses.isEmpty {
                    emptyState
                } else {
                    courseList
                }
            }
            .navigationTitle("Courses")
            .searchable(text: $searchText, prompt: "Search courses")
            .task {
                if courseVM.courses.isEmpty {
                    await courseVM.loadCourses(refresh: true)
                }
            }
            .refreshable {
                await courseVM.loadCourses(refresh: true)
            }
        }
    }

    private var courseList: some View {
        List {
            ForEach(filteredCourses) { course in
                NavigationLink {
                    CourseDetailView(course: course)
                        .environmentObject(courseVM)
                } label: {
                    CourseRowView(course: course)
                }
            }

            if courseVM.courses.count < courseVM.totalCount {
                ProgressView()
                    .frame(maxWidth: .infinity)
                    .onAppear {
                        Task { await courseVM.loadMoreCourses() }
                    }
            }
        }
        .listStyle(.plain)
    }

    private var emptyState: some View {
        VStack(spacing: 16) {
            Image(systemName: "books.vertical")
                .font(.system(size: 50))
                .foregroundColor(.secondary)
            Text("No courses available")
                .font(.headline)
            if let error = courseVM.errorMessage {
                Text(error)
                    .font(.caption)
                    .foregroundColor(.red)
                    .multilineTextAlignment(.center)
            }
            Button("Retry") {
                Task { await courseVM.loadCourses(refresh: true) }
            }
        }
        .padding()
    }
}

// MARK: - CourseRowView

struct CourseRowView: View {
    let course: Course

    var body: some View {
        HStack(spacing: 12) {
            // Thumbnail
            AsyncImage(url: URL(string: course.thumb ?? "")) { image in
                image.resizable().scaledToFill()
            } placeholder: {
                Color.blue.opacity(0.1)
                    .overlay(
                        Image(systemName: "play.rectangle.fill")
                            .foregroundColor(.blue.opacity(0.5))
                    )
            }
            .frame(width: 70, height: 50)
            .cornerRadius(8)
            .clipped()

            // Info
            VStack(alignment: .leading, spacing: 4) {
                Text(course.title)
                    .font(.headline)
                    .lineLimit(2)

                if let desc = course.shortDescription {
                    Text(desc)
                        .font(.caption)
                        .foregroundColor(.secondary)
                        .lineLimit(1)
                }

                HStack(spacing: 8) {
                    Label("\(course.userCount ?? 0)", systemImage: "person.2")
                        .font(.caption2)
                        .foregroundColor(.secondary)

                    if course.isFreeAccess {
                        Text("FREE")
                            .font(.caption2)
                            .fontWeight(.semibold)
                            .foregroundColor(.green)
                            .padding(.horizontal, 6)
                            .padding(.vertical, 2)
                            .background(Color.green.opacity(0.1))
                            .cornerRadius(4)
                    }
                }
            }
        }
        .padding(.vertical, 4)
    }
}

// MARK: - CourseDetailView

struct CourseDetailView: View {
    @EnvironmentObject var courseVM: CourseViewModel
    let course: Course
    @State private var selectedTab = 0

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                // Hero image
                AsyncImage(url: URL(string: course.thumb ?? "")) { image in
                    image.resizable().scaledToFill()
                } placeholder: {
                    Color.blue.opacity(0.15)
                }
                .frame(maxWidth: .infinity)
                .frame(height: 200)
                .clipped()

                VStack(alignment: .leading, spacing: 16) {
                    // Title & meta
                    VStack(alignment: .leading, spacing: 8) {
                        Text(course.title)
                            .font(.title2)
                            .fontWeight(.bold)

                        HStack {
                            Label("\(course.userCount ?? 0) students", systemImage: "person.2")
                                .font(.subheadline)
                                .foregroundColor(.secondary)
                            Spacer()
                            if course.isFreeAccess {
                                Text("FREE")
                                    .font(.headline)
                                    .fontWeight(.bold)
                                    .foregroundColor(.green)
                            } else {
                                Text("Premium")
                                    .font(.headline)
                                    .fontWeight(.bold)
                                    .foregroundColor(.orange)
                            }
                        }
                    }

                    Divider()

                    // Tab: Overview / Videos
                    Picker("", selection: $selectedTab) {
                        Text("Overview").tag(0)
                        Text("Videos").tag(1)
                    }
                    .pickerStyle(.segmented)

                    if selectedTab == 0 {
                        overviewTab
                    } else {
                        videosTab
                    }
                }
                .padding()
            }
        }
        .navigationTitle(course.title)
        .navigationBarTitleDisplayMode(.inline)
        .task {
            await courseVM.loadCourseVideos(courseID: course.id)
        }
    }

    private var overviewTab: some View {
        VStack(alignment: .leading, spacing: 12) {
            if let desc = course.shortDescription {
                Text(desc)
                    .font(.body)
                    .foregroundColor(.secondary)
            }
            if let desc = course.renderDesc ?? course.originalDesc {
                Text(desc)
                    .font(.body)
            }
        }
    }

    private var videosTab: some View {
        VStack(spacing: 0) {
            if courseVM.isLoading {
                ProgressView().padding()
            } else if courseVM.courseVideos.isEmpty {
                Text("No videos available")
                    .foregroundColor(.secondary)
                    .padding()
            } else {
                ForEach(courseVM.courseVideos) { video in
                    VideoRowView(video: video)
                    Divider()
                }
            }
        }
    }
}

// MARK: - VideoRowView

struct VideoRowView: View {
    let video: CourseVideo

    var durationText: String {
        guard let duration = video.duration else { return "" }
        let minutes = duration / 60
        let seconds = duration % 60
        return String(format: "%d:%02d", minutes, seconds)
    }

    var body: some View {
        HStack(spacing: 12) {
            Image(systemName: "play.circle.fill")
                .font(.system(size: 30))
                .foregroundColor(.blue)

            VStack(alignment: .leading, spacing: 2) {
                Text(video.title)
                    .font(.subheadline)
                    .lineLimit(2)
                if !durationText.isEmpty {
                    Text(durationText)
                        .font(.caption)
                        .foregroundColor(.secondary)
                }
            }

            Spacer()

            if video.isFree == 1 {
                Image(systemName: "lock.open")
                    .font(.caption)
                    .foregroundColor(.green)
            } else {
                Image(systemName: "lock")
                    .font(.caption)
                    .foregroundColor(.orange)
            }
        }
        .padding(.vertical, 8)
    }
}
