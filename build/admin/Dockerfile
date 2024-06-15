FROM golang:1.22.3-alpine as build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/admin

RUN go build -o admin .

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/cmd/admin/config.toml .
COPY --from=build-stage /app/cmd/admin/admin .

EXPOSE 6988

CMD ["./admin"]
