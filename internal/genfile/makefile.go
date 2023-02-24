package genfile

import "fmt"

func Make(s string) string {
	return fmt.Sprintf(`BINARY_NAME=%s

build:
	go build -o ${BINARY_NAME} cmd/%s/main.go

run:
	go run cmd/%s/main.go

tidy:
	go mod tidy

`, s, s, s)
}
