BINARY_NAME=app

build:
	go build -o ${BINARY_NAME} cmd/app/main.go

deps:
	go mod download
