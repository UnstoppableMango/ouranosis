_ != mkdir -p .make

GO     := go
BUF    := $(GO) tool buf
DEVCTL := $(GO) tool devctl

GO_SRC    != $(DEVCTL) list --go
PROTO_SRC != $(BUF) ls-files

build: bin/client bin/server .make/buf-build
fmt format: .make/buf-fmt .make/go-fmt
lint: .make/buf-lint .make/go-vet
tidy: go.sum buf.lock

bin/client bin/server: bin/%: go.mod ${GO_SRC}
	$(GO) build -o $@ ./cmd/$*

buf.lock: buf.yaml ${PROTO_SRC}
	$(BUF) dep update

go.sum: go.mod ${GO_SRC}
	$(GO) mod tidy

.make/buf-build: ${PROTO_SRC}
	$(BUF) build $(addprefix --path ,$?)
	@touch $@

.make/buf-fmt: ${PROTO_SRC}
	$(BUF) format --write $(addprefix --path ,$?)
	@touch $@

.make/buf-lint: ${PROTO_SRC}
	$(BUF) lint $(addprefix --path ,$?)
	@touch $@

.make/go-fmt: ${GO_SRC}
	$(GO) fmt $(addprefix ./,$(sort $(dir $?)))
	@touch $@

.make/go-vet: ${GO_SRC}
	$(GO) vet $(addprefix ./,$(sort $(dir $?)))
	@touch $@
