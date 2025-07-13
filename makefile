.PHONY: build run clean test

build:
	go build -o dlink-keygen main.go

run:
	go run main.go

example:
	go run main.go 00:1B:2F:12:34:56

clean:
	rm -f dlink-keygen

test:
	go build -o /dev/null main.go

deps:
	go mod tidy

fmt:
	go fmt ./...

help:
	@echo "Available targets:"
	@echo "  build   - Build the binary"
	@echo "  run     - Run the program"
	@echo "  example - Run with example BSSID"
	@echo "  clean   - Clean build artifacts"
	@echo "  test    - Test build"
	@echo "  deps    - Install dependencies"
	@echo "  fmt     - Format code"