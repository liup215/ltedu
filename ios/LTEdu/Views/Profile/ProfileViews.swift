import SwiftUI

// MARK: - ProfileView

struct ProfileView: View {
    @EnvironmentObject var authViewModel: AuthViewModel
    @EnvironmentObject var planVM: LearningPlanViewModel
    @State private var showEditProfile = false
    @State private var showChangePassword = false
    @State private var showLearningPlans = false
    @State private var showNotificationSettings = false

    var body: some View {
        NavigationStack {
            List {
                // User header
                Section {
                    userHeaderView
                }

                // Learning
                Section("Learning") {
                    NavigationLink {
                        LearningPlanListView()
                            .environmentObject(planVM)
                    } label: {
                        Label("Learning Plans", systemImage: "list.bullet.clipboard")
                    }

                    NavigationLink {
                        Text("Attempt History")
                    } label: {
                        Label("Attempt History", systemImage: "clock.arrow.circlepath")
                    }

                    NavigationLink {
                        Text("Knowledge Map")
                    } label: {
                        Label("Knowledge Map", systemImage: "map")
                    }
                }

                // Account
                Section("Account") {
                    Button {
                        showEditProfile = true
                    } label: {
                        Label("Edit Profile", systemImage: "person.crop.circle")
                    }
                    .foregroundColor(.primary)

                    Button {
                        showChangePassword = true
                    } label: {
                        Label("Change Password", systemImage: "lock")
                    }
                    .foregroundColor(.primary)

                    Button {
                        showNotificationSettings = true
                    } label: {
                        Label("Notifications", systemImage: "bell")
                    }
                    .foregroundColor(.primary)
                }

                // App Info
                Section("App") {
                    HStack {
                        Label("Version", systemImage: "info.circle")
                        Spacer()
                        Text(Bundle.main.infoDictionary?["CFBundleShortVersionString"] as? String ?? "1.0.0")
                            .foregroundColor(.secondary)
                    }
                }

                // Sign Out
                Section {
                    Button(role: .destructive) {
                        authViewModel.logout()
                    } label: {
                        HStack {
                            Spacer()
                            Label("Sign Out", systemImage: "arrow.right.square")
                            Spacer()
                        }
                    }
                }
            }
            .navigationTitle("Profile")
            .sheet(isPresented: $showEditProfile) {
                EditProfileView()
                    .environmentObject(authViewModel)
            }
            .sheet(isPresented: $showChangePassword) {
                ChangePasswordView()
            }
            .sheet(isPresented: $showNotificationSettings) {
                NotificationSettingsView()
            }
        }
    }

    private var userHeaderView: some View {
        HStack(spacing: 14) {
            // Avatar
            Group {
                if let avatar = authViewModel.currentUser?.avatar,
                   let url = URL(string: avatar) {
                    AsyncImage(url: url) { image in
                        image.resizable().scaledToFill()
                    } placeholder: {
                        avatarPlaceholder
                    }
                } else {
                    avatarPlaceholder
                }
            }
            .frame(width: 60, height: 60)
            .clipShape(Circle())

            VStack(alignment: .leading, spacing: 4) {
                Text(authViewModel.currentUser?.displayName ?? "User")
                    .font(.headline)
                Text(authViewModel.currentUser?.email ?? "")
                    .font(.caption)
                    .foregroundColor(.secondary)

                HStack(spacing: 6) {
                    if authViewModel.currentUser?.isAdmin == true {
                        badgeView("Admin", color: .red)
                    }
                    if authViewModel.currentUser?.isTeacher == true {
                        badgeView("Teacher", color: .blue)
                    }
                    if authViewModel.currentUser?.isVip == true {
                        badgeView("VIP", color: .orange)
                    }
                }
            }
        }
        .padding(.vertical, 4)
    }

    private var avatarPlaceholder: some View {
        Image(systemName: "person.circle.fill")
            .font(.system(size: 60))
            .foregroundColor(.blue.opacity(0.6))
    }

    private func badgeView(_ text: String, color: Color) -> some View {
        Text(text)
            .font(.caption2)
            .fontWeight(.semibold)
            .padding(.horizontal, 8)
            .padding(.vertical, 2)
            .background(color.opacity(0.1))
            .foregroundColor(color)
            .cornerRadius(6)
    }
}

// MARK: - EditProfileView

struct EditProfileView: View {
    @EnvironmentObject var authViewModel: AuthViewModel
    @Environment(\.dismiss) var dismiss

    @State private var realname = ""
    @State private var nickname = ""
    @State private var mobile = ""
    @State private var sex = 0

    var body: some View {
        NavigationStack {
            Form {
                Section("Personal Info") {
                    TextField("Real Name", text: $realname)
                    TextField("Nickname", text: $nickname)
                    TextField("Mobile", text: $mobile)
                        .keyboardType(.phonePad)
                }

                Section("Gender") {
                    Picker("Gender", selection: $sex) {
                        Text("Not specified").tag(0)
                        Text("Male").tag(1)
                        Text("Female").tag(2)
                    }
                    .pickerStyle(.segmented)
                }

                if let error = authViewModel.errorMessage {
                    Section {
                        Text(error)
                            .foregroundColor(.red)
                            .font(.caption)
                    }
                }
            }
            .navigationTitle("Edit Profile")
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .navigationBarLeading) {
                    Button("Cancel") { dismiss() }
                }
                ToolbarItem(placement: .navigationBarTrailing) {
                    Button("Save") {
                        Task {
                            await authViewModel.updateProfile(
                                realname: realname.isEmpty ? nil : realname,
                                nickname: nickname.isEmpty ? nil : nickname,
                                sex: sex == 0 ? nil : sex,
                                mobile: mobile.isEmpty ? nil : mobile
                            )
                            if authViewModel.errorMessage == nil { dismiss() }
                        }
                    }
                    .disabled(authViewModel.isLoading)
                }
            }
        }
        .onAppear {
            realname = authViewModel.currentUser?.realname ?? ""
            nickname = authViewModel.currentUser?.nickname ?? ""
            mobile = authViewModel.currentUser?.mobile ?? ""
            sex = authViewModel.currentUser?.sex ?? 0
        }
    }
}

// MARK: - ChangePasswordView

struct ChangePasswordView: View {
    @Environment(\.dismiss) var dismiss
    @State private var oldPassword = ""
    @State private var newPassword = ""
    @State private var confirmPassword = ""
    @State private var showOld = false
    @State private var showNew = false
    @State private var errorMessage: String?
    @State private var isLoading = false

    var body: some View {
        NavigationStack {
            Form {
                Section("Current Password") {
                    LTEduSecureField(text: $oldPassword, placeholder: "Current password", showPassword: $showOld)
                        .listRowInsets(EdgeInsets(top: 4, leading: 16, bottom: 4, trailing: 16))
                        .listRowBackground(Color.clear)
                }

                Section("New Password") {
                    LTEduSecureField(text: $newPassword, placeholder: "New password (min 6 chars)", showPassword: $showNew)
                        .listRowInsets(EdgeInsets(top: 4, leading: 16, bottom: 4, trailing: 16))
                        .listRowBackground(Color.clear)
                    LTEduSecureField(text: $confirmPassword, placeholder: "Confirm new password", showPassword: $showNew)
                        .listRowInsets(EdgeInsets(top: 4, leading: 16, bottom: 4, trailing: 16))
                        .listRowBackground(Color.clear)
                }

                if let error = errorMessage {
                    Text(error)
                        .foregroundColor(.red)
                        .font(.caption)
                }
            }
            .navigationTitle("Change Password")
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .navigationBarLeading) {
                    Button("Cancel") { dismiss() }
                }
                ToolbarItem(placement: .navigationBarTrailing) {
                    Button("Save") { changePassword() }
                        .disabled(isLoading)
                }
            }
        }
    }

    private func changePassword() {
        guard newPassword == confirmPassword else {
            errorMessage = "Passwords do not match"
            return
        }
        guard newPassword.count >= 6 else {
            errorMessage = "Password must be at least 6 characters"
            return
        }
        isLoading = true
        Task {
            do {
                try await AuthService.shared.changePassword(oldPassword: oldPassword, newPassword: newPassword)
                dismiss()
            } catch {
                errorMessage = (error as? APIError)?.errorDescription ?? error.localizedDescription
            }
            isLoading = false
        }
    }
}

// MARK: - NotificationSettingsView

struct NotificationSettingsView: View {
    @EnvironmentObject var notificationManager: PushNotificationManager
    @Environment(\.dismiss) var dismiss
    @State private var studyReminders = true
    @State private var reviewAlerts = true

    var body: some View {
        NavigationStack {
            Form {
                Section("Push Notifications") {
                    if notificationManager.isAuthorized {
                        Toggle("Study Reminders", isOn: $studyReminders)
                        Toggle("Review Due Alerts", isOn: $reviewAlerts)
                    } else {
                        VStack(alignment: .leading, spacing: 8) {
                            Text("Notifications are disabled")
                                .font(.subheadline)
                            Text("Enable notifications in Settings to receive study reminders and review alerts.")
                                .font(.caption)
                                .foregroundColor(.secondary)
                            Button("Open Settings") {
                                if let url = URL(string: UIApplication.openSettingsURLString) {
                                    UIApplication.shared.open(url)
                                }
                            }
                            .font(.subheadline)
                        }
                    }
                }
            }
            .navigationTitle("Notifications")
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .navigationBarTrailing) {
                    Button("Done") { dismiss() }
                }
            }
            .onAppear {
                notificationManager.checkAuthorizationStatus()
            }
        }
    }
}

// MARK: - LearningPlanListView

struct LearningPlanListView: View {
    @EnvironmentObject var planVM: LearningPlanViewModel

    var body: some View {
        Group {
            if planVM.isLoading {
                ProgressView("Loading plans...")
            } else if planVM.plans.isEmpty {
                VStack(spacing: 16) {
                    Image(systemName: "list.bullet.clipboard")
                        .font(.system(size: 40))
                        .foregroundColor(.secondary)
                    Text("No learning plans yet")
                        .font(.headline)
                    if let error = planVM.errorMessage {
                        Text(error)
                            .font(.caption)
                            .foregroundColor(.red)
                    }
                }
            } else {
                List(planVM.plans) { plan in
                    LearningPlanRowView(plan: plan)
                }
            }
        }
        .navigationTitle("Learning Plans")
        .task {
            if planVM.plans.isEmpty {
                await planVM.loadPlans()
            }
        }
        .refreshable {
            await planVM.loadPlans()
        }
    }
}

// MARK: - LearningPlanRowView

struct LearningPlanRowView: View {
    let plan: LearningPlan

    var planTypeLabel: String {
        switch plan.planType {
        case "long": return "Long-term"
        case "mid": return "Mid-term"
        case "short": return "Short-term"
        default: return plan.planType ?? "Plan"
        }
    }

    var planTypeColor: Color {
        switch plan.planType {
        case "long": return .blue
        case "mid": return .purple
        case "short": return .orange
        default: return .gray
        }
    }

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            HStack {
                Text(planTypeLabel)
                    .font(.caption)
                    .fontWeight(.semibold)
                    .padding(.horizontal, 8)
                    .padding(.vertical, 3)
                    .background(planTypeColor.opacity(0.1))
                    .foregroundColor(planTypeColor)
                    .cornerRadius(6)

                Spacer()

                if let date = plan.updatedAt {
                    Text(date.prefix(10))
                        .font(.caption2)
                        .foregroundColor(.secondary)
                }
            }

            if let content = plan.content {
                Text(content)
                    .font(.subheadline)
                    .lineLimit(3)
                    .foregroundColor(.primary)
            }

            if plan.isPersonal == true {
                Label("Personal Plan", systemImage: "person")
                    .font(.caption2)
                    .foregroundColor(.secondary)
            }
        }
        .padding(.vertical, 4)
    }
}
