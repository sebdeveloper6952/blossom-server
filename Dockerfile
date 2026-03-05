FROM dhi.io/golang:1.26 AS builder
WORKDIR /go/src/app
COPY . .
RUN mkdir ./bin && go build -o ./bin/app ./cmd/api/main.go

FROM dhi.io/static:20250419
COPY --from=builder /go/src/app/bin/app /app
COPY --from=builder /go/src/app/db /db
EXPOSE 8000/tcp
CMD ["/app"]
