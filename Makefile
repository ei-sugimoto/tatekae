.PHONY: apibuf
apibuf:
	cd api && buf lint && buf format -w && buf generate

.PHONY: webbuf
webbuf:
	docker compose exec web bash -c "npx buf generate"


.PHONY: buf
buf:
	make apibuf
	make webbuf

.PHONY: entgen
entgen:
	cd api && go generate ./infrastructure/ent

.PHONY: new-schema
new-schema:
	cd api/infrastructure && test -n "$(name)" || (echo "name is not set" && exit 1)
	cd api/infrastructure && go run -mod=mod entgo.io/ent/cmd/ent new $(name)

.PHONY: up
up:
	docker compose up

.PHONY: down
down:
	docker compose down

.PHONY: up-build
up-build:
	docker compose up --build

.PHONY: down-vol
down-vol:
	docker compose down -v
.PHONY: prod-up
prod-up:
	docker compose -f docker-compose.prod.yml up --build

.PHONY: prod-image-build
prod-image-build:
	cd api && docker build -f Dockerfile.Prod -t api-image:latest .
	cd web && docker build -f Dockerfile.Prod -t web-image:latest .
	make prod-up