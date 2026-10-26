GO_SRC   ?= $(shell find . -name '*.go')
PACKAGES ?= $(shell go list ./... 2>/dev/null)
ARCH_DIR ?= docs/architecture
LIKEC4   ?= likec4

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

arch:
	${LIKEC4} start ${ARCH_DIR}

arch-check:
	nix flake check

arch-build:
	nix build .#architecture
