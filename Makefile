build:
	@go fmt ./...
	@go build -o bin/mercury

run: build
	@./bin/mercury

test:
	@go test -v ./...
