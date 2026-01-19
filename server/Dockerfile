# ---Build stage---
FROM golang:1.25.6-alpine AS builder

# Build arguments for multi-architecture support
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

# Copy go.mod and go.sum first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code (done last to maximize cache efficiency)
COPY . .

# Builds the application as a statically linked binary, to allow it to run on alpine
# TARGETOS and TARGETARCH are automatically set by Docker BuildKit
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -installsuffix cgo -o app .

# ---Final stage---
FROM alpine:3.23.2

WORKDIR /app

# Add packages and create non-root user
RUN apk -U upgrade \
    && apk add --no-cache dumb-init ca-certificates \
    && addgroup -g 1000 appuser \
    && adduser -D -u 1000 -G appuser appuser \
    && chown -R appuser:appuser /app

# Copy binary and assets
COPY --from=builder /app/app .
COPY --chown=appuser:appuser ./static ./static

# Make the binary executable
RUN chmod +x /app/app

# Switch to non-root user
USER appuser

EXPOSE 8080

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./app"]
