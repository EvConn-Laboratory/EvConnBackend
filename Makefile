.PHONY: run migrate test

run:
    go run cmd/api/main.go

migrate:
    go run cmd/migrate/main.go

test:
    go test ./... -v

build:
    go build -o bin/api cmd/api/main.go
    go build -o bin/migrate cmd/migrate/main.go