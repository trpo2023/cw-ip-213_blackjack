BIN_NAME=main

# Launching the application from main.go
dev:
	go run .

# Starting the application build
build:
	go build -o dist/$(BIN_NAME) main.go

# Launching an application from a binary file
run:
	./dist/$(BIN_NAME)

# Building and launching the application
prod: build run

.PHONY: dev build run prod
