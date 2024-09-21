.PHONY protoc:
	protoc --go_out=./api/web --go-grpc_out=./api/web proto/*.proto