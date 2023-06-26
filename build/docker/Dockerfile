# Use the official Go image as the base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

COPY ./cmd/api .
# Copy the Go module files to the working directory
COPY go.mod .
COPY go.sum .

# Download and install the application dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o app

# Expose the port that the application listens on
EXPOSE 8000

# Set the entry point for the container
CMD ["./app"]