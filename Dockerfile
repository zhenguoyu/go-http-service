# 构建阶段
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/httpserver

# 运行阶段  
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY config/config.yaml ./config/

EXPOSE 8080
CMD ["./server"]
