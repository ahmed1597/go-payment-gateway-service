# Use official Golang image as the base
FROM golang:1.20-alpine

# Set environment variables
ENV GO111MODULE=on

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o /app/main ./cmd/main.go

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the Go application
CMD ["/app/main"]
