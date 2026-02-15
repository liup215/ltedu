# Active Context: LT-Edu

## 1. Current Focus

- **System Architecture & Repository Pattern**: Formalizing the unified monorepo structure (Web + Server) and the dedicated Repository module for data access abstraction.
- **Backend Optimization**: Improving query capabilities (e.g., `filterRoot` in ChapterQuery) and adding Preloads to repositories for better performance.
- **Deployment & DevOps**: Refining Docker configurations and GitHub Actions workflows for automated build and deployment.
- **MCP Integration**: Maintaining and updating Model Context Protocol (MCP) integrations.
- **Maintenance**: Keeping documentation aligned with backend and infrastructure updates.

**Phase: Infrastructure & Backend Refinement**

The project is currently focusing on backend robustness, query optimization, and streamlining the deployment process.

## 2. Main Website Features

1. **Frontend Framework**
   - Vue.js + TypeScript, modular structure, TailwindCSS for UI.
2. **Registration & Login**
   - JWT-based authentication, registration, password management.
   - **[2025-08-25] Image verification code (captcha) feature implemented for the login page. Backend generates and validates the captcha, and the frontend displays it, including a refresh mechanism on login failure.**
3. **User Management**
   - Profile management, role assignment (Student, Teacher, Admin).
   - Admin privileges can now be revoked directly from the User Management page (Remove Admin button, full-stack implementation).
   - [2025-08-27] Admins can now grant 1 month VIP to any user via UserManagement.vue. VIP expiry date is displayed for each user. Role column removed for clarity.
4. **Organisation Management**
   - CRUD for organisations, cascading selects in forms.
5. **Qualification Management**
   - CRUD for qualifications, linked to organisations.
6. **Syllabus Management**
   - CRUD for syllabuses, linked to qualifications.
7. **Practice / Quick Practice**
   - Student quick practice feature, cascading selects (Organisation → Qualification → Syllabus), question rendering, grading, strict ApiResponse usage.
   - [2025-08-27] Chapter filter added to QuickPractice.vue. Users can select multiple chapters after choosing a syllabus; selected chapterIds are sent to the backend for question filtering. UI and logic follow the admin/QuestionManagement.vue pattern and LTEDU frontend conventions.
   - **[2025-08-28] Quick Practice pagination & sidebar refactor (ing me):**
     - Backend /practice/quick API now returns only question ID array; frontend loads question details by ID per page. Data models updated.
     - Frontend QuickPractice.vue implements question ID pagination, only shows current question, "Previous"/"Next" buttons centered to avoid sidebar overlap.
     - Sidebar PracticeSidebar.vue displays question numbers, answer status, grading result; supports jump and close; popup height is half page, fixed at bottom right.
     - Grading result only shows current question's answer and score, using computed resultItem, fully type-safe.
     - UX improvements: sidebar close/show button, centered navigation buttons and question number, grading result syncs with question page.
     - Key tech: Vue3 + TS type safety, componentized sidebar/main page linkage, grading result sync with question number.
     - Next steps: integration testing, mobile adaptation, further UI/UX polish, more features as needed.
7.1 **Past Paper Practice**
   - [2025-08-28] New feature: Past Paper Practice (PaperPractice.vue).
     - Users select filters (organisation, qualification, syllabus, year, paper code, series, past paper) and start a full past paper practice session.
     - Backend returns question ID array for the selected past paper; frontend loads question details per ID.
     - PracticeSidebar.vue reused for navigation, answer status, grading result.
     - Fully type-safe, strict ApiResponse usage, matches LTEDU frontend patterns.
     - Implementation: getPaperId removed, now uses selectedPastPaperId directly for backend requests.
     - Status: Feature implemented, tested, and verified working.
8. **Course Management**
   - Course creation/editing, video management, backend complete, frontend views in progress.
9. **Exam Management**
   - Question bank, exam paper builder, random/past paper workflows.
   - **[2025-08-17] Refactored ExamPaperForm.vue to display questionContent for each question part using the same format and logic as QuestionManagement.vue, including options, correct answers, and rich text editor for short answer questions. All question types are rendered consistently.**
   - **[2025-08-24] Admin QuestionManagement page now supports searching questions by paper name. Backend and frontend fully integrated.**
   - **[2025-08-24] ExamPaperForm page now supports searching questions by paper name. UI and backend logic match QuestionManagement.**
   - **[2025-08-27] ExamPaperForm.vue chapter selection refactored to match QuickPractice.vue: now uses a tree selector (ChapterOption component) for batch selection, and the redundant "All Chapters" dropdown filter has been removed. All prompt messages are now in English.**
10. **Media Management**
    - Document/image/video/slide management, Qiniu cloud integration.
11. **AI Features**
    - Vocabulary sets, AI-powered question generation, recommendations.

## 3. Recent Changes

- **[2026-02-16] Quill Editor Enhancement**:
  - Added `quill-better-table` integration for enhanced table functionality in the Quill rich text editor.
  - Improves content creation capabilities for courses, questions, and other rich text content.
- **[2026-02-02] MCP Tools & Question Repository Improvements**:
  - Added MCP tools for managing past papers, qualifications, questions, and syllabuses.
  - Enhanced argument parsing in tools and added debug logging for question editing.
  - Fixed question repository to exclude specific fields from updates for safer data operations.
- **[2026-01-31] Practice Question Retrieval Logic**:
  - Implemented question retrieval logic for quick and paper practice generation.
  - Fixed question retrieval logic in GradePracticeSubmission for accurate grading.
  - Improved error handling when creating or finding app configuration keys.
- **[2026-01-30] System Architecture & Repository Pattern**:
  - Formalized the unified system architecture containing both Web (Frontend) and Server (Backend).
  - Implemented a dedicated `repository` module with strictly typed interfaces (e.g., `IUserRepository`) to abstract GORM operations and support complex query logic.
- **[2026-01-30] Backend Enhancements:**
  - Added `filterRoot` option to `ChapterQuery` and enhanced filtering logic in `chapter_repository`.
  - Enhanced generic repository and API functionalities with additional Preloads and query parameters support.
- **[2026-01-30] Deployment & DevOps:**
  - Refactored GitHub Actions workflow for Docker build and deployment.
  - Updated deployment scripts to manage Docker Compose and directory structures better.
  - Added and updated Docker configuration files for build processes.
- **[2026-01-30] MCP Integration:**
  - Updated MCP endpoint URL in token management to use the correct API path.
- **[2025-08-28] Past Paper Practice feature created (PaperPractice.vue):**
  - Users can now select a past paper and start a practice session.
  - Backend returns question IDs for the selected paper; frontend loads and renders questions.
  - PracticeSidebar.vue reused for navigation and grading.
  - Implementation uses selectedPastPaperId directly, removing getPaperId helper.
  - Feature tested and verified.
- **[2025-08-27] User Management: Added "Grant VIP" button for each user (admin only), calls backend API to grant 1 month VIP. VIP expiry date now shown in user table. Role column removed.**
- **[2025-08-27] Global donation popup: Randomly appears on route change (10% chance), suppressed for VIP users and admins. Popup and donation page both remind users to include username in donation remark for advanced permissions.**
- **[2025-08-27] QuickPractice.vue: Added chapter filter. Users can now select chapters for practice; chapterIds are passed to the backend. UI and logic match admin/QuestionManagement.vue.**
- **[2025-08-27] ExamPaperForm.vue: Chapter selection refactored to match QuickPractice.vue, using a tree selector for batch selection. Redundant chapter dropdown filter removed. All prompt messages are now in English.**
- **[2025-08-26] Unified Verification Logic (Email/Phone):**
  - Backend now uses a unified verification service for sending codes to both email and phone targets. The SendCode API validates captcha, determines verification type, and delegates to a central GenerateAndSendCode method, improving maintainability and extensibility.
- **[2025-08-26] Register.vue Email Verification Code Countdown:**
  - Frontend now disables the "Send Code" button after sending, displays a countdown for 60 seconds, and re-enables the button when finished. This improves registration security and user experience. Implementation follows LTEDU frontend conventions.
- **[2025-08-25] Image Captcha Feature Full-Stack Implementation:**
  - **Backend**: Implemented a new captcha service (`captcha.go`) using `base64Captcha` with an in-memory store. A new public API endpoint (`/api/v1/captcha`) was created to serve captcha images. The login authentication flow (`auth.go`) was updated to require and validate the captcha.
  - **Frontend**: The `Login.vue` component was updated to display the captcha image and input field. Logic was added to fetch the captcha on component mount, on click, and automatically after a failed login attempt to improve user experience.
- **[2025-08-24] Admin QuestionManagement page: Added search by paper name. Backend model/service and frontend UI/service updated for full integration.**
- **[2025-08-24] ExamPaperForm page: Added search by paper name to question search. UI and backend logic match QuestionManagement.**
- **[2025-08-20] Remove Admin feature:** Added "Remove Admin" button to User Management page (Vue, English UI). Backend API and route /api/v1/user/removeAdmin implemented and registered. Admin privileges can now be revoked for any user.
- **[2025-08-20] Global Error Handling:** Implemented a global API error handling system using Notivue. All API errors are now caught by an Axios interceptor and displayed as user-friendly popup notifications.
- **[2025-08-19] Home page redesigned to highlight two core features: Quick Practice for students and Exam Paper Builder for teachers. The new homepage provides direct entry points for both roles and emphasizes support for major international exam syllabuses (Cambridge, IB, AP, SAT, A-Level, IGCSE, etc.).**
- **[2025-08-20] Change Password feature full-stack implementation:**
  - Backend: Unified user retrieval logic, registered change password route, tokenSalt field design, ensures all devices require re-login after password change.
  - Frontend: Profile.vue page added/updated change password form, fields aligned with backend (oldPassword/newPassword), added new password confirmation validation.
  - ChangePassword types merged into auth.model.ts.
  - Implemented changePassword method in authService.ts, calls /api/v1/change-password.
  - Parameter alignment: currentPassword → oldPassword, avoids backend validation errors.
  - Error handling: Form validation and API errors are displayed in the UI for better UX.
  - Routing: Home and quick practice pages are now public, no login required.
  - Code structure: All model/service naming, API calls, and form logic follow LTEDU conventions.
- **[2025-08-17] QuickPractice feature is now fully complete, including backend and frontend. All requirements for question rendering, grading, cascading selects, strict ApiResponse usage, and reset functionality are implemented and verified.**
- Quick Practice feature fully implemented:
  - Backend returns joined fields for questions.
  - Frontend maps student answers per question part.
  - Added reset button to clear all data and restart session.
- Refactored QuickPractice.vue to use ApiResponse wrappers for all service responses.
- Integrated cascading selects for Organisation, Qualification, Syllabus in QuickPractice.vue.
- Updated practiceService.ts and practice.model.ts for strict typing and response consistency.
- Improved documentation structure to follow feature-based organization.
- **[2025-08-17] ExamPaperForm.vue now matches QuestionManagement.vue for questionContent rendering, supporting all question types and correct answer display.**

## 4. Next Steps

- Ensure all usages of ApiResponse wrappers in practiceService and practice.model are consistent.
- Complete frontend views for course/video management.
- Expand AI features and analytics in admin dashboard.
- Continue documentation-first workflow for all new features.

## 5. Key Decisions & Patterns

- **Global Notifications**: Centralized API error handling via Axios interceptors and Notivue for consistent, non-intrusive user feedback.
- **Feature-Oriented Documentation**: All Memory Bank files now track by main website features.
- **Strict Typing & Response Wrappers**: All API responses use ApiResponse<T> for consistency.
- **Cascading Selects**: Organisation → Qualification → Syllabus pattern standardized in forms.
- **Modular Design**: Enables rapid onboarding and maintainability.
