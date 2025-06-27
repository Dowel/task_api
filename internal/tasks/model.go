package tasks

import "time"

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusProcessing TaskStatus = "processing"
	StatusCompleted  TaskStatus = "completed"
	StatusFailed     TaskStatus = "failed"
)

type Task struct {
	ID          string
	Status      TaskStatus
	CreatedAt   time.Time
	StartedAt   *time.Time
	CompletedAt *time.Time
	Result      *string
	Error       *string
}
