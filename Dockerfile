# Stage 1: Build the React application
FROM node:18-alpine AS frontend-build

# Set working directory
WORKDIR /app/frontend

# Copy React app files
COPY frontend/package*.json ./
COPY frontend/ ./

# Install dependencies and build the React app
RUN npm install && npm run build

# Stage 2: Build the Go application
FROM golang:1.20-alpine AS backend-build

# Set working directory
WORKDIR /app/backend

# Copy Go app files
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./

# Embed the React build files into the Go binary (if applicable)
COPY --from=frontend-build /app/frontend/build ./static

# Build the Go binary
RUN go build -o server .

# Stage 3: Final image
FROM alpine:latest

# Install a simple static file server (optional)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy the Go binary
COPY --from=backend-build /app/backend/server .

# Copy React build files (if not embedded in Go binary)
COPY --from=frontend-build /app/frontend/build ./static

# Expose the necessary port
EXPOSE 8080

# Command to run the Go server
CMD ["./server"]