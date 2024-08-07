# Use a Golang base image to build the Go app
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Use a minimal base image for the runtime
FROM scratch

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main /app/main

# Set the entrypoint for the container
ENTRYPOINT ["/app/main"]
