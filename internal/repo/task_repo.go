package repo

import (
	"awesomeProject2/internal/tasks"
	"sync"
)

type TaskRepository struct {
	mu    sync.RWMutex
	tasks map[string]*tasks.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]*tasks.Task),
	}
}

func (r *TaskRepository) Create(task *tasks.Task) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID] = task
}

func (r *TaskRepository) Get(id string) (*tasks.Task, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	task, exists := r.tasks[id]
	return task, exists
}

func (r *TaskRepository) Delete(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.tasks, id)
}
