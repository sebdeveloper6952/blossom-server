# search for test files everywehere except cmd/
test:
    go test -v $(go list ./... | grep -v /cmd/)

dbgen:
    sqlc generate

dev:
    go run ./cmd/api/main.go
