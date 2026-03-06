import SwiftUI
import UserNotifications

@main
struct LTEduApp: App {
    @UIApplicationDelegateAdaptor(AppDelegate.self) var appDelegate
    @StateObject private var authViewModel = AuthViewModel()
    @StateObject private var notificationManager = PushNotificationManager.shared

    var body: some Scene {
        WindowGroup {
            ContentView()
                .environmentObject(authViewModel)
                .environmentObject(notificationManager)
                .onAppear {
                    notificationManager.requestAuthorization()
                }
        }
    }
}
