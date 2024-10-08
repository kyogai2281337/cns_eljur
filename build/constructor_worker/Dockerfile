FROM golang:1.22.3-alpine as build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/constructor_worker
RUN ls -la

RUN go build -o constructor .

# Используем более лёгкий образ для запуска
FROM alpine:latest

# Определяем переменные окружения
ARG NATS_URL=""
ARG BUF="10"

WORKDIR /app

# Копируем собранный бинарный файл из этапа сборки
COPY --from=build-stage /app/cmd/constructor_worker/constructor .

# Используем sh для выполнения команды с интерполированными переменными
CMD sh -c './constructor -broker "${NATS_URL}" -buffer "${BUF}"'
