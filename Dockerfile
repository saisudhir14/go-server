# Use the official Go image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o go-server .

# Expose the port your app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./go-server"]


##### has no mod file