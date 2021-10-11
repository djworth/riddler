
tidy:
	go mod tidy

build: tidy
	go build -o build/api-server cmd/api-server/main.go

run: tidy
	go run cmd/api-server/main.go