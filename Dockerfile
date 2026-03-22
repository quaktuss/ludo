FROM golang:1.21-alpine AS builder

# Install dependencies for building
RUN apk add --no-cache git

WORKDIR /app

# Copy source code first
COPY . .

# Download dependencies and generate go.sum
RUN go mod tidy && go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ludo-bot cmd/bot/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/ludo-bot .

# Copy config files
COPY --from=builder /app/configs ./configs

# Create non-root user
RUN adduser -D -s /bin/sh appuser
USER appuser

# Expose port (if needed for health checks)
EXPOSE 8080

# Run the bot
CMD ["./ludo-bot"]