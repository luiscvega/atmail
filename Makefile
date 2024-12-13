default: test

atmail:
	PORT=8080 go run ./cmd/atmail

test:
	go test ./...

.PHONY: api
api:
	go generate ./api

db:
	mysql atmail < setup.sql
