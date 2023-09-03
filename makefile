.PHONY: generate

generate:
	go generate ./...

generate-sql:
	sqlc generate

start:
	go run ./cmd/webserver/...