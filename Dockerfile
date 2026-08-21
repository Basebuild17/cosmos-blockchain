# Multi-stage build for smaller final image
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make gcc musl-dev linux-headers

# Copy source code
COPY . .

# Download dependencies
RUN go mod download

# Build the binary
RUN make build

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates curl jq bash

# Create app user
RUN addgroup -g 1000 myblockchain && \
    adduser -D -u 1000 -G myblockchain myblockchain

# Set working directory
WORKDIR /home/myblockchain

# Copy binary from builder
COPY --from=builder /app/myblockaind /usr/local/bin/
COPY --from=builder /app/myblockchainicli /usr/local/bin/

# Create necessary directories
RUN mkdir -p /home/myblockchain/.myblockchain && \
    chown -R myblockchain:myblockchain /home/myblockchain

# Switch to non-root user
USER myblockchain

# Expose ports
# 26656 - P2P
# 26657 - RPC
# 1317 - REST API
# 9090 - gRPC
EXPOSE 26656 26657 1317 9090

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=40s --retries=3 \
    CMD curl -f http://localhost:26657/status || exit 1

# Default command
ENTRYPOINT ["myblockaind"]
CMD ["start"]
