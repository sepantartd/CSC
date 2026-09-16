# مرحله کامپایل (Build stage)
FROM golang:1.22-alpine AS builder

WORKDIR /app

# کپی فایل‌های مدیریت وابستگی
COPY go.mod go.sum ./
RUN go mod download

# کپی سورس‌کد و کامپایل پروژه
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o cloud-saver main.go

# مرحله نهایی (Final lightweight image)
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/cloud-saver .

ENTRYPOINT ["./cloud-saver"]
CMD ["analyze"]
