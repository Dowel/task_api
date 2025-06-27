package main

import (
	"awesomeProject2/internal/handlers"
	"awesomeProject2/internal/repo"
	"awesomeProject2/internal/service"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	// Инициализация зависимостей
	taskRepo := repo.NewTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Роутер
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Task Manager API"))
	})

	taskHandler.RegisterRoutes(r)

	// Запуск сервера
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
