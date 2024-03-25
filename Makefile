.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: test
test: tidy
	@go test -v ./...
