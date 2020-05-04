.PHONY: lint
lint: ## Run linter
	@go vet `go list ./...`

.PHONY: test
test: ## Run all tests
	@go test -v -cover ./...

.PHONY: help
help: ## Help command
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
