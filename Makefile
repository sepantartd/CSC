# متغیرها
BINARY_NAME=cloud-saver
GO_FILES=main.go

.PHONY: all build run test docker-build clean lint help

all: test build

build: ## کامپایل پروژه و ساخت باینری
@echo "🔨 در حال کامپایل پروژه..."
go build -o $(BINARY_NAME) $(GO_FILES)

run: ## اجرای پروژه در حالت تحلیل
@go run $(GO_FILES) analyze

test: ## اجرای تمام تست‌های واحد
@echo "🧪 در حال اجرای تست‌ها..."
go test -v ./...

docker-build: ## ساخت ایمیج داکر
@echo "🐳 در حال ساخت ایمیج داکر..."
docker build -t cloud-saver:latest .

clean: ## پاکسازی فایل‌های کامپایل شده
@echo "🧹 در حال پاکسازی..."
rm -f $(BINARY_NAME)
rm -rf dist/

help: ## نمایش دستورات قابل استفاده در Makefile
@echo "دستورات قابل استفاده:"
@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'
