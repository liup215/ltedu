package model

import "time"

// Task type constants
const (
	TaskTypeLearn  = "learn"  // Learning new knowledge points
	TaskTypeDrill  = "drill"  // Practice with questions
	TaskTypeReview = "review" // Review previously learned content (SRS)
	TaskTypeTest   = "test"   // Milestone test for a chapter/module
	TaskTypeMock   = "mock"   // Mock exam or past paper practice
)

// Task status constants
const (
	TaskStatusPending   = "pending"   // Not started yet
	TaskStatusInProgress = "in_progress" // Started but not completed
	TaskStatusCompleted = "completed" // Completed successfully
	TaskStatusSkipped   = "skipped"   // User skipped this task
	TaskStatusFailed    = "failed"    // User failed the task (e.g., test below threshold)
)

// Task represents a daily task card in the learning stream
type Task struct {
	Model
	UserId         uint       `json:"userId" gorm:"index:idx_user_goal_date"`
	User           User       `json:"user" gorm:"foreignKey:UserId"`
	GoalId         uint       `json:"goalId" gorm:"index:idx_user_goal_date"`
	Goal           Goal       `json:"goal" gorm:"foreignKey:GoalId"`
	Type           string     `json:"type" gorm:"index"`              // learn, drill, review, test, mock
	Status         string     `json:"status" gorm:"default:pending"` // pending, in_progress, completed, skipped, failed
	TargetDate     time.Time  `json:"targetDate" gorm:"index:idx_user_goal_date"` // Date this task is assigned to
	ChapterId      *uint      `json:"chapterId,omitempty" gorm:"index"`           // For learn/drill/review tasks
	Chapter        *Chapter   `json:"chapter,omitempty" gorm:"foreignKey:ChapterId"`
	PaperId        *uint      `json:"paperId,omitempty"`                          // For mock tasks
	Paper          *Paper     `json:"paper,omitempty" gorm:"foreignKey:PaperId"`
	Title          string     `json:"title"`                                      // Human-readable task title
	Description    string     `json:"description"`                                // Task description
	EstimatedMinutes int      `json:"estimatedMinutes"`                           // Estimated time to complete
	QuestionCount  int        `json:"questionCount"`                              // Number of questions for drill/test
	Priority       int        `json:"priority" gorm:"default:0"`                  // Higher priority = more important
	IsLocked       bool       `json:"isLocked" gorm:"default:false"`              // User-locked task (can't be rescheduled)
	CompletedAt    *time.Time `json:"completedAt,omitempty"`                      // When the task was completed
	PlanVersion    int        `json:"planVersion" gorm:"default:1"`               // Plan version for tracking changes
}

// TaskQuery for filtering tasks
type TaskQuery struct {
	ID         uint      `json:"id"`
	UserId     uint      `json:"userId"`
	GoalId     uint      `json:"goalId"`
	Type       string    `json:"type"`
	Status     string    `json:"status"`
	TargetDate time.Time `json:"targetDate"`
	ChapterId  uint      `json:"chapterId"`
	DateFrom   time.Time `json:"dateFrom"`
	DateTo     time.Time `json:"dateTo"`
	Page
}

// TaskCreateRequest for creating a new task
type TaskCreateRequest struct {
	GoalId           uint      `json:"goalId" binding:"required"`
	Type             string    `json:"type" binding:"required"`
	TargetDate       time.Time `json:"targetDate" binding:"required"`
	ChapterId        *uint     `json:"chapterId,omitempty"`
	PaperId          *uint     `json:"paperId,omitempty"`
	Title            string    `json:"title" binding:"required"`
	Description      string    `json:"description"`
	EstimatedMinutes int       `json:"estimatedMinutes"`
	QuestionCount    int       `json:"questionCount"`
	Priority         int       `json:"priority"`
}

// TaskUpdateRequest for updating a task
type TaskUpdateRequest struct {
	ID               uint       `json:"id" binding:"required"`
	Status           *string    `json:"status,omitempty"`
	TargetDate       *time.Time `json:"targetDate,omitempty"`
	Title            *string    `json:"title,omitempty"`
	Description      *string    `json:"description,omitempty"`
	EstimatedMinutes *int       `json:"estimatedMinutes,omitempty"`
	QuestionCount    *int       `json:"questionCount,omitempty"`
	Priority         *int       `json:"priority,omitempty"`
	IsLocked         *bool      `json:"isLocked,omitempty"`
	CompletedAt      *time.Time `json:"completedAt,omitempty"`
}

// TaskStreamResponse for showing tasks in a date range
type TaskStreamResponse struct {
	TodayTasks   []Task `json:"todayTasks"`   // Tasks for today
	UpcomingTasks []Task `json:"upcomingTasks"` // Tasks for the next 7 days
	OverdueTasks []Task `json:"overdueTasks"`  // Overdue tasks
}

// TaskCompleteRequest for marking a task as complete
type TaskCompleteRequest struct {
	ID uint `json:"id" binding:"required"`
}
