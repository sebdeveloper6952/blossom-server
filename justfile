# search for test files everywehere except cmd/
test:
    go test -v $(go list ./... | grep -v /cmd/)

dbgen:
    sqlc generate

dev:
    gowatch -p ./cmd/api/main.go
