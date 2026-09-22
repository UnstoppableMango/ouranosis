GO_SRC   ?= $(shell find . -name '*.go')
PACKAGES ?= $(shell go list ./... 2>/dev/null)

build: $(if ${PACKAGES},bin/ouranosis)

bin/ouranosis: go.mod ${GO_SRC}
	go build -o $@ ./cmd/ouranosis

run: bin/ouranosis
	./bin/ouranosis

test:
	$(if ${PACKAGES},go tool ginkgo run -r,@echo no packages to test)

check lint:
	$(if ${PACKAGES},go vet ./...,@echo no packages to check)

format fmt:
	gofmt -w .

tidy: go.sum

go.sum: go.mod ${GO_SRC}
	go mod tidy
