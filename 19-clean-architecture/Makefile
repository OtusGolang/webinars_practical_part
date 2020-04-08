gen:
	protoc --go_out=plugins=grpc:internal/adapters/grpc api/*.proto

lint:
	golangci-lint run ./...

build: gen
	go build -o clncnd main.go
