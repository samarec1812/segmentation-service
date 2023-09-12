.PHONY: lint
lint:
	$(info Run go linters in project...)
	golangci-lint run -c ./.golangci.yaml ./...


.PHONY: run
run:
	$(info Run project in docker-compose...)
	docker-compose up