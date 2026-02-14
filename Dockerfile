# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Stage 2: Run the binary in a minimal container
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]