package handlers

import (
	"awesomeProject2/internal/service"
	"awesomeProject2/internal/tasks"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) RegisterRoutes(r *chi.Mux) {
	r.Post("/tasks", h.createTask)
	r.Get("/tasks/{id}", h.getTask)
	r.Delete("/tasks/{id}", h.deleteTask)
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	task := h.service.CreateTask(r.Context())

	writeJSON(w, http.StatusCreated, toTaskResponse(task))
}

func (h *TaskHandler) getTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, exists := h.service.GetTask(id)
	if !exists {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}

	writeJSON(w, http.StatusOK, toTaskResponse(task))
}

func (h *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	h.service.DeleteTask(id)

	w.WriteHeader(http.StatusNoContent)
}

type TaskResponse struct {
	ID          string     `json:"id"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	DurationSec *float64   `json:"duration_sec,omitempty"`
	Result      *string    `json:"result,omitempty"`
	Error       *string    `json:"error,omitempty"`
}

func toTaskResponse(t *tasks.Task) TaskResponse {
	resp := TaskResponse{
		ID:          t.ID,
		Status:      string(t.Status),
		CreatedAt:   t.CreatedAt,
		StartedAt:   t.StartedAt,
		CompletedAt: t.CompletedAt,
		Result:      t.Result,
		Error:       t.Error,
	}

	if t.StartedAt != nil && t.CompletedAt != nil {
		dur := t.CompletedAt.Sub(*t.StartedAt).Seconds()
		resp.DurationSec = &dur
	}

	return resp
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
