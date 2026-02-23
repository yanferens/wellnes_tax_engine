# СТАДІЯ 1: Збірка (використовуємо офіційний образ)
FROM golang:alpine AS builder

# Встановлюємо необхідні системні пакети для збірки
RUN apk add --no-cache git

WORKDIR /app

# Спочатку копіюємо залежності для кешування
COPY go.mod go.sum ./
RUN go mod download

# Копіюємо код і збираємо бінарник
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o worker ./cmd/worker/main.go

# СТАДІЯ 2: Запуск (мінімальний образ)
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
# Копіюємо тільки скомпільований файл
COPY --from=builder /app/worker .

# Запуск
CMD ["./worker"]