# Stage 1: Build
FROM golang:1.24.6-alpine3.21 AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Install migrate CLI in builder
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Build the app binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./server_go ./cmd/main.go

# Stage 2: Run
FROM alpine:3.21

WORKDIR /app

# Install CA certificates for TLS and mysql client for the entrypoint healthcheck
RUN apk add --no-cache ca-certificates mysql-client

# Copy the app binary from builder
COPY --from=builder /app/server_go ./

# Copy the migrate binary from builder
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Copy runtime files — templates, static assets, migrations, entrypoint
COPY --from=builder /app/pkg/static ./pkg/static
COPY --from=builder /app/database ./database
COPY --from=builder /app/entrypoint.sh ./entrypoint.sh

RUN chmod +x ./entrypoint.sh

EXPOSE 8000

ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["./server_go"]
