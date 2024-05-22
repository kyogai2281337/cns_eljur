FROM golang:1.22.3-alpine as build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN ls -a

COPY . .

RUN ls -a

WORKDIR /app/cmd/auth
RUN ls -a
RUN go build -o auth .

FROM alpine:latest
WORKDIR /app
COPY --from=build-stage /app/cmd/auth/auth.toml .
COPY --from=build-stage /app/cmd/auth/auth .

RUN cd .. && ls -a

RUN ls -a

EXPOSE 6987
CMD ["./auth"]
