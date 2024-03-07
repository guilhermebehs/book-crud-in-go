build:
	@go build -o bins/app

run: build
	@./app   