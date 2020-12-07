generate:
	mkdir -p elections/pb
	protoc --proto_path=api/elections/ --go_out=elections/pb --go-grpc_out=elections/pb api/elections/*.proto

	protoc --proto_path=api/elections_with_admin/ --go_out=elections_with_admin --go-grpc_out=elections_with_admin api/elections_with_admin/*.proto
	protoc --proto_path=api/elections_with_stats/ --go_out=elections_with_stats --go-grpc_out=elections_with_stats api/elections_with_stats/*.proto
