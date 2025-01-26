FROM golang:1.22.5-bookworm AS builder
WORKDIR /go/src/app
COPY . .
RUN mkdir ./bin && go build -o ./bin/app ./cmd/api/main.go

FROM gcr.io/distroless/cc
COPY --from=builder /go/src/app/bin/app /app
COPY --from=builder /go/src/app/db /db
EXPOSE 8000/tcp
CMD ["/app"]
