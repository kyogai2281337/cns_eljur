FROM golang:latest as build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY  . .

RUN ls -a

WORKDIR /app/cmd/auth

RUN ls -a

RUN GOOS=linux && go build -o auth .

FROM alpine

COPY --from=build-stage /app/cmd/auth/auth /app/cmd/auth/auth.toml ./app/

EXPOSE 6987

CMD ["./app/auth"]
