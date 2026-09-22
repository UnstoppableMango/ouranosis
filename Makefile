GO_SRC ?= $(shell find . -name '*.go')

build: bin/ouranosis

bin/ouranosis: go.mod ${GO_SRC}
	go build -o $@ ./cmd/ouranosis

run: bin/ouranosis
	./bin/ouranosis

test:
	go tool ginkgo run -r

check lint:
	go vet ./...

format fmt:
	gofmt -w .

tidy: go.sum

go.sum: go.mod ${GO_SRC}
	go mod tidy
