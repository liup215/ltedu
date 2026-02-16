# Project Progress: LT-Edu

## 1. Current Status

- **[2026-02-16] Quill Editor Table Library Migration & UI Refinement**: Migrated from `quill-better-table` to `quill-table-better` for advanced table functionality, added `tableWidth` prop, keyboard bindings, and refined UI font styling (16px base, normal font weights).**
- **[2026-02-02] MCP tools expanded and question repository improved with safer update operations.**
- **[2026-01-31] Practice question retrieval logic implemented and grading fixes applied.**
- **[2026-01-30] System architecture standardization, repository pattern implementation, and deployment automation.**
  - Formalized Web + Backend unified architecture.
  - Implemented dedicated Repository module.
  - Refined Docker configurations and GitHub Actions workflows.
- [2025-09-13] All pages, dialogs, and popups in ltedu-web are fully internationalized (Chinese/English). All static text uses $t(key) with en.ts/zh.ts maintained.

**Phase: Infrastructure & Backend Refinement**

The project is moving towards better performace, query capabilities, and automated deployment processes.

## 2. What Works (by Website Feature)

1. **Frontend Framework**: Vue.js + TypeScript structure, TailwindCSS UI, routing and modular views are stable.
   - Global error handling is implemented using Notivue, providing automatic popup notifications for all API errors.
2. **Registration & Login**: Fully implemented JWT-based authentication, registration, password management.
   - **[2025-08-26] Unified backend verification logic for email/phone: SendCode API now uses a central service to handle both email and phone verification codes, with captcha validation and type-based dispatch. Improves maintainability and extensibility.**
   - **[2025-08-26] Register.vue email verification code countdown implemented: "Send Code" button disables after sending, shows countdown for 60 seconds, then re-enables. Improves registration security and user experience.**
   - **[2025-08-25] Image verification code (captcha) feature is fully implemented. This includes the backend service for generation/validation and frontend integration on the login page, with an automatic refresh on login failure.**
   - **[2025-08-20] Change Password feature full-stack implementation: backend unified user retrieval, registered change password route, tokenSalt field design, ensures all devices require re-login after password change; frontend Profile.vue change password form, fields aligned with backend (oldPassword/newPassword), new password confirmation validation, error handling and API errors displayed in UI. All model/service naming, API calls, and form logic follow LTEDU conventions.
3. **User Management**: Profile management, role assignment, and user CRUD are functional.
   - **[2025-08-27] Admins can grant 1 month VIP to users via UserManagement.vue. VIP expiry date is shown for each user. Role column removed.**
   - **[2025-08-20] Remove Admin feature implemented: User Management page now supports removing admin privileges via "Remove Admin" button (English UI). Backend API and route /api/v1/user/removeAdmin fully integrated.
4. **Organisation Management**: Organisation CRUD, cascading selects in forms, fully integrated.
5. **Qualification Management**: Qualification CRUD, linked to organisations, cascading selects operational.
6. **Syllabus Management**: Syllabus CRUD, linked to qualifications, cascading selects operational.
7. **Practice / Quick Practice**: 
   - **[2025-08-17] QuickPractice feature is now fully complete, including backend and frontend. All requirements for question rendering, grading, cascading selects, strict ApiResponse usage, and reset functionality are implemented and verified.**
   - **[2025-08-27] Chapter filter is now implemented in QuickPractice.vue. Users can select chapters after syllabus selection; selected chapterIds are sent to the backend for question filtering. UI and logic match admin/QuestionManagement.vue.**
   - Fully implemented. Backend returns joined fields for questions, frontend maps student answers per question part, reset button added to clear all data and restart session. Cascading selects (Organisation → Qualification → Syllabus), chapter filter, question rendering, grading, strict ApiResponse usage.
7.1 **Past Paper Practice**:
   - **[2025-08-28] Past Paper Practice feature (PaperPractice.vue) implemented and verified.**
   - Users select filters and a past paper, then start a practice session.
   - Backend returns question IDs for the selected paper; frontend loads and renders questions.
   - PracticeSidebar.vue reused for navigation and grading.
   - Implementation uses selectedPastPaperId directly, removing getPaperId helper.
   - Feature tested and confirmed working.
8. **Course Management**: Backend services complete; frontend views pending.
9. **Exam Management**: Question bank, exam paper builder, random/past paper workflows implemented and integrated.
   - **[2025-08-17] ExamPaperForm.vue now displays questionContent for each question part using the same format and logic as QuestionManagement.vue, including options, correct answers, and rich text editor for short answer questions. All question types are rendered consistently.**
   - **[2025-08-24] Admin QuestionManagement page now supports searching questions by paper name. Backend and frontend fully integrated.**
   - **[2025-08-24] ExamPaperForm page now supports searching questions by paper name. UI and backend logic match QuestionManagement.**
   - **[2025-08-27] ExamPaperForm.vue chapter selection refactored to match QuickPractice.vue: now uses a tree selector (ChapterOption component) for batch selection, and the redundant "All Chapters" dropdown filter has been removed. All prompt messages are now in English.**
10. **Media Management**: Document/image/video/slide management works; Qiniu cloud integration in place.
11. **AI Features**: Vocabulary sets and initial AI-powered features (question generation, recommendations) are integrated.
12. **Homepage**: 
   - **[2025-08-19] Homepage redesigned to highlight Quick Practice for students and Exam Paper Builder for teachers. Direct entry points for both roles are provided, and support for major international exam syllabuses is emphasized.**
   - **[2025-08-20] Home and quick practice pages are now public, no login required.**
13. **Donation & VIP**:
   - **[2025-08-27] Global donation popup randomly appears (10% chance) on route change, suppressed for VIP users and admins. Popup and donation page remind users to include username in donation remark for advanced permissions.**
14. **Backend & Infrastructure**:
   - **[2026-02-16] Quill Editor Table Library Migration**: Migrated from `quill-better-table` to `quill-table-better` for advanced table functionality, added `tableWidth` prop and keyboard bindings.
   - **[2026-02-02] MCP Tools**: Added tools for managing past papers, qualifications, questions, and syllabuses via Model Context Protocol.
   - **[2026-02-02] Question Repository Safety**: Fixed repository to exclude specific fields from updates, preventing accidental data overwrites.
   - **[2026-01-31] Practice Question Retrieval**: Implemented question retrieval logic for quick and paper practice generation; fixed grading logic in GradePracticeSubmission.
   - **[2026-01-30] System Architecture**: Documented the integral Web + Backend structure and layered architecture.
   - **[2026-01-30] Repository Module**: Implemented dedicated repository layer for data access abstraction and standardized query logic.
   - **[2026-01-30] Chapter Filter Optimization**: Implemented `filterRoot` option in `ChapterQuery` logic.
   - **[2026-01-30] Repository Performance**: Enhanced generic repository and API functionalities with additional Preloads and support for more query parameters.
   - **[2026-01-30] Deployment Automation**: Refactored GitHub Actions workflow and deployment scripts for streamlined Docker-based deployment.

## 3. What's Left to Build

1. **Course Management**: Complete and polish frontend views for course/video management.
2. **Admin Dashboard Enhancements**: Add analytics, batch operations, advanced controls.
3. **AI Features**: Expand smart content, feedback, and assessment capabilities.
4. **Testing & QA**: Comprehensive unit, integration, and E2E tests for all modules.
5. **Performance & Security**: Optimize API, frontend, and database for scalability and security.

## 4. Known Issues

- None at this stage.

## 5. Evolution of Decisions

- Documentation and architecture now follow a feature-oriented structure for clarity and maintainability.
- Modular design enables rapid onboarding and feature expansion.
- **[2025-08-17] ExamPaperForm.vue questionContent rendering refactor complete and documented.**
- **[2025-08-27] ExamPaperForm.vue chapter selection refactored to match QuickPractice.vue: now uses a tree selector (ChapterOption component) for batch selection, and the redundant "All Chapters" dropdown filter has been removed. All prompt messages are now in English.**

---

## [2025-08-28] Past Paper Practice feature complete (PaperPractice.vue)
- Users can select a past paper and start a practice session.
- Backend returns question IDs for the selected paper; frontend loads and renders questions.
- PracticeSidebar.vue reused for navigation and grading.
- Implementation uses selectedPastPaperId directly, removing getPaperId helper.
- Feature tested and confirmed working.

## [2025-08-28] Quick Practice pagination & sidebar refactor (ing me)

- Backend /practice/quick API now returns only question ID array; frontend loads question details by ID per page. Data models updated.
- Frontend QuickPractice.vue implements question ID pagination, only shows current question, "Previous"/"Next" buttons centered to avoid sidebar overlap.
- Sidebar PracticeSidebar.vue displays question numbers, answer status, grading result; supports jump and close; popup height is half page, fixed at bottom right.
- Grading result only shows current question's answer and score, using computed resultItem, fully type-safe.
- UX improvements: sidebar close/show button, centered navigation buttons and question number, grading result syncs with question page.
- Key tech: Vue3 + TS type safety, componentized sidebar/main page linkage, grading result sync with question number.
- Next steps: integration testing, mobile adaptation, further UI/UX polish, more features as needed.

---
