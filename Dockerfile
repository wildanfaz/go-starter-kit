# Stage 1: Build the Golang binary
FROM golang:1.23.4-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download and install Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-starter-kit .

# Stage 2: Create a minimal image
FROM alpine:latest

# Set the timezone to Asia/Jakarta
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone && \
    apk del tzdata

# Set the working directory inside the container
WORKDIR /app

# Copy only the compiled binary from the previous stage
COPY --from=builder /app/go-starter-kit .

# Expose the port your application will run on
EXPOSE 1323

# Command to run the application
CMD ["./go-starter-kit", "start"]