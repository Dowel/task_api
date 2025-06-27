package service

import (
	"awesomeProject2/internal/repo"
	"awesomeProject2/internal/tasks"
	"context"
	"math/rand"
	"time"
)

type TaskService struct {
	repo *repo.TaskRepository
}

func NewTaskService(repo *repo.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context) *tasks.Task {
	task := &tasks.Task{
		ID:        generateID(),
		Status:    tasks.StatusPending,
		CreatedAt: time.Now(),
	}

	s.repo.Create(task)

	go s.processTask(task.ID)

	return task
}

func (s *TaskService) processTask(id string) {
	task, exists := s.repo.Get(id)
	if !exists {
		return
	}

	now := time.Now()
	task.Status = tasks.StatusProcessing
	task.StartedAt = &now

	// Имитация долгой работы
	time.Sleep(3 * time.Minute)

	now = time.Now()
	task.CompletedAt = &now
	task.Status = tasks.StatusCompleted
	result := "task completed successfully"
	task.Result = &result
}

func (s *TaskService) GetTask(id string) (*tasks.Task, bool) {
	return s.repo.Get(id)
}

func (s *TaskService) DeleteTask(id string) {
	s.repo.Delete(id)
}

func generateID() string {
	return time.Now().Format("20060102150405") + "-" + randString(6)
}

func randString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
