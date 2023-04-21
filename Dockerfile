# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

ENV PROJECT-KEY-DOCKER=/app/config

ENV PROJECT_WEB_PORT=8080

ENV PROJECT_ADDRESS_PSQL host=project-postgres port=5432 user=rakuten password=rakuten@123 dbname=rakuten sslmode=disable

ENV PROJECT_SCHEMA_PSQL=rakuten

COPY . .

RUN go build -o rakuten

EXPOSE 8080

CMD ./rakuten