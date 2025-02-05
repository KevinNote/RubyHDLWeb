# Build stage
FROM golang:latest AS builder

WORKDIR /app

RUN apt update && apt install -y git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o build/serv cmd/serv/main.go

# Run stage
FROM ubuntu:latest

WORKDIR /app

RUN apt update && apt install -y smlnj && apt clean

COPY --from=builder /app/build .

COPY docker-config.json .
RUN mv docker-config.json config.json

RUN mkdir -p /app/tmp/task
COPY ruby/ /ruby/

EXPOSE 8080

ENTRYPOINT ["/app/serv"]
