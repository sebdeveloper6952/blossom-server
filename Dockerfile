# syntax=docker/dockerfile:1.7

# --- Admin UI build ---
FROM node:22-alpine AS ui-builder
WORKDIR /src/ui
RUN corepack enable
COPY ui/package.json ui/pnpm-lock.yaml* ./
RUN pnpm install --frozen-lockfile
COPY ui/ ./
RUN pnpm build

# --- Go build, headless (no UI embedded) ---
FROM dhi.io/golang:1.26-dev AS builder-headless
WORKDIR /go/src/app
COPY . .
RUN mkdir -p ./bin && CGO_ENABLED=1 go build \
    -ldflags "-linkmode external -extldflags '-static'" \
    -o ./bin/app ./cmd/api/main.go

# --- Go build, with UI embedded ---
FROM dhi.io/golang:1.26-dev AS builder-ui
WORKDIR /go/src/app
COPY . .
COPY --from=ui-builder /src/ui/build ./ui/build
RUN mkdir -p ./bin && CGO_ENABLED=1 go build \
    -tags ui \
    -ldflags "-linkmode external -extldflags '-static'" \
    -o ./bin/app ./cmd/api/main.go

# --- Runtime, headless ---
FROM dhi.io/static:20250419 AS runtime-headless
COPY --from=builder-headless /go/src/app/bin/app /app
COPY --from=builder-headless /go/src/app/db /db
EXPOSE 8000/tcp
CMD ["/app"]

# --- Runtime, with UI (default target) ---
FROM dhi.io/static:20250419 AS runtime-ui
COPY --from=builder-ui /go/src/app/bin/app /app
COPY --from=builder-ui /go/src/app/db /db
EXPOSE 8000/tcp
CMD ["/app"]
