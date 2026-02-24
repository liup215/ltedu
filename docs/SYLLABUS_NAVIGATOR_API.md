# Syllabus Navigator API Documentation

## Overview

The Syllabus Navigator is a learning engine that helps students achieve their learning goals through a structured, data-driven approach. It provides:

- **Goal Management**: Set and track learning objectives with target exam dates
- **Knowledge State Tracking**: Monitor mastery and retention for each chapter
- **Task Streaming**: Daily task cards (Learn, Drill, Review, Test, Mock)
- **Spaced Repetition**: Intelligent review scheduling (SRS)
- **Progress Tracking**: Coverage, mastery, and stability metrics

## Architecture

### Data Models

1. **Goal** - Learning project with target syllabus and exam date
2. **KnowledgeState** - User's mastery and retention per chapter
3. **Task** - Daily task cards in the learning stream
4. **Attempt** - Learning activity logs (question attempts)
5. **TaskLog** - Task completion tracking

### Layer Structure

```
Controller (API v1)
    ↓
Service Layer
    ↓
Repository Layer
    ↓
Database (GORM)
```

## API Endpoints

All endpoints require JWT authentication via `Authorization: Bearer {token}` header.

### Goal Management

#### Create Goal
```
POST /api/v1/goal/create
```

**Request:**
```json
{
  "syllabusId": 1,
  "examDate": "2026-06-15T00:00:00Z",
  "targetScore": 85,
  "targetGrade": "A",
  "weeklyHours": 10,
  "mode": "self"
}
```

**Response:**
```json
{
  "code": 0,
  "message": "Goal created successfully!",
  "data": {
    "id": 1,
    "userId": 1,
    "syllabusId": 1,
    "examDate": "2026-06-15T00:00:00Z",
    "targetScore": 85,
    "targetGrade": "A",
    "weeklyHours": 10,
    "mode": "self",
    "status": "active",
    "diagnosticDone": false,
    "startDate": "2026-02-17T18:00:00Z"
  }
}
```

#### Update Goal
```
POST /api/v1/goal/edit
```

**Request:**
```json
{
  "id": 1,
  "examDate": "2026-07-01T00:00:00Z",
  "weeklyHours": 12
}
```

#### Get Goal by ID
```
POST /api/v1/goal/byId
```

**Request:**
```json
{
  "id": 1
}
```

#### List Goals
```
POST /api/v1/goal/list
```

**Request:**
```json
{
  "pageSize": 20,
  "pageIndex": 1,
  "status": "active"
}
```

#### Get Active Goals
```
POST /api/v1/goal/active
```

**Request:**
```json
{}
```

#### Delete Goal
```
POST /api/v1/goal/delete
```

**Request:**
```json
{
  "id": 1
}
```

#### Complete Diagnostic
```
POST /api/v1/goal/diagnostic/complete
```

**Request:**
```json
{
  "goalId": 1
}
```

### Knowledge State

#### Get Knowledge State by Chapter
```
POST /api/v1/knowledge-state/byChapter
```

**Request:**
```json
{
  "goalId": 1,
  "chapterId": 1
}
```

**Response:**
```json
{
  "code": 0,
  "message": "Data retrieved successfully!",
  "data": {
    "id": 1,
    "userId": 1,
    "goalId": 1,
    "chapterId": 1,
    "masteryLevel": 0.75,
    "stabilityScore": 0.6,
    "correctCount": 15,
    "totalCount": 20,
    "consecutiveCorrect": 3,
    "consecutiveWrong": 0,
    "isCovered": true,
    "lastPracticeAt": "2026-02-17T17:30:00Z",
    "nextReviewAt": "2026-02-20T17:30:00Z"
  }
}
```

#### List Knowledge States
```
POST /api/v1/knowledge-state/list
```

**Request:**
```json
{
  "goalId": 1
}
```

#### Get Progress
```
POST /api/v1/knowledge-state/progress
```

**Request:**
```json
{
  "goalId": 1
}
```

**Response:**
```json
{
  "code": 0,
  "message": "Data retrieved successfully!",
  "data": {
    "coverage": 65.5,
    "mastery": 72.3,
    "stability": 68.9,
    "chapterStates": [
      {
        "chapterId": 1,
        "chapterName": "Introduction",
        "masteryLevel": 0.85,
        "stabilityScore": 0.75,
        "isCovered": true
      }
    ]
  }
}
```

#### Get Due for Review
```
POST /api/v1/knowledge-state/due-review
```

**Request:**
```json
{
  "goalId": 1
}
```

### Task Management

#### Create Task
```
POST /api/v1/task/create
```

**Request:**
```json
{
  "goalId": 1,
  "type": "drill",
  "targetDate": "2026-02-18T00:00:00Z",
  "chapterId": 1,
  "title": "Practice: Chapter 1",
  "description": "Complete 10 practice questions",
  "estimatedMinutes": 30,
  "questionCount": 10,
  "priority": 10
}
```

#### Update Task
```
POST /api/v1/task/edit
```

**Request:**
```json
{
  "id": 1,
  "status": "completed"
}
```

#### Get Task by ID
```
POST /api/v1/task/byId
```

**Request:**
```json
{
  "id": 1
}
```

#### Get Task Stream
```
POST /api/v1/task/stream
```

**Request:**
```json
{
  "goalId": 1
}
```

**Response:**
```json
{
  "code": 0,
  "message": "Data retrieved successfully!",
  "data": {
    "todayTasks": [
      {
        "id": 1,
        "type": "learn",
        "status": "pending",
        "targetDate": "2026-02-17T00:00:00Z",
        "title": "Learn: Chapter 1",
        "estimatedMinutes": 30
      }
    ],
    "upcomingTasks": [
      {
        "id": 2,
        "type": "drill",
        "status": "pending",
        "targetDate": "2026-02-18T00:00:00Z",
        "title": "Practice: Chapter 1",
        "estimatedMinutes": 30
      }
    ],
    "overdueTasks": []
  }
}
```

#### Complete Task
```
POST /api/v1/task/complete
```

**Request:**
```json
{
  "id": 1
}
```

#### Generate Initial Plan
```
POST /api/v1/task/generate-plan
```

**Request:**
```json
{
  "goalId": 1
}
```

#### List Tasks
```
POST /api/v1/task/list
```

**Request:**
```json
{
  "pageSize": 20,
  "pageIndex": 1,
  "goalId": 1,
  "status": "pending"
}
```

#### Delete Task
```
POST /api/v1/task/delete
```

**Request:**
```json
{
  "id": 1
}
```

### Attempt Logging

#### Create Attempt
```
POST /api/v1/attempt/create
```

**Request:**
```json
{
  "goalId": 1,
  "taskId": 1,
  "questionId": 1,
  "chapterId": 1,
  "questionContentId": 1,
  "isCorrect": true,
  "studentAnswer": "B",
  "correctAnswer": "B",
  "timeSpentSeconds": 45,
  "score": 1,
  "maxScore": 1
}
```

**Note:** Creating an attempt automatically updates the knowledge state for the chapter.

#### Get Recent Attempts
```
POST /api/v1/attempt/recent
```

**Request:**
```json
{
  "goalId": 1,
  "limit": 20
}
```

#### Get Attempt Stats
```
POST /api/v1/attempt/stats
```

**Request:**
```json
{
  "goalId": 1,
  "chapterId": 1
}
```

**Response:**
```json
{
  "code": 0,
  "message": "Data retrieved successfully!",
  "data": {
    "totalAttempts": 20,
    "correctAttempts": 15,
    "wrongAttempts": 5,
    "accuracyRate": 75.0,
    "totalTimeSpent": 900,
    "averageTimePerQ": 45,
    "totalScore": 15,
    "totalMaxScore": 20
  }
}
```

#### List Attempts
```
POST /api/v1/attempt/list
```

**Request:**
```json
{
  "pageSize": 20,
  "pageIndex": 1,
  "goalId": 1,
  "chapterId": 1
}
```

## Task Types

- **learn**: Learning new knowledge points
- **drill**: Practice with questions
- **review**: Review previously learned content (SRS-driven)
- **test**: Milestone test for a chapter/module
- **mock**: Mock exam or past paper practice

## Task Statuses

- **pending**: Not started yet
- **in_progress**: Started but not completed
- **completed**: Completed successfully
- **skipped**: User skipped this task
- **failed**: User failed the task

## Spaced Repetition System (SRS)

The SRS algorithm calculates the next review date based on:

1. **Mastery Level** (0.0 to 1.0): Correct attempts / Total attempts
2. **Stability Score** (0.0 to 1.0): Memory retention strength
3. **Consecutive Performance**: Bonus for correct streaks, penalty for wrong streaks

**Interval Formula:**
```
baseInterval = 1 day
masteryFactor = 1 + (masteryLevel * 5)  // 1x to 6x
stabilityFactor = 1 + (stabilityScore * 3)  // 1x to 4x
consecutiveBonus = min(2.0, 1 + (consecutiveCorrect * 0.2))
intervalDays = baseInterval * masteryFactor * stabilityFactor * consecutiveBonus
intervalDays = clamp(intervalDays, 1, 30)  // Min 1 day, max 30 days
```

## Progress Metrics

### Coverage
Percentage of chapters that have been touched/covered:
```
Coverage = (CoveredChapters / TotalChapters) * 100
```

### Mastery
Average mastery level across all chapters:
```
Mastery = (Sum of MasteryLevels / TotalChapters) * 100
```

### Stability
Average stability score across all chapters:
```
Stability = (Sum of StabilityScores / TotalChapters) * 100
```

## Usage Flow

### 1. Create a Goal
```bash
POST /api/v1/goal/create
{
  "syllabusId": 1,
  "examDate": "2026-06-15T00:00:00Z",
  "weeklyHours": 10,
  "mode": "self"
}
```

### 2. Generate Initial Plan
```bash
POST /api/v1/task/generate-plan
{
  "goalId": 1
}
```

### 3. Get Today's Tasks
```bash
POST /api/v1/task/stream
{
  "goalId": 1
}
```

### 4. Complete Tasks and Log Attempts
```bash
POST /api/v1/attempt/create
{
  "goalId": 1,
  "taskId": 1,
  "questionId": 1,
  "chapterId": 1,
  "questionContentId": 1,
  "isCorrect": true,
  "studentAnswer": "B",
  "correctAnswer": "B",
  "timeSpentSeconds": 45,
  "score": 1,
  "maxScore": 1
}
```

### 5. Track Progress
```bash
POST /api/v1/knowledge-state/progress
{
  "goalId": 1
}
```

## Error Handling

All endpoints return a standard response format:

**Success:**
```json
{
  "code": 0,
  "message": "Success message",
  "data": { ... }
}
```

**Error:**
```json
{
  "code": 1,
  "message": "Error message",
  "data": null
}
```

## Implementation Notes

1. All API endpoints require JWT authentication
2. User ID is extracted from JWT claims using `auth.GetCurrentUser(c)`
3. All data is isolated per user - users can only access their own goals/tasks/attempts
4. Knowledge states are automatically initialized when a goal is created
5. Attempts automatically update knowledge states and trigger SRS scheduling
6. Task plan generation creates tasks for the next 7 days
7. Database migrations are handled automatically via GORM AutoMigrate

## Future Enhancements

- Adaptive task difficulty based on performance
- Machine learning-based SRS optimization
- Collaborative learning features
- Multi-device synchronization
- Offline mode support
- Advanced analytics and insights
- Personalized learning paths
- Integration with external learning resources
