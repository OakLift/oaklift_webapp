# Use a multi-stage build to minimize the final image size
# Stage 1: Build the Go application
FROM golang:1.18 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the application
RUN go build -o oaklift ./main.go ./oauth.go

# Stage 2: Create a smaller image with the application
FROM alpine:3.14

# Install ca-certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/oaklift .

ENV DB_USER=${DB_USER}
ENV DB_NAME=${DB_NAME}
ENV DB_PORT=${DB_PORT}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_HOST=${DB_HOST}
ENV GOOGLE_CLIENT_ID=${GOOGLE_CLIENT_ID}
ENV GOOGLE_CLIENT_SECRET=${GOOGLE_CLIENT_SECRET}


# Expose the application port
EXPOSE 8000

# Run the application
CMD ["./oaklift"]
