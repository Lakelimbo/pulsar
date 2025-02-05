default: list

list:
    @echo "Available commands:"
    @echo "     run      - Run the application"
    @echo "     build    - Build the application"
    @echo "     fmt      - Format all codebase"
    @echo "     swag-gen - Regenerate swagger reference file + docs.go" 

run:
    go run ./cmd/pulsar serve

build:
    go build -o bin/app ./cmd/pulsar

fmt:
    go fmt ./... && swag fmt

swag-gen:
    swag init --dir ./internal/server -g routes.go -o ./internal/server/docs