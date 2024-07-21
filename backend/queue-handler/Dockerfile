# Use the official Golang image as the base image
FROM golang:1.22.3-alpine

# Install curl
RUN apk update && apk add curl

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .


# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 80

# Command to run the application
CMD ["./main"]


