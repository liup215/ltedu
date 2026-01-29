# Product Context: LT-Edu

## 1. Problem Space

The demand for flexible and accessible online education is rapidly growing. Traditional learning environments often face limitations in terms of scalability, content delivery, and personalized learning paths. LT-Edu aims to address these challenges by providing a modern, all-in-one platform for online learning and teaching.

## 2. Target Audience

- **Students**: Learners of all ages seeking to acquire new knowledge and skills through structured online courses.
- **Teachers/Instructors**: Educators who need a platform to create, manage, and deliver their educational content to a broad audience.
- **Administrators**: Staff responsible for managing the platform, overseeing content, and managing user accounts.

## 3. Core Features & Functional Goals

- **User Management**:
  - Distinct roles for Students, Teachers, and Administrators.
  - User registration and JWT-based authentication.
  - Dedicated account management section for users.
  - Admin-only section for platform management.
  - Registration and login now require image captcha and email verification code.
  - Registration uses unified API (`/api/v1/register`) with fields: username, email, password, passwordConfirm, verificationCode.
  - Mobile phone is optional and not validated during registration.
  - All API responses and prompts are in English.

- **Content & Course Management**:
  - Instructors can create and manage courses.
  - Support for rich text content creation using editors like **Quill** and **TinyMCE**.
  - Ability to generate and export content as `.docx` files.
  - Integration with **Qiniu Cloud** for storing and delivering media assets (videos, documents).

- **Assessment & Examinations**:
  - A comprehensive **Exam Paper Builder** for teachers to create, edit, and manage exams.
  - Ability for teachers to preview exam papers before publishing.
  - Students can take exams and view their results.

- **AI-Powered Features**:
  - The integration of the **Aliyun Bailian LLM SDK** suggests potential for AI-driven features, such as:
    - Automated question generation.
    - Smart content recommendations.
    - AI-assisted grading or feedback.

- **Student Experience**:
  - Students can browse the course catalog, enroll in courses, and track their learning progress.
  - A clean, responsive UI built with Vue.js and TailwindCSS.

## Recent Feature Update (2025-08-26)

- Added image captcha and email verification code for login and registration.
- Registration uses unified API (`/api/v1/register`) with fields: username, email, password, passwordConfirm, verificationCode.
- Mobile phone is optional and not validated during registration.
- All API responses and prompts are in English.
