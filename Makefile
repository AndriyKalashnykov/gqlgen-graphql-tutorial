.DEFAULT_GOAL := help
GOFLAGS=-mod=mod

#help: @ List available tasks
help:
	@clear
	@echo "Usage: make COMMAND"
	@echo "Commands :"
	@grep -E '[a-zA-Z\.\-]+:.*?@ .*$$' $(MAKEFILE_LIST)| tr -d '#' | awk 'BEGIN {FS = ":.*?@ "}; {printf "\033[32m%-14s\033[0m - %s\n", $$1, $$2}'

#docker-up: @ Docker Compose Up
docker-up: docker-down
	docker-compose -f "docker-compose.yml" up

#docker-down: @ Docker Compose Down
docker-down:
	docker-compose -f "docker-compose.yml" down --volumes

#generate: @ Generate GraphQL go source code
generate:
	@rm -rf graph/models
	rm -rf graph/generated
	@export GOFLAGS=$(GOFLAGS); go run github.com/99designs/gqlgen generate

#test: @ Run tests
test: generate
	@export GOFLAGS=$(GOFLAGS); go test -v ./...

#build: @ Build Threeport GraphQL API
build: generate
	@export GOFLAGS=$(GOFLAGS); go build -o server server.go

#run: @ Run Threeport GraphQL API
run: build
	@export GOFLAGS=$(GOFLAGS); go run server.go -env-file=.env

#get: @ Download and install packages
get:
	@export GOFLAGS=$(GOFLAGS); go get . ; go mod tidy

#deps: @ Download and install dependencies
deps:
	@export GOFLAGS=$(GOFLAGS); go install github.com/99designs/gqlgen@latest

#update: @ Update dependencies to latest versions
update:
	@export GOPRIVATE=$(GOPRIVATE); export GOFLAGS=$(GOFLAGS); go get -u; go mod tidy
