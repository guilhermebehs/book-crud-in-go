build:
	@go build -o bin/app

run: build
	@./bin/app   

test:
	@go test -v ./...

test-cov:
	@go test -cover -v ./...

