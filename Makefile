
help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

compose-up: ## Run docker compose containers
	docker-compose up -d
	docker exec -it app sh

dev: ## Run the api in development mode
	go run main.go --dev

build: ## Build the binary
	CGO_ENABLED=0 go build -a -o piadocas .
