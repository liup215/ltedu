# LTEdu HarmonyOS Native App

A native HarmonyOS application for the LTEdu platform, built with ArkTS and ArkUI. Provides a full AI-native learning experience leveraging HarmonyOS distributed capabilities.

## Features

- **AI Chat Interface** – Primary interaction surface powered by the LTEdu AI backend
- **Practice Sessions** – Quick practice, paper-based practice, and AI-powered answer feedback
- **Learning Plans** – View personal and class learning plans (long/mid/short term)
- **User Profile & Settings** – Account management and API server configuration
- **Atomic Service Widget** – Home-screen widget for quick access without full app launch
- **HarmonyOS Distributed** – Multi-device collaboration via HarmonyOS data management

## Project Structure

```
harmonyos/
├── AppScope/                        # Application-level resources and config
│   ├── app.json5                    # App bundle configuration
│   └── resources/base/
│       └── element/string.json      # App-level strings
├── entry/                           # Main entry module
│   ├── build-profile.json5          # Module build configuration
│   ├── oh-package.json5             # Module dependencies
│   └── src/main/
│       ├── module.json5             # Module manifest (abilities, permissions)
│       ├── ets/
│       │   ├── entryability/
│       │   │   ├── EntryAbility.ets          # App entry point, auth routing
│       │   │   └── FormExtensionAbility.ets  # Atomic Service widget backend
│       │   ├── pages/
│       │   │   ├── LoginPage.ets             # Authentication page
│       │   │   ├── RegisterPage.ets          # Registration with email verification
│       │   │   ├── MainPage.ets              # Tab navigator (Chat/Practice/Plans/Profile)
│       │   │   ├── ChatPage.ets              # AI chat interface (primary page)
│       │   │   ├── PracticePage.ets          # Practice mode selection and session
│       │   │   ├── PracticeDetailPage.ets    # Paper-based practice with AI feedback
│       │   │   ├── LearningPlanPage.ets      # Learning plan viewer
│       │   │   └── ProfilePage.ets           # User profile and settings
│       │   ├── components/
│       │   │   ├── ChatMessageItem.ets       # Chat bubble component
│       │   │   ├── QuestionCard.ets          # Question display with answer selection
│       │   │   └── CommonComponents.ets      # LoadingOverlay, ErrorBanner, EmptyState
│       │   ├── services/
│       │   │   ├── ApiService.ets            # HTTP client wrapping @ohos.net.http
│       │   │   ├── AuthService.ets           # Login, register, logout
│       │   │   ├── AiService.ets             # AI chat and answer checking
│       │   │   ├── PracticeService.ets       # Practice and grading
│       │   │   ├── LearningPlanService.ets   # Learning plan CRUD
│       │   │   └── StorageService.ets        # HarmonyOS Preferences wrapper
│       │   ├── models/
│       │   │   └── ApiModels.ets             # Shared TypeScript interfaces
│       │   └── widget/pages/
│       │       └── WidgetCard.ets            # Atomic Service home-screen widget
│       └── resources/
│           └── base/
│               ├── element/
│               │   ├── string.json           # UI strings
│               │   └── color.json            # Design token colours
│               └── profile/
│                   ├── main_pages.json       # Page routing configuration
│                   └── form_config.json      # Widget form configuration
├── build-profile.json5              # Root build configuration
├── oh-package.json5                 # Root package manifest
└── hvigor/
    └── hvigor-config.json5          # Hvigor build tool configuration
```

## Development Setup

### Prerequisites

- [DevEco Studio](https://developer.harmonyos.com/en/develop/deveco-studio) 4.1 or higher
- HarmonyOS SDK API 12 or higher
- A HarmonyOS device or emulator for testing

### Getting Started

1. **Open the project** – Launch DevEco Studio and open the `harmonyos/` directory.

2. **Configure the API URL** – The default backend URL is `https://api.ltedu.com`. You can update it in the app's Profile → Settings screen.

3. **Build & Run** – Click **Run** in DevEco Studio to build and deploy to a connected device or emulator.

### Configuration

The app communicates with the LTEdu backend API. Ensure the following is reachable:

| Endpoint | Purpose |
|---|---|
| `POST /api/v1/login` | User authentication |
| `POST /api/v1/register` | User registration |
| `GET  /api/v1/user` | Get current user profile |
| `POST /api/v1/practice/quick` | Quick practice questions |
| `POST /api/v1/practice/paper` | Paper-based practice |
| `POST /api/v1/practice/grade` | Grade submitted answers |
| `POST /api/v1/learning-plan/list` | List learning plans |
| `POST /api/v1/learning-plan/byId` | Get a single plan |
| `POST /api/v1/ai/chat` | AI chat messages |
| `POST /api/v1/ai/check-vocabulary` | AI vocabulary feedback |
| `POST /api/v1/ai/check-question` | AI question answer feedback |

See the main [API documentation](../docs/swagger.yaml) for full endpoint details.

## Architecture

```
┌────────────────────────────────────────────┐
│              HarmonyOS App                 │
│                                            │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐ │
│  │ ChatPage │  │Practice  │  │Learning  │ │
│  │ (primary)│  │  Page    │  │Plan Page │ │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘ │
│       │              │              │       │
│  ┌────▼──────────────▼──────────────▼────┐ │
│  │           Service Layer               │ │
│  │  AiService  PracticeService  Plan...  │ │
│  └────────────────────┬──────────────────┘ │
│                       │                    │
│  ┌────────────────────▼──────────────────┐ │
│  │  ApiService (HTTP + Auth headers)     │ │
│  └────────────────────┬──────────────────┘ │
│                       │                    │
│  ┌────────────────────▼──────────────────┐ │
│  │    StorageService (HOS Preferences)   │ │
│  └───────────────────────────────────────┘ │
└────────────────────────────────────────────┘
                        │
                        ▼
              LTEdu Backend API
             (Go + Gin + PostgreSQL)
```

### Key Design Decisions

- **Singleton services** – Each service class exposes a module-level singleton instance to avoid redundant initialisation and share state (e.g., conversation history in `AiService`).
- **StorageService init** – `StorageService.init()` must be called in `EntryAbility.onCreate()` before any page loads to ensure Preferences is available.
- **Auth-first routing** – `EntryAbility.onWindowStageCreate()` checks the stored token and redirects to `LoginPage` or `MainPage` accordingly.
- **HarmonyOS Preferences** – Used instead of plain file I/O to leverage the OS-managed, encrypted key-value store for auth tokens.

## Atomic Service Widget

The app includes a 2×2 home-screen widget (`LTEduWidget`) that shows study status and provides one-tap access to the AI chat or practice session. The widget is defined in `form_config.json` and rendered by `WidgetCard.ets`.

## Permissions

| Permission | Reason |
|---|---|
| `ohos.permission.INTERNET` | Backend API communication |
| `ohos.permission.GET_NETWORK_INFO` | Network availability checks |

## Huawei AppGallery Compliance

- Bundle name: `com.ltedu.app`
- Min API version: 12
- Target API version: 12
- Device types: phone, tablet
- Distribution filter: excludes circular screen (watch)

## Testing

Unit tests are in `entry/src/ohosTest/ets/tests/AppTest.ets` and cover:

- `AiService` conversation history management
- Date formatting utilities
- Plan type colour mapping
- Practice scoring calculations
- Input validation for login and registration

Run tests via DevEco Studio → **Run** → **Run 'ohosTest'**.
