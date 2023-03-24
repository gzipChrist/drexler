BINARY_NAME=drexler

build:
	go build -o ${BINARY_NAME} cmd/drexler/main.go

deps:
	go mod download

tidy:
	go mod tidy