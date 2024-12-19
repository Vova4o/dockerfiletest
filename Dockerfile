FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .

FROM alpine:latest

WORKDIR /app

RUN mkdir filestorage
RUN chmod 755 filestorage

COPY --from=builder /app/main .

CMD ["./main"]