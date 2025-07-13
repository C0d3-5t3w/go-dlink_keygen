# D-Link Key Generator Makefile

.PHONY: build run clean test

# Build the binary
build:
	go build -o dlink-keygen main.go

# Run the program with a sample BSSID
run:
	go run main.go

# Run with example BSSID
example:
	go run main.go 00:1B:2F:12:34:56

# Clean build artifacts
clean:
	rm -f dlink-keygen

# Test the build
test:
	go build -o /dev/null main.go

# Install dependencies (if any)
deps:
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Show help
help:
	@echo "Available targets:"
	@echo "  build   - Build the binary"
	@echo "  run     - Run the program"
	@echo "  example - Run with example BSSID"
	@echo "  clean   - Clean build artifacts"
	@echo "  test    - Test build"
	@echo "  deps    - Install dependencies"
	@echo "  fmt     - Format code"