build:
	@go build -o bins/app

run: build
	@./app   

test:
	@go test -v ./...

test-cover:
	@go test -cover -v ./...

