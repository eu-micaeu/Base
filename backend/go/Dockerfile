## Build stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download && \
    go build -o /base-api ./cmd

## Runtime stage
FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=builder /base-api /base-api
ENV ADDR=:8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/base-api"]
