# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

ENV PROJECT-KEY-DOCKER=/app/config

ENV PROJECT_WEB_PORT=8080

ENV PROJECT_ADDRESS_PSQL host=project-postgres port=5432 user=codetest password=codetest@123 dbname=codetest sslmode=disable

ENV PROJECT_SCHEMA_PSQL=codetest

COPY . .

RUN go build -o codetest

EXPOSE 8080

CMD ./codetest docker