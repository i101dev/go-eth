build:
	@go build -o bin/goeth

run: build
	@./bin/goeth

api:
	@cd api && go run .

test:
	@go test -v ./...


hardhat:
	@cd hardhat && npx hardhat node

.PHONY: hardhat
.PHONY: api