build:
	@go build -o bin/goeth

run: build
	@./bin/goeth

test:
	@go test -v ./...