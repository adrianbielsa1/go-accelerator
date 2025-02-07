# Use the official Golang image as the build stage.
FROM golang:latest AS builder

WORKDIR /app

# Copy go modules and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the application.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o accelerator-service .

# Use a minimal base image for running the application.
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage.
COPY --from=builder /app/accelerator-service .
RUN chmod +x "/app/accelerator-service"

# Expose the application port.
EXPOSE 8080

# Run the application.
CMD ["/app/accelerator-service"]
