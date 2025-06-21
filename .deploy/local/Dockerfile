# Билд этап для компиляции приложения
FROM golang:1.21-alpine AS builder

# Устанавливаем зависимости для сборки
RUN apk add --no-cache \
    git \
    make \
    gcc \
    musl-dev

# Настраиваем рабочую директорию
WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Компилируем приложение с флагами для уменьшения размера
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/bin/main ./cmd/app

# Финальный этап с минимальным образом
FROM alpine:3.18

# Создаем пользователя для безопасности
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Настраиваем рабочую директорию
WORKDIR /app

# Копируем бинарник из builder-этапа
COPY --from=builder --chown=appuser:appgroup /app/bin/main ./app
# Копируем миграции (если есть)
COPY --chown=appuser:appgroup migrations ./migrations

# Переключаемся на непривилегированного пользователя
USER appuser

# Экспортируем порт (если нужно)
EXPOSE 8080

# Команда запуска
CMD ["./app"]