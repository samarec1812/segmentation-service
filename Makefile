.PHONY: lint
lint:
	$(info Run go linters in project...)
	golangci-lint run -c ./.golangci.yaml ./...