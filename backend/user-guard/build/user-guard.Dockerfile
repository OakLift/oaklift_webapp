# syntax=docker/dockerfile:1

FROM golang:1.19 as builder

# Set destination for COPY
WORKDIR /app

COPY ../go.mod ../go.sum ./
COPY ../user-guard/main.go ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/user-guard

ENTRYPOINT ["/opt/user-guard"]
