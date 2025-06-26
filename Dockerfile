# Dockerfile

FROM golang:1.22-alpine

# Install system deps and CompileDaemon
RUN apk add --no-cache git
RUN go install github.com/githubnemo/CompileDaemon@latest

# Set working directory
WORKDIR /app

# Copy go.mod and install deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the full project
COPY . .

# Expose app port
EXPOSE 8080

# Run CompileDaemon (rebuild & restart on change)
CMD ["CompileDaemon", "--build=go build -o main ./cmd/app", "--command=./main"]
