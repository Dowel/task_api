FROM golang:1.24.2
LABEL authors="demidovartem"
WORKDIR /app

# Копируем файлы модулей для кэширования
COPY go.mod .
COPY go.sum .
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник с отключенным CGO и оптимизациями
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /task-manager ./cmd/server/

WORKDIR /app

# Копируем бинарник и шаблоны
COPY --from=builder /task-manager /app/task-manager

# Необходимые пакеты для alpine (для работы с TLS и т.д.)
RUN apk --no-cache add ca-certificates

# Открываем порт сервера
EXPOSE 8080

# Запускаем сервер
CMD ["/app/task-manager"]