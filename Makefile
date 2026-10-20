GO_SRC   ?= $(shell find . -name '*.go')
PACKAGES ?= $(shell go list ./... 2>/dev/null)
ARCH_DIR ?= docs/architecture
LIKEC4   ?= cd ${ARCH_DIR} && npx --no-install likec4

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

arch: ${ARCH_DIR}/node_modules
	${LIKEC4} start

arch-check: ${ARCH_DIR}/node_modules
	${LIKEC4} validate

# Relative base and hash history let one build serve from any URL prefix.
arch-build: ${ARCH_DIR}/node_modules
	${LIKEC4} build --base ./ --use-hash-history --title 'ouranosis architecture' -o dist

${ARCH_DIR}/node_modules: ${ARCH_DIR}/package.json ${ARCH_DIR}/package-lock.json
	npm ci --prefix ${ARCH_DIR}
	@touch $@
