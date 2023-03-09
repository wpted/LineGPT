FROM        golang:1.20-alpine
WORKDIR     /app

# Download dependencies
COPY        go.mod ./
COPY        go.sum ./
RUN         go mod download

# Add code to the container
COPY        cmd/ ./cmd/
COPY        configs/ ./configs/
COPY        pkg/ ./pkg/
COPY        . .

RUN         go build -o /main ./cmd/main.go

EXPOSE      8000
ENTRYPOINT  ["/main"]
