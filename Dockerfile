# Use an official Golang runtime as a parent image
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Install Compile Daemon for go. We'll use it to watch changes in go files
RUN go install github.com/githubnemo/CompileDaemon@latest

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Ensure go mod download for any new dependencies added after copying the source
RUN go mod tidy

# Build the Go app
RUN go build -o main ./cmd

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
