# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o html2pdf .

# Runtime stage - using Debian for wkhtmltopdf availability
FROM debian:bookworm-slim

WORKDIR /app

# Install wkhtmltopdf and dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    wkhtmltopdf \
    fonts-dejavu-core \
    fonts-freefont-ttf \
    fontconfig \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && fc-cache -f

# Copy binary from builder
COPY --from=builder /app/html2pdf .

# Expose port
EXPOSE 8080

# Set environment variables
ENV PORT=8080

# Run the application
CMD ["./html2pdf"]
