# search for test files everywehere except cmd/
test:
    go test -v $(go list ./... | grep -v /cmd/)
