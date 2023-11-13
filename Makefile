
.PHONY: test


lint:
	go fmt ./...
	go vet ./...

test:
	@mkdir -p test
	go test ./... -race -coverprofile=test/coverage.out ./...

