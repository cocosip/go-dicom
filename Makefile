.PHONY: all build test test-coverage lint fmt vet clean install-tools help

# 默认目标
all: fmt vet lint test build

# 构建所有包
build:
	@echo "Building..."
	@go build ./...

# 构建命令行工具
build-tools:
	@echo "Building command-line tools..."
	@go build -o bin/dicominfo ./cmd/dicominfo

# 运行所有测试
test:
	@echo "Running tests..."
	@go test -race -v ./...

# 运行测试并生成覆盖率报告
test-coverage:
	@echo "Running tests with coverage..."
	@go test -race -coverprofile=coverage.out -covermode=atomic ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# 运行单个包的测试
test-pkg:
	@echo "Usage: make test-pkg PKG=./pkg/dicom"
	@go test -race -v $(PKG)

# 运行基准测试
bench:
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

# 运行 linter
lint:
	@echo "Running linter..."
	@golangci-lint run

# 格式化代码
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# 运行 go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

# 整理依赖
tidy:
	@echo "Tidying dependencies..."
	@go mod tidy

# 下载依赖
deps:
	@echo "Downloading dependencies..."
	@go mod download

# 清理构建产物
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@go clean -cache -testcache

# 安装开发工具
install-tools:
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# 生成代码（如果需要）
generate:
	@echo "Generating code..."
	@go generate ./...

# 显示帮助信息
help:
	@echo "Available targets:"
	@echo "  all            - Format, vet, lint, test, and build (default)"
	@echo "  build          - Build all packages"
	@echo "  build-tools    - Build command-line tools"
	@echo "  test           - Run all tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  test-pkg       - Run tests for a specific package (usage: make test-pkg PKG=./pkg/dicom)"
	@echo "  bench          - Run benchmarks"
	@echo "  lint           - Run golangci-lint"
	@echo "  fmt            - Format code with gofmt"
	@echo "  vet            - Run go vet"
	@echo "  tidy           - Tidy dependencies"
	@echo "  deps           - Download dependencies"
	@echo "  clean          - Clean build artifacts"
	@echo "  install-tools  - Install development tools"
	@echo "  generate       - Generate code"
	@echo "  help           - Show this help message"
