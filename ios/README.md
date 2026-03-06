# LTEdu iOS App

Native iOS application for the LTEdu AI-powered education platform, built with SwiftUI.

## Features

- **AI Chat Interface** — Primary interaction via an AI assistant (powered by LTEdu's AI backend)
- **Course Catalog** — Browse and view all courses and videos
- **Practice Zone** — Quick practice sessions, custom difficulty settings, and past papers
- **Learning Plans** — View and track personalized learning plans
- **Knowledge Progress** — Visual progress tracking and spaced-repetition review scheduling
- **Offline Caching** — Core Data-backed caching for key resources
- **Push Notifications** — Study reminders and review-due alerts
- **Secure Token Storage** — JWT stored in iOS Keychain
- **iPhone & iPad** — Adaptive layout for all iOS form factors (iOS 16+)

## Architecture

```
ios/
├── Package.swift                  # Swift Package Manager manifest (CI/testing)
├── LTEdu/
│   ├── App/                       # App entry point & delegates
│   │   ├── LTEduApp.swift
│   │   ├── AppDelegate.swift
│   │   └── ContentView.swift
│   ├── Network/                   # HTTP client layer
│   │   ├── APIClient.swift        # URLSession-based generic client
│   │   ├── APIEndpoints.swift     # All endpoint constants & config
│   │   └── NetworkModels.swift    # Codable request/response models
│   ├── Services/                  # Business logic / API calls
│   │   ├── AuthService.swift      # Login, register, captcha
│   │   └── Services.swift         # Course, Practice, Class, LearningPlan, etc.
│   ├── ViewModels/                # ObservableObject state managers (MVVM)
│   │   ├── AuthViewModel.swift
│   │   └── ViewModels.swift       # Course, Practice, Chat, Knowledge, Plan VMs
│   ├── Views/                     # SwiftUI views
│   │   ├── Auth/AuthViews.swift   # Login, Register, shared input fields
│   │   ├── Home/                  # Home dashboard & tab container
│   │   ├── Chat/AIChatView.swift  # AI chat interface
│   │   ├── Courses/               # Course list & detail
│   │   ├── Practice/              # Practice session flow & results
│   │   └── Profile/               # Profile, settings, learning plans
│   ├── CoreData/                  # Offline persistence
│   │   ├── LTEdu.xcdatamodeld/    # Core Data model (User, Course, ChatSession)
│   │   └── PersistenceController.swift
│   ├── Notifications/
│   │   └── PushNotificationManager.swift
│   └── Utilities/
│       ├── KeychainManager.swift  # Secure JWT storage
│       └── Extensions.swift       # String, Date, View helpers
└── LTEduTests/
    └── LTEduTests.swift           # Unit tests for models, keychain, API config
```

**Pattern:** MVVM with SwiftUI + Combine  
**Persistence:** Core Data (offline cache) + Keychain (auth tokens)  
**Networking:** `URLSession` with `async/await` via `APIClient`

## Requirements

| Tool | Version |
|------|---------|
| Xcode | 15.0+ |
| Swift | 5.9+ |
| iOS Deployment Target | 16.0+ |
| macOS (build host) | 13.0+ |

## Getting Started

### 1. Create the Xcode Project

Because Xcode project files (`.xcodeproj`) are binary, this repository contains
only the Swift source files. To create the Xcode project:

```bash
# Open Xcode → File → New → Project
# Choose: iOS → App
# Product Name: LTEdu
# Bundle Identifier: com.yourorg.ltedu
# Interface: SwiftUI
# Language: Swift
# Include Core Data: checked

# Then drag all source folders from ios/LTEdu/ into the project
```

### 2. Configure the API Base URL

Set the `LTEduAPIBaseURL` key in your **Info.plist**:

```xml
<key>LTEduAPIBaseURL</key>
<string>https://your-ltedu-server.com/api</string>
```

Or set the environment variable `LTEduAPIBaseURL` in the Xcode scheme for local development.

For local development against a macOS-hosted server:
```
LTEduAPIBaseURL = http://localhost:80/api
```

### 3. Enable Push Notifications

In Xcode:
1. Select the `LTEdu` target → **Signing & Capabilities**
2. Add **Push Notifications** capability
3. Add **Background Modes** → check **Remote notifications**

### 4. Add the Core Data Model

Copy `ios/LTEdu/CoreData/LTEdu.xcdatamodeld/` to your Xcode project root and add it to the target.

### 5. Build & Run

Select a simulator (iPhone 15 Pro recommended) and press **⌘R**.

## Configuration

### API Base URL

Priority order:
1. `LTEduAPIBaseURL` environment variable (Xcode scheme)
2. `LTEduAPIBaseURL` key in `Info.plist`
3. Default: `http://localhost:80/api`

### Authentication

- JWT tokens are stored in the iOS Keychain under `com.ltedu.auth.token`
- Token expiry is tracked and validated before each request
- Expired tokens trigger automatic logout

## Testing

Run tests from Xcode with **⌘U** or via command line (network-independent tests only):

```bash
cd ios
swift test --filter LTEduTests
```

Tests cover:
- `APIConfigTests` — URL construction and defaults
- `NetworkModelsTests` — JSON encoding/decoding for all models
- `KeychainManagerTests` — Token save, retrieve, expiry, and clear
- `ExtensionsTests` — Email/password validation, HTML stripping
- `APIErrorTests` — Error descriptions
- `AnyCodableTests` — Type-erased encoding/decoding

## Offline Capabilities

The app caches the following data locally using Core Data:
- **User profile** — Persisted after login; displayed when offline
- **Courses** — Cached after first fetch; stale cache invalidated after 1 hour
- **AI Chat Sessions** — All conversations persisted locally as JSON blobs

## Push Notifications

Local notifications are scheduled for:
- **Study reminders** — Configurable time via Settings → Notifications
- **Review due alerts** — Triggered when spaced-repetition topics are due

Remote push notifications require backend APNs integration (not yet implemented).

## App Store Submission Checklist

- [ ] Set correct Bundle ID and provisioning profile
- [ ] Configure `Info.plist` privacy usage descriptions (if requesting permissions)
- [ ] Set `LTEduAPIBaseURL` to production server URL
- [ ] Enable App Transport Security exceptions if using HTTP (not recommended)
- [ ] Run on physical device to test Keychain and push notifications
- [ ] Test on both iPhone and iPad form factors
- [ ] Submit for App Store review via App Store Connect

## AI Integration

The current AI chat implementation uses a local mock response generator
(`AIChatViewModel.generateAIResponse`). To integrate a real AI backend:

1. The LTEdu backend supports **MCP tokens** for AI agent authentication
2. Replace the mock generator with a call to your AI service endpoint
3. Stream responses using `URLSession.bytes(for:)` for real-time display

## Contributing

See the main [repository README](../README.md) for contribution guidelines.
