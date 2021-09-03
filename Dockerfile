# Build Stage
FROM golang:1.15-alpine3.13 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run Stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY .env /app
COPY view.html /app
EXPOSE ${APP_PORT}
CMD ["/app/main"]
