# Используем официальный образ Go в качестве базового
FROM golang:1.22-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы
COPY . .

# Собираем приложение
RUN go build -o main cmd/database/main.go

# Создаем директорию для логов
RUN mkdir -p /app/logs

# Запускаем приложение
CMD ["./main"]
