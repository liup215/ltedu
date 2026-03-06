// swift-tools-version: 5.9
// This Package.swift enables building and testing the LTEdu iOS app logic
// using Swift Package Manager for CI/CD pipelines and Linux environments.
// Full iOS app features (SwiftUI, UIKit, CoreData) require Xcode on macOS.

import PackageDescription

let package = Package(
    name: "LTEdu",
    platforms: [
        .iOS(.v16),
        .macOS(.v13)
    ],
    products: [
        .library(
            name: "LTEduCore",
            targets: ["LTEduCore"]
        )
    ],
    targets: [
        .target(
            name: "LTEduCore",
            path: "LTEdu",
            exclude: [
                "App/LTEduApp.swift",
                "App/AppDelegate.swift",
                "App/ContentView.swift",
                "Views",
                "CoreData",
                "Notifications/PushNotificationManager.swift"
            ],
            sources: [
                "Network/APIClient.swift",
                "Network/APIEndpoints.swift",
                "Network/NetworkModels.swift",
                "Services/AuthService.swift",
                "Services/Services.swift",
                "Utilities/KeychainManager.swift",
                "Utilities/Extensions.swift"
            ]
        ),
        .testTarget(
            name: "LTEduTests",
            dependencies: ["LTEduCore"],
            path: "../LTEduTests"
        )
    ]
)
