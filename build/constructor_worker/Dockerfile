FROM golang:1.22.3-alpine as build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/constructor_worker
RUN ls -la

RUN go build -o constructor .

FROM alpine:latest

ARG NATS_URL=""
ARG BUF="10"

WORKDIR /app

COPY --from=build-stage /app/cmd/constructor_worker/constructor .

CMD sh -c './constructor -broker "${NATS_URL}" -buffer "${BUF}"'
