PKG := "github.com/psyb0t/potentially-hazardous-asteroids"
PKG_LIST := $(shell go list $(PKG)/...)

all: build

dep: ## Get the dependencies + remove unused ones
	@go mod tidy
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status $(PKG_LIST)

build: dep ## Build the executable binary
	@go build -o build/potentially-hazardous-asteroids cmd/*.go

build-dockerimage: build ## Build the executable binary and rebuild docker image then remove build dir
	@docker build . --tag psyb0t/potentially-hazardous-asteroids
	@rm -rf build

run: dep ## Run without building
	@go run cmd/*.go

clean: ## Remove the build directory
	@rm -rf build

test: ## Run tests
	@go test -v $(PKG_LIST)

vet: ## Run go vet
	@go vet $(PKG_LIST)

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@cat cover.out >> coverage.txt

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
