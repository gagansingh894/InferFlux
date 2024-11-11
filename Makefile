.PHONY: all lint generate-mocks generate-protos test

all: lint test

lint:
	@echo "Running linter..."
	#todo: fix thi
	#golangci-lint run

test:
	@echo "Running tests..."
	go test ./...

test_coverage:
	@echo "Running tests with coverage..."
	go test ./... -coverprofile=coverage.out