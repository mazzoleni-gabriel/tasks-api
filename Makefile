run:
	@docker-compose build
	@docker-compose up

test:
	go test ./...

fmt:
	gofmt -l -s -w .