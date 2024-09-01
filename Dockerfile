# Use the official Golang image
FROM golang:1.17 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o crm-backend .

# Use a smaller image for the final build
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/crm-backend .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./crm-backend"]