# Use a multi-stage build with Go 1.23 (Alpine-based)
FROM golang:1.23-alpine AS builder

# Install required packages for module resolution and CA certs
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy dependency files first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy application code
COPY . .

# Build a static binary suitable for scratch
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o fbref_exporter

# ---- Final scratch image ----
FROM scratch

# Add CA certs (needed for HTTPS scraping)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy built exporter binary
COPY --from=builder /app/fbref_exporter /fbref_exporter

# Expose metrics port
EXPOSE 9101

# Entry point
ENTRYPOINT ["/fbref_exporter"]
