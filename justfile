default: list

list:
    @echo "Available commands:"
    @echo "    run      - Run the application"
    @echo "    build    - Build the application"
    @echo "    fmt      - Format all codebase"

run:
    go run ./cmd/pulsar serve

build:
    go build -o bin/app ./cmd/pulsar

fmt:
    go fmt ./...