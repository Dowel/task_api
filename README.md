# task_api
Микросервис для управления долгоживущими задачами (3-5 минут выполнения) с HTTP API. Хранит состояние в памяти.

Особенности

Простое API для управления задачами
Поддержка долгих операций (3-5 минут)
Потокобезопасная реализация
Отслеживание статуса и времени выполнения
Готовность к развертыванию в Docker
Быстрый старт

1. Запуск без Docker

# Клонировать репозиторий
git clone https://github.com/dowel/task_api.git
cd task-manager

# Запустить сервер
go run cmd/server/main.go

# Запустить сервер
go run cmd/server/main.go

2. Запуск с Docker

docker build -t task-manager .
docker run -p 8080:8080 task-manager
Или с docker-compose:

docker-compose up -d
API Endpoints

Метод	Путь	Описание
POST	/tasks	Создать новую задачу
GET	/tasks/{id}	Получить статус задачи
DELETE	/tasks/{id}	Удалить задачу
Примеры использования

Создание задачи

curl -X POST http://localhost:8080/tasks
Ответ:

json
{
  "id": "20230601123045-abc123",
  "status": "pending",
  "created_at": "2023-06-01T12:30:45Z"
}
Проверка статуса

bash
curl http://localhost:8080/tasks/20230601123045-abc123
Ответ (во время выполнения):

json
{
  "id": "20230601123045-abc123",
  "status": "processing",
  "created_at": "2023-06-01T12:30:45Z",
  "started_at": "2023-06-01T12:30:46Z"
}
Удаление задачи

curl -X DELETE http://localhost:8080/tasks/20230601123045-abc123
