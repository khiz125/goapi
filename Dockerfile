# stage 1: Build
FROM golang:1.25 AS builder
WORKDIR /app

RUN --mount=type=bind,source=.,target=/src,readonly \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    cd /src && CGO_ENABLED=0 GOOS=linux go build -o /app/goapi .

# stage 2: Run
FROM alpine:3.22.1
WORKDIR /app
COPY --from=builder /app/goapi .
CMD ["./goapi"]

