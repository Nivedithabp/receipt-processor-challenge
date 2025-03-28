# Stage 1: Build the Go application
FROM golang:1.24.1 AS builder
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application files
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create the final image with Alpine (lightweight and secure)
FROM alpine:latest
WORKDIR /app

# Add required certificates and utilities
RUN apk --no-cache add ca-certificates

# Copy the Go binary and config.json
COPY --from=builder /app/main .
COPY config.json ./config.json
COPY docs ./docs

# Set the environment for config
ENV CONFIG_FILE=config.json

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]

