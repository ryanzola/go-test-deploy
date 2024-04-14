run: build
	@./bin/tester

build:
	@go build -o bin/tester main.go

install:
	@go install