# search for test files everywehere except cmd/
test:
    go test -v $(go list ./... | grep -v /cmd/)

dbgen:
    sqlc generate

dev:
    gowatch -p ./cmd/api/main.go

# build the admin UI (requires pnpm)
ui-build:
    cd ui && pnpm install --frozen-lockfile=false && pnpm build

# run the Svelte dev server (proxies to Go API on :8000)
ui-dev:
    cd ui && pnpm dev

# build the server with the admin UI baked in
build: ui-build
    go build -tags ui -o bin/app ./cmd/api

