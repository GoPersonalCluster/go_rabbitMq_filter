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

COPY --from=build /app/app .

EXPOSE 8080

ENTRYPOINT ["./app/app"]