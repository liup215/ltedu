# Tech Context: LT-Edu

## 1. Backend Technology Stack

- **Language**: Go
- **Database**: SQLite
- **Web Framework**: Gin
- **Architecture Pattern**: Controller-Service-Repository Layered Architecture
- **Data Access**: Repository Pattern (Custom Implementation)
- **ORM**: GORM (with gorm/gen for type-safe queries)
- **Authentication**: JWT
- **Logging**: Zap
- **Async Tasks**: NSQ (potential)
- **Cloud Services**: Qiniu Cloud SDK (object storage), Aliyun Bailian SDK (AI/LLM)
- **Containerization**: Docker

## 2. Frontend Technology Stack

- **Framework**: Vue.js 3
- **Language**: TypeScript
- **State Management**: Pinia
- **Routing**: Vue Router
- **HTTP Client**: Axios
- **Styling**: TailwindCSS
- **Rich Text Editing**: Quill, TinyMCE
- **Build Tool**: Vite
- **Package Manager**: npm

## 3. Functional Mapping

- **User & Access Management**: JWT, Pinia, Vue Router, Gin, GORM
- **System Administration**: Admin APIs, Pinia, Vue components, Gin controllers
- **Educational Structure Management**: GORM models, Vue forms, API services
- **Course Management**: Backend course/video services, frontend course views
- **Examination & Assessment**: Question/paper services, exam builder UI, AI integration
- **Content & Media Management**: Qiniu SDK, document/image/video APIs, upload components
- **AI & Advanced Features**: Aliyun Bailian SDK, vocabulary APIs, smart features

## 4. Development Environment

- **Reverse Proxy**: Traefik
- **Container Orchestration**: Docker Compose
- **Workspace**: VS Code workspace (`ltedu.code-workspace`)
