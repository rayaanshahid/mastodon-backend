# Use an official Golang runtime as a parent image
FROM golang:1.22.4

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o main .

# Make port 8000 available to the world outside this container
EXPOSE 8000

# Run the executable
CMD ["./main"]
