FROM dhi.io/golang:1.26-dev AS builder
WORKDIR /go/src/app
COPY . .
RUN mkdir ./bin && CGO_ENABLED=1 go build \
    -ldflags "-linkmode external -extldflags '-static'" \
    -o ./bin/app ./cmd/api/main.go

FROM dhi.io/static:20250419
COPY --from=builder /go/src/app/bin/app /app
COPY --from=builder /go/src/app/db /db
EXPOSE 8000/tcp
CMD ["/app"]
