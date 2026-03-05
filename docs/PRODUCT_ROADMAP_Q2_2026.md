# LTEdu Product Roadmap — Q2 2026

## Overview

This document translates the findings from [MARKET_RESEARCH_2026.md](./MARKET_RESEARCH_2026.md) into an actionable, time-boxed roadmap for Q2 2026 (April–June). Each feature is described with priority, effort, dependencies, technical feasibility, and a phased delivery plan.

---

## Roadmap Summary

| # | Feature | Priority | Effort | Start | Target |
|---|---|---|---|---|---|
| F1 | AI Step-by-Step Explanations | 🔴 P1 | 3 weeks | Week 1 | Week 3 |
| F2 | Student Progress Dashboard | 🔴 P1 | 3.5 weeks | Week 2 | Week 5 |
| F3 | Spaced-Repetition Review Queue | 🟡 P2 | 2.5 weeks | Week 4 | Week 6 |
| F4 | Study-Reminder Notifications | 🟡 P2 | 1.5 weeks | Week 6 | Week 7 |
| F5 | Shareable Achievement Badges | 🟢 P3 | 2 weeks | Week 8 | Week 10 |

---

## F1 — AI Step-by-Step Explanations

### Problem
Students abandon practice sessions when they cannot understand answer explanations. Current explanations are static and brief.

### Solution
On-demand AI-generated step-by-step explanations, cached per question in the database to minimise API cost.

### Technical Design

**Backend**
- New GORM model `model/question_explanation.go`:
  ```go
  type QuestionExplanation struct {
      ID         uint   `gorm:"primaryKey"`
      QuestionID uint   `gorm:"uniqueIndex;not null"`
      Content    string `gorm:"type:text"`
      CreatedAt  time.Time
      UpdatedAt  time.Time
  }
  ```
- Service method `service/question.go` → `GetOrGenerateExplanation(questionID uint) (string, error)`:
  1. Check cache in `lt_question_explanation`.
  2. On miss, build prompt from question stem + answer, call AI service.
  3. Store result, return content.
- New endpoint `POST /api/v1/question/explain` (auth required):
  ```json
  // Request
  { "questionId": 42 }
  // Response
  { "code": 0, "data": { "explanation": "Step 1: ..." } }
  ```

**Frontend**
- "Show Explanation" button in the attempt result overlay.
- Lazy-load explanation text on button click.
- Display as collapsible section below the static answer.

### Feasibility Assessment
- **Risk: Low.** `service/ai.go` already wraps an LLM client. The caching layer prevents runaway API costs. AutoMigrate handles table creation.
- **Constraint:** Requires a valid LLM API key in `conf/config.yaml`.

### Acceptance Criteria
- [ ] Explanation generated and cached on first request.
- [ ] Subsequent requests return cached value without calling AI.
- [ ] UI button visible after submitting a practice question.
- [ ] Response time < 3 s for cached; < 10 s for uncached.

---

## F2 — Student Progress Dashboard

### Problem
Students lack visibility into their own mastery progression, making it hard to self-direct study.

### Solution
A dedicated student dashboard page showing per-exam-node mastery, accuracy trends by question type, and a study-streak calendar.

### Technical Design

**Backend**
- Extend `service/statistic.go` with:
  - `GetStudentMasteryMap(userID uint) ([]ExamNodeMastery, error)` — aggregates `lt_knowledge_state` by `exam_node_id`.
  - `GetAccuracyTrend(userID uint, days int) ([]DailyAccuracy, error)` — groups `lt_attempt` by date.
  - `GetStudyStreak(userID uint) (int, error)` — counts consecutive days with at least one attempt.
- New endpoint `GET /api/v1/student/dashboard` (auth required):
  ```json
  {
    "code": 0,
    "data": {
      "masteryMap": [{ "examNodeId": 1, "name": "...", "masteryPct": 72 }],
      "accuracyTrend": [{ "date": "2026-04-01", "correct": 18, "total": 25 }],
      "studyStreak": 7,
      "totalPracticed": 340
    }
  }
  ```

**Frontend**
- New route `/student/dashboard` → `web/src/views/student/Dashboard.vue`.
- Components: `MasteryMap.vue` (bar chart by exam node), `AccuracyChart.vue` (line chart), `StreakCalendar.vue`.
- Uses existing Axios service pattern; add `web/src/services/studentService.ts`.

### Feasibility Assessment
- **Risk: Low–Medium.** Data exists in `lt_attempt` and `lt_knowledge_state`. No schema changes required. Frontend chart library (e.g., ECharts, already common in Vue edtech projects) may need to be added.
- **Constraint:** Query performance — add composite index on `(user_id, created_at)` in `lt_attempt`.

### Acceptance Criteria
- [ ] Dashboard loads in < 2 s for a student with 500+ attempts.
- [ ] Mastery percentage matches manual calculation from `lt_knowledge_state`.
- [ ] Study streak resets correctly if a day is skipped.
- [ ] Mobile-responsive layout.

---

## F3 — Spaced-Repetition Review Queue

### Problem
Students do not revisit previously practised material at optimal intervals, leading to forgetting. Competitor platforms (Anki, Magoosh) win on retention.

### Solution
Implement SM-2 scheduling on top of `KnowledgeState`. Surface a daily "Review Today" queue of due items on the home screen.

### Technical Design

**Schema Change**
Add columns to `lt_knowledge_state` via AutoMigrate:
```go
NextReviewAt  *time.Time `gorm:"index"`
IntervalDays  int        `gorm:"default:1"`
EaseFactor    float32    `gorm:"default:2.5"`
RepCount      int        `gorm:"default:0"`
```

**Service**
- `service/knowledgeState.go` → `RecordReview(userID, questionID uint, quality int) error`:
  - `quality` 0–5 (SM-2 scale).
  - Update `IntervalDays`, `EaseFactor`, `NextReviewAt` per SM-2 formula.
- `GetReviewQueue(userID uint, limit int) ([]Question, error)`:
  - `WHERE user_id = ? AND next_review_at <= NOW() ORDER BY next_review_at ASC LIMIT ?`

**API**
- `GET /api/v1/review/queue` → returns list of due questions (reuses existing Question response shape).
- `POST /api/v1/review/submit` → accepts `{ questionId, quality }`, calls `RecordReview`.

**Frontend**
- "Review Today (N)" badge on student home / dashboard.
- Dedicated `/student/review` page with flashcard-style flip UI.

### Feasibility Assessment
- **Risk: Low.** SM-2 is a well-documented, copyright-free algorithm. No external dependency needed. Schema change is additive (nullable columns with defaults).
- **Constraint:** Existing `lt_knowledge_state` rows will have `next_review_at = NULL` until first review — treat NULL as "due now" in the queue query.

### Acceptance Criteria
- [ ] SM-2 intervals match reference implementation for quality scores 0–5.
- [ ] Review queue returns only questions due today or earlier.
- [ ] Queue is empty for a brand-new student (no prior practice).
- [ ] After reviewing, next due date is updated correctly.

---

## F4 — Study-Reminder Notifications

### Problem
Students forget to return to the platform. Email reminders with a personalised "you have N cards due" message drive re-engagement.

### Solution
A nightly scheduled task that emails students whose review queue is non-empty, with user-configurable reminder preferences.

### Technical Design

**Schema**
Extend `lt_user` (or a new `lt_user_reminder_pref` table) with:
```go
ReminderEnabled bool      `gorm:"default:false"`
ReminderHour    int       `gorm:"default:8"` // local hour 0–23
```

**Backend**
- New task `task/study_reminder.go`:
  - Runs hourly via the existing task scheduler.
  - Selects users with `reminder_enabled = true` and `reminder_hour = current_hour`.
  - For each user, calls `GetReviewQueue` (F3) — if non-empty, sends email via `service/email.go`.
- Email template: "Hi {name}, you have {N} items to review today. [Start reviewing →]"
- Preference API: `POST /api/v1/user/reminder-settings` accepts `{ enabled, hour }`.

**Frontend**
- Settings toggle + hour picker in the student profile page.

### Feasibility Assessment
- **Risk: Low.** `service/email.go` and the task scheduler already exist. This feature composes two existing capabilities.
- **Constraint:** Email deliverability depends on SMTP configuration in `conf/config.yaml`.

### Acceptance Criteria
- [ ] Reminder email sent only when review queue is non-empty.
- [ ] No duplicate emails if task runs multiple times within the same hour.
- [ ] User can disable reminders and stop receiving emails immediately.

---

## F5 — Shareable Achievement Badges

### Problem
No viral acquisition loop. Competitors with certificate-sharing grow user bases through LinkedIn and social media.

### Solution
Auto-issue SVG/PNG badges when a student completes a learning plan phase or achieves ≥ 80 % mastery on an exam node. Each badge has a public verification URL.

### Technical Design

**Schema**
```go
// model/badge.go
type Badge struct {
    ID           uint      `gorm:"primaryKey"`
    UserID       uint      `gorm:"index;not null"`
    Type         string    // "phase_complete" | "exam_node_mastery"
    RefID        uint      // phase_plan_id or exam_node_id
    Token        string    `gorm:"uniqueIndex;size:64"` // 32-byte hex
    IssuedAt     time.Time
    SVGContent   string    `gorm:"type:text"`
}
```

**Service**
- `service/badge.go`:
  - `IssueBadge(userID uint, badgeType string, refID uint) (*Badge, error)` — idempotent (check existing by user+type+ref).
  - `GenerateSVG(user User, badgeType, refName string) string` — Go `text/template` based SVG.
- Hook `IssueBadge` into:
  - `service/phase_plan.go` when a phase plan is marked complete.
  - `service/knowledgeState.go` when mastery crosses 80 %.

**API**
- `GET /api/v1/badge/:token` — public, no auth; returns badge metadata + SVG URL.
- `GET /api/v1/user/badges` — authenticated; returns user's badge list.

**Frontend**
- Badge gallery in student profile.
- "Share on WeChat / Copy link" button per badge.

### Feasibility Assessment
- **Risk: Medium.** SVG templating and idempotent issuance are straightforward. Integration hooks into phase completion require careful coordination with existing phase-plan update logic to avoid double-issuing.
- **Constraint:** SVG design requires a UX asset (template). Provide a default geometric design in code; allow custom templates via config path.

### Acceptance Criteria
- [ ] Badge issued exactly once per user+type+ref combination.
- [ ] Public verification URL resolves without authentication.
- [ ] SVG renders correctly at 300×300 px.
- [ ] "Share" button copies public URL to clipboard.

---

## Implementation Sequencing

```
Week 1-3:   F1 (AI Explanations) — backend + frontend
Week 2-5:   F2 (Progress Dashboard) — backend + frontend [parallel with F1]
Week 4-6:   F3 (Spaced Repetition) — backend + frontend
Week 6-7:   F4 (Reminders) — backend + frontend
Week 8-10:  F5 (Badges) — backend + frontend
```

---

## Risk Register

| Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|
| LLM API latency for F1 | Medium | Medium | Aggressive caching; async generation with polling |
| DB query performance for F2 | Low | Medium | Add indexes before deploying; test with 10k-row seed data |
| SM-2 implementation bug in F3 | Low | High | Unit-test against published SM-2 test vectors |
| SMTP misconfiguration for F4 | Medium | Low | Fail silently, log error; do not block user flow |
| Double-badge issuance in F5 | Low | Medium | DB unique constraint as safety net |

---

*Roadmap generated: 2026-03-05 | Author: Copilot autonomous product agent*
