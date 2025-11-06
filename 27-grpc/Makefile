generate:
	rm -rf elections/pb
	mkdir -p elections/pb

	protoc --proto_path=api/elections/ --go_out=elections/pb --go-grpc_out=elections/pb api/elections/*.proto

	protoc --proto_path=api/elections-with-admin/ --go_out=elections-with-admin/pb --go-grpc_out=elections-with-admin/pb api/elections-with-admin/*.proto
	protoc --proto_path=api/elections-with-stats/ --go_out=elections-with-stats/pb --go-grpc_out=elections-with-stats/pb api/elections-with-stats/*.proto

evans:
	evans --proto api/elections/elections.proto repl

evans-stats:
	evans --proto api/elections-with-stats/elections.proto repl

evans-admin:
	evans --proto api/elections-with-admin/elections.proto repl
