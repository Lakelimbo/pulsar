# "Pulsar" - Go starting project

This is an API built in Go as a starting point for my other web-related projects. It is probably **_not as good_** as other starting projects you may find around, but this is also serving as a study project for me.

## Requirements

- Go (obviously)
- Docker
- Just (not mandatory, but recommended)

## Start the project

Before everything, you might want to tidy up the dependencies:

```sh
go mod tidy
```

Run the project with Just:

```sh
just run
```

Or if you prefer not using Just, you can run in the classic way:

```sh
# point to the directory where the main.go lives
go run ./cmd/pulsar serve
```

## Other commands

You can find other commands with Just:

```sh
just

# outputs:
#
# Available commands:
#    run      - Run the application
#    build    - Build the application
#    fmt      - Format all codebase
```

## TO-DO

- [x] Basic API functionality
  - [x] Routing
- [x] PostgreSQL for the database
- [x] Redis for caching
- [ ] Make routes OpenAPI compliant
  - [ ] Add an OpenAPI explorer frontend (e.g. Swagger, ReDocly, Scalar...)
- [ ] OpenTelemetry
  - [ ] Tracing
  - [ ] Logging
- [ ] Frontend
- [ ] Containerize backend
- [ ] Better docs overall
