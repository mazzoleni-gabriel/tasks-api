FROM golang:1.18-alpine

WORKDIR /app/src/tasks-api

ENV GOPATH=/app

COPY . /app/src/tasks-api

RUN go build cmd/api/main.go

ENTRYPOINT ["./main"]

EXPOSE 8080