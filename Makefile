.PHONY: protoc
protoc:
	protoc --go_out=./api/web --go-grpc_out=./api/web proto/*.proto

.PHONY: entgen
entgen:
	cd api && go generate ./infrastructure/ent

.PHONY: new-schema
new-schema:
	cd api/infrastructure && test -n "$(name)" || (echo "name is not set" && exit 1)
	cd api/infrastructure && go run -mod=mod entgo.io/ent/cmd/ent new $(name)