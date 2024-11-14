# Base image
FROM golang:1.19

# Set up work directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o /go/go_app

RUN chown -R 1000:1000 /go

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./go_app"]
