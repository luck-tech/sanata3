FROM golang:1.23 AS app
WORKDIR /app

RUN go install github.com/air-verse/air@latest

CMD ["air","-c",".docker/app/.air.toml"]