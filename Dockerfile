FROM golang:1.22.5-bookworm AS builder
WORKDIR /go/src/app
COPY . .
RUN go build -o /go/bin/app

FROM gcr.io/distroless/cc
COPY --from=builder /go/bin/app /
COPY --from=builder /go/src/app/index.html /
EXPOSE 8000/tcp
ENTRYPOINT ["/app"]
