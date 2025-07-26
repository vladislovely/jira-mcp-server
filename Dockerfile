# ──────────────────────────────
# 1. Build stage
# ──────────────────────────────
FROM golang:1.24-alpine AS builder

# Собираем статически, для Linux/amd64
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src

# Кэшируем зависимости
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Копируем остальной код и собираем
COPY . .
RUN cp .env.example .env

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -v -o /bin/app ./cmd/...

# ──────────────────────────────
# 2. Runtime stage
# ──────────────────────────────
FROM debian:12-slim AS runtime

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /bin/app /app
COPY --from=builder /src/.env ./

ENTRYPOINT ["/app"]