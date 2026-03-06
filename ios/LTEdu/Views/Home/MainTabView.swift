import SwiftUI

struct MainTabView: View {
    @EnvironmentObject var authViewModel: AuthViewModel
    @StateObject private var courseVM = CourseViewModel()
    @StateObject private var practiceVM = PracticeViewModel()
    @StateObject private var planVM = LearningPlanViewModel()
    @StateObject private var chatVM = AIChatViewModel()
    @StateObject private var knowledgeVM = KnowledgeViewModel()

    var body: some View {
        TabView {
            AIChatView()
                .environmentObject(chatVM)
                .tabItem {
                    Label("AI Chat", systemImage: "brain.head.profile")
                }

            HomeView()
                .environmentObject(courseVM)
                .environmentObject(knowledgeVM)
                .tabItem {
                    Label("Home", systemImage: "house.fill")
                }

            CourseListView()
                .environmentObject(courseVM)
                .tabItem {
                    Label("Courses", systemImage: "book.fill")
                }

            PracticeView()
                .environmentObject(practiceVM)
                .tabItem {
                    Label("Practice", systemImage: "pencil.and.list.clipboard")
                }

            ProfileView()
                .environmentObject(authViewModel)
                .environmentObject(planVM)
                .tabItem {
                    Label("Profile", systemImage: "person.fill")
                }
        }
        .accentColor(.blue)
    }
}
