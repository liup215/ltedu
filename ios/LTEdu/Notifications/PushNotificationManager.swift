import Foundation
import UserNotifications

// MARK: - PushNotificationManager

final class PushNotificationManager: NSObject, ObservableObject {
    static let shared = PushNotificationManager()

    @Published var isAuthorized = false
    @Published var pendingDeepLink: String?

    private var deviceToken: String?

    private override init() {
        super.init()
    }

    // MARK: - Authorization

    func requestAuthorization() {
        UNUserNotificationCenter.current().requestAuthorization(
            options: [.alert, .sound, .badge]
        ) { [weak self] granted, error in
            DispatchQueue.main.async {
                self?.isAuthorized = granted
            }
            if granted {
                DispatchQueue.main.async {
                    UIApplication.shared.registerForRemoteNotifications()
                }
            }
            if let error = error {
                print("Push notification authorization error: \(error.localizedDescription)")
            }
        }
    }

    func checkAuthorizationStatus() {
        UNUserNotificationCenter.current().getNotificationSettings { [weak self] settings in
            DispatchQueue.main.async {
                self?.isAuthorized = settings.authorizationStatus == .authorized
            }
        }
    }

    // MARK: - Device Token

    func setDeviceToken(_ data: Data) {
        let tokenString = data.map { String(format: "%02.2hhx", $0) }.joined()
        self.deviceToken = tokenString
        print("Device token: \(tokenString)")
        // TODO: Send token to server when backend push notification support is added
    }

    // MARK: - Local Notifications

    func scheduleStudyReminder(title: String, body: String, at date: Date) {
        let content = UNMutableNotificationContent()
        content.title = title
        content.body = body
        content.sound = .default
        content.categoryIdentifier = NotificationCategory.studyReminder

        let components = Calendar.current.dateComponents([.year, .month, .day, .hour, .minute], from: date)
        let trigger = UNCalendarNotificationTrigger(dateMatching: components, repeats: false)
        let identifier = "study-reminder-\(UUID().uuidString)"
        let request = UNNotificationRequest(identifier: identifier, content: content, trigger: trigger)

        UNUserNotificationCenter.current().add(request) { error in
            if let error = error {
                print("Failed to schedule study reminder: \(error.localizedDescription)")
            }
        }
    }

    func scheduleReviewReminder(for topic: String, dueDate: Date) {
        let content = UNMutableNotificationContent()
        content.title = "Review Due"
        content.body = "Time to review: \(topic)"
        content.sound = .default
        content.categoryIdentifier = NotificationCategory.reviewDue
        content.userInfo = ["topic": topic]

        let components = Calendar.current.dateComponents([.year, .month, .day, .hour], from: dueDate)
        let trigger = UNCalendarNotificationTrigger(dateMatching: components, repeats: false)
        let identifier = "review-\(topic.hashValue)"
        let request = UNNotificationRequest(identifier: identifier, content: content, trigger: trigger)

        UNUserNotificationCenter.current().add(request) { error in
            if let error = error {
                print("Failed to schedule review reminder: \(error.localizedDescription)")
            }
        }
    }

    func cancelAllNotifications() {
        UNUserNotificationCenter.current().removeAllPendingNotificationRequests()
    }

    func cancelNotification(identifier: String) {
        UNUserNotificationCenter.current().removePendingNotificationRequests(withIdentifiers: [identifier])
    }

    // MARK: - Deep Link Handling

    func handleNotificationResponse(_ response: UNNotificationResponse) {
        let userInfo = response.notification.request.content.userInfo
        if let deepLink = userInfo["deepLink"] as? String {
            DispatchQueue.main.async { [weak self] in
                self?.pendingDeepLink = deepLink
            }
        }
    }

    // MARK: - Notification Categories

    func registerNotificationCategories() {
        let studyAction = UNNotificationAction(
            identifier: NotificationAction.startStudy,
            title: "Start Studying",
            options: .foreground
        )
        let dismissAction = UNNotificationAction(
            identifier: NotificationAction.dismiss,
            title: "Dismiss",
            options: .destructive
        )

        let studyCategory = UNNotificationCategory(
            identifier: NotificationCategory.studyReminder,
            actions: [studyAction, dismissAction],
            intentIdentifiers: [],
            options: []
        )

        let reviewCategory = UNNotificationCategory(
            identifier: NotificationCategory.reviewDue,
            actions: [studyAction, dismissAction],
            intentIdentifiers: [],
            options: []
        )

        UNUserNotificationCenter.current().setNotificationCategories([studyCategory, reviewCategory])
    }
}

// MARK: - Constants

private enum NotificationCategory {
    static let studyReminder = "STUDY_REMINDER"
    static let reviewDue = "REVIEW_DUE"
}

private enum NotificationAction {
    static let startStudy = "START_STUDY"
    static let dismiss = "DISMISS"
}
