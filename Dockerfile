# Frontend build stage
FROM node:20-alpine AS frontend-builder
WORKDIR /app
# Copy package files first for better caching
COPY frontend/package*.json ./
RUN npm ci
# Copy frontend source
COPY frontend/ ./
# Build frontend
RUN npm run build

# Backend build stage
FROM golang:1.23-alpine AS backend-builder
WORKDIR /app
# Copy go.mod and go.sum first for better caching
COPY backend/go.* ./
RUN go mod download
# Copy backend source
COPY backend/ ./
# Build backend
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Final stage
FROM alpine:3.20
WORKDIR /app
# Copy backend binary
COPY --from=backend-builder /app/server .
# Copy frontend dist files to the expected path
COPY --from=frontend-builder /app/dist/frontend/browser/ ./frontend/dist/frontend/browser/

# Set production mode
ENV GIN_MODE=release

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["/app/server"]
