FROM golang:1.22

WORKDIR /app

# Install CompileDaemon for live reloading
RUN go install github.com/githubnemo/CompileDaemon@latest

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

# Automatically rebuild and run on source change
CMD ["CompileDaemon", "--build=go build -o main .", "--command=./main"]
