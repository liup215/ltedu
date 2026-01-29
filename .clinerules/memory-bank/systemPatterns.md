# System Patterns: LT-Edu

## 1. Architecture Overview

LT-Edu uses a modular client-server architecture, organized by functional modules:

- **Frontend (Vue.js SPA)**: User interface for all roles, built from reusable components and views mapped to functional modules (User, Admin, Content, Course, Exam, Media, AI).
- **Backend (Go + Gin)**: RESTful API with layered design (Model, DAO, Service, Controller), each service mapped to a functional module.
- **Database (SQLite + GORM)**: Data persistence for all entities (users, courses, exams, media, etc.).
- **Reverse Proxy (Traefik)**: Routes traffic to frontend/backend containers.
- **Cloud Services**: Qiniu for media storage, Aliyun Bailian for AI features.

## 2. Functional Module Patterns

- **User & Access Management**: JWT authentication, Pinia state, Vue Router guards, Gin controllers, GORM models.
- **System Administration**: Admin dashboard, user management, system settings, teacher application review (frontend views, backend services).
- **Educational Structure Management**: Organisation, qualification, syllabus, chapter management (forms, APIs, models).
- **Course Management**: Course/video management (pending frontend views, backend services).
- **Examination & Assessment**: Question bank, exam paper builder, paper series/code, past/random papers (SPA views, backend services).
- **Content & Media Management**: Document/image/video/slide management, Qiniu integration (upload components, APIs).
- **AI & Advanced Features**: Vocabulary sets, AI-powered features (Aliyun Bailian SDK, smart APIs).

## 3. Deployment & Infrastructure

- **Containerization**: All services run in Docker containers, orchestrated by Docker Compose.
- **CI/CD**: Build and deploy steps for frontend and backend.
- **Cloud Integration**: Specialized SDKs for storage and AI.

## 4. Design Principles

- **Function-Oriented Separation**: Each module is independently developed and documented.
- **Documentation-First**: Memory Bank tracks architecture and progress by functional module.
- **Scalability & Maintainability**: Modular design enables easy feature expansion and onboarding.
