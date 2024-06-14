FROM ubuntu:latest
LABEL authors="levandr"

FROM golang:1.22.3-alpine as build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN ls -a

COPY . .

RUN ls -a

WORKDIR /app/cmd/admin

RUN ls -a

RUN go build -o auth .

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/cmd/admin/config.toml .

COPY --from=build-stage /app/cmd/admin .

RUN cd .. && ls -a

RUN ls -a

EXPOSE 6988

CMD ["./admin"]
ENTRYPOINT ["top", "-b"]