# Используем официальный образ Go для сборки
FROM golang:1.23.2-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /build

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем все файлы проекта в контейнер
COPY . .

# Собираем приложение
RUN go build -o obstore_api ./cmd/apiserver

# Используем минимальный образ для запуска
FROM gcr.io/distroless/base-debian12

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем собранный бинарник из предыдущего этапа
COPY --from=builder /build/obstore_api .

# Открываем порт
EXPOSE 8090

# Запускаем приложение
CMD ["/app/obstore_api"]