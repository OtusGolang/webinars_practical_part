generate:
	rm -rf elections/pb
	mkdir -p elections/pb
	protoc --proto_path=api/elections/ --go_out=elections/pb --go-grpc_out=elections/pb api/elections/*.proto

	protoc \
		--proto_path=api/elections-with-admin/ \
		--go_out=elections-with-admin \
		--go-grpc_out=elections-with-admin \
		api/elections-with-admin/*.proto
	protoc \
		--proto_path=api/elections-with-stats/ \
		--go_out=elections-with-stats \
		--go-grpc_out=elections-with-stats \
		api/elections-with-stats/*.proto
