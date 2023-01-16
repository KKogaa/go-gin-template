tidy:
	go mod tidy
build:
	go build -o bin/cmd cmd/main.go
run:
	go run cmd/main.go