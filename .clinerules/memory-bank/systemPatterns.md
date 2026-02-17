# System Patterns: LT-Edu

## 1. Architecture Overview

LT-Edu follows a unified system architecture where both the Web application and Backend API are integral parts of the system monorepo.

- **System Structure**:
    - **Web (Frontend)**: Vue.js + TypeScript SPA located in `web/`, serving as the user interface.
    - **Server (Backend)**: Go + Gin API located in `server/`, `service/`, `repository/`, providing business logic and data access.
- **Database**: SQLite + GORM for data persistence.
- **Cloud Services**: Qiniu for media storage, Aliyun Bailian for AI features.

## 2. Technical Architecture implemented methods

### Backend Layered Architecture
The backend follows a strict layered architecture to ensure separation of concerns:

- **Controller Layer (`server/api/v1/`)**:
  - Handles HTTP requests and responses.
  - Validates input and calls the Service layer.
  - Returns formatted responses (Success/Error).

- **Service Layer (`service/`)**:
  - Encapsulates business logic.
  - Orchestrates operations between multiple Repositories.
  - Initialized as singleton instances (e.g., `service.UserSvr`).

- **Repository Layer (`repository/`)**:
  - **New Functional Implementation**: A dedicated `repository` module handles all data access.
  - Defines interfaces (e.g., `IUserRepository`) and implementations (e.g., `userRepository`).
  - Abstracts GORM queries and database operations.
  - Supports Preloading and complex filtering (e.g., `ChapterQuery` with `filterRoot`).
  - Initialized globally via `repository.InitRepositories(db)`.

- **Model Layer (`model/`)**:
  - Defines GORM structs and database schema.

### Frontend Architecture
The frontend (`web/src/`) follows a component-based architecture:

- **Views**: Page-level components mapping to routes.
- **Components**: Reusable UI elements.
- **Services**: API client wrappers interacting with the backend.
- **Stores**: Pinia state management for global data.

## 3. Functional Module Patterns

- **User & Access Management**: JWT authentication, Pinia state, Vue Router guards, Gin controllers, GORM models.
- **System Administration**: Admin dashboard, user management, system settings, teacher application review (frontend views, backend services).
- **Educational Structure Management**: Organisation, qualification, syllabus, chapter management (forms, APIs, models).
- **Syllabus Navigator (Learning Engine)**:
  - Built on **Syllabus/Chapter/Question/Paper** as the knowledge map foundation.
  - Core concepts: **Goal** (learning project), **Diagnostic** (calibration), **User Knowledge State** (mastery + retention/stability), **Task Stream** (daily task cards).
  - Planning structure: **7-day hard plan** + **4-week soft plan**, with **plan versioning** and explainable adjustments.
  - Task types: Learn / Drill / Review (SRS) / Test / Mock or Past.
  - Update strategy: event-driven writes (Attempt/TaskLog) + async partial replanning + scheduled backfill/maintenance jobs.
- **Course Management**: Course/video management (pending frontend views, backend services).
- **Examination & Assessment**: Question bank, exam paper builder, paper series/code, past/random papers (SPA views, backend services).
- **Content & Media Management**: Document/image/video/slide management, Qiniu integration (upload components, APIs).
- **AI & Advanced Features**: Vocabulary sets, AI-powered features (Aliyun Bailian SDK, smart APIs).

## 4. Deployment & Infrastructure

- **Containerization**: All services run in Docker containers, orchestrated by Docker Compose.
- **CI/CD**: GitHub Actions workflows handle automated build and deployment processes (Docker image generation, deployment scripts).
- **Cloud Integration**: Specialized SDKs for storage and AI.

## 5. Design Principles

- **Function-Oriented Separation**: Each module is independently developed and documented.
- **Documentation-First**: Memory Bank tracks architecture and progress by functional module.
- **Scalability & Maintainability**: Modular design enables easy feature expansion and onboarding.
