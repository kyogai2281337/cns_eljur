FROM golang:1.22.3-alpine as build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/constructor

RUN go build -o constructor .

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/cmd/constructor/config.toml .
COPY --from=build-stage /app/cmd/constructor/constructor .

EXPOSE 6989

CMD ["./constructor"]
