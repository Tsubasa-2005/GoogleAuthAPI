BIN_DIR:=$(shell pwd)/bin

.PHONY: run
run:
	go run main.go serve

.PHONY: ogen
ogen:
	ogen -target ui/api -clean docs/openapi.yaml

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: docker
docker:
	docker compose up -d
