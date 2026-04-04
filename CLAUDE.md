# CLAUDE.md

## Project Overview

Go GraphQL API tutorial using [gqlgen](https://gqlgen.com/) code-first framework. Exposes a `/query` endpoint with a GraphQL playground at `/`. Uses PostgreSQL for persistence via go-pg, Chi router for HTTP, and JWT for authentication.

## Tech Stack

- **Language**: Go 1.26
- **GraphQL**: gqlgen (code-first schema generation)
- **Database**: PostgreSQL via go-pg/v10
- **Router**: go-chi/v5
- **Auth**: JWT (dgrijalva/jwt-go)

## Project Structure

```
server.go              # Entrypoint - HTTP server setup
schema.graphql         # GraphQL schema definition
gqlgen.yml             # gqlgen code generation config
domain/                # Business logic layer
graphql/               # Generated resolvers, dataloaders, validation
models/                # Data models (User, Meetup) + validation
postgres/              # Repository implementations (go-pg)
middleware/             # Auth middleware
validator/             # Input validation helpers
```

## Build / Test / Run

```bash
make generate   # Regenerate GraphQL Go source from schema
make test       # Generate + run tests
make build      # Generate + compile binary
make run        # Build + start server (reads .env)
make get        # Download deps + tidy
make update     # Update all deps to latest
make deps       # Install gqlgen CLI tool
```

## Development Notes

- Code generation: run `make generate` after changing `schema.graphql`
- Server listens on `PORT` env var (default 8080)
- Docker: `make docker-up` / `make docker-down` for local PostgreSQL
- Generated files: `graphql/generated.go`, `models/models_gen.go`
