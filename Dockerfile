# Use the official Go image as the base image
FROM golang:1.22.2-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy the source code into the container
COPY . .

# Build the Go application
RUN templ generate && go build -o ./tmp/main cmd/app/main.go
# Start a new stage from scratch
FROM scratch

# Set the current working directory inside the container

# Copy the executable from the builder stage into the final image
COPY --from=builder /app/tmp/main/ .
COPY web/ web/

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
