all: help

.PHONY:
help:
	@echo "tasks api"
	@echo ""
	@echo "run                       Build docker and run application"
	@echo "test                      Run tests locally"
	@echo "fmt                       Run go fmt locally"
	@echo ""
	@echo "swagger doc"
	@echo ""
	@echo "generate-index            generate swagger documentation on index.yaml"
	@echo "serve                     Creates a server to render the content of index.yaml"
	@echo ""

run:
	@docker-compose build
	@docker-compose up

test:
	go test ./...

fmt:
	gofmt -l -s -w .


generate-index:
	@swagger generate spec -o ./docs/index.yaml --scan-models

serve:
	@swagger-ui-watcher docs/index.yaml -p 8081