#===============================================================================
# Environmrntal variables and constants
#===============================================================================
CURRENT_UID :=$(shell id -u):$(shell id -g)
GOLANGCI_LINT_TAG := v2.4.0

#===============================================================================
# Documentation
#===============================================================================
help:
	@echo "Makefile commands:"
	@echo "  start   - Start the backend template service"
	@echo "  stop    - Stop the backend template service"
	@echo "  logs    - View logs for the backend template service"
	@echo "  tests   - Run tests for the backend template service"
	@echo "  help    - Show this help message"

#===============================================================================
# Service management
#===============================================================================
# TODO: replace project name and docker-compose service name with actual values
start:
	env CURRENT_UID=${CURRENT_UID} docker compose -f zarf/local/docker-compose.yaml --project-name bkndtmp up -d --build backend_template

stop:
	env CURRENT_UID=${CURRENT_UID} docker compose -f zarf/local/docker-compose.yaml --project-name bkndtmp down

logs:
	env CURRENT_UID=${CURRENT_UID} docker compose -f zarf/local/docker-compose.yaml --project-name bkndtmp logs -f backend_template

#===============================================================================
# Development tools
#===============================================================================
tests:
	env CURRENT_UID=${CURRENT_UID} docker compose -f zarf/local/docker-compose.yaml --project-name bkndtmp exec backend_template sh -c "go clean -testcache && CI_PIPELINE_ID=true go test ./... -v -race -short -cover -json | tparse"

lint:
	docker run --env-file zarf/local/.env --rm -v $(shell pwd):/api -w /api --memory=4g golangci/golangci-lint:$(GOLANGCI_LINT_TAG) golangci-lint run -v --timeout 4m 