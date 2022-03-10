ifeq (, $(shell which golint))
	$(error "golint is not installed; go install golang.org/x/lint/golint@latest")
endif

ifeq (, $(shell which golangci-lint))
	$(error "golangci-lint is not installed; go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0")
endif

.PHONY: lint
lint: deps
	golint -set_exit_status ./...
	go vet ./...
	golangci-lint run

.PHONY: deps
deps:
	go get -v -t -d ./...
