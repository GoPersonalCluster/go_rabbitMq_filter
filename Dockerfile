FROM golang:1.25 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -o /app/golang-app ./app/internal

FROM debian:12-slim

WORKDIR /app

ENV rabbitmq_username=admin \
    rabbitmq_password=admin \
    rabbitmq_host=rabbitmq \
    rabbitmq_port=5672 \
    postgres_host=postgres \
    postgres_port=5432 \
    postgres_db=pg_rabbitmq \
    postgres_user=postgres \
    postgres_password=postgres \
    host_name=filterQueue

COPY --from=build /app/golang-app .

EXPOSE 8080

ENTRYPOINT ["./golang-app"]