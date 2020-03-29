gen:
	protoc --go_out=plugins=grpc:internal/adapters/grpc api/*.proto
	
build: gen
	go build -o clncnd main.go
