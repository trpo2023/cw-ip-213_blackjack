APP_MAIN_PATH=./cmd/app/main.go
BIN_NAME=main
BUILD_DIR=./dist

# Launching the application from main.go
dev:
	go run $(APP_MAIN_PATH)

# Starting the application build
build:
	go build -o $(BUILD_DIR)/$(BIN_NAME) $(APP_MAIN_PATH)

# Launching an application from a binary file
run:
	$(BUILD_DIR)/$(BIN_NAME)

# Building and launching the application
prod: build run

# Running the application test
test:
	go test -v -cover -short ./...

.PHONY: dev build run prod test
