_ != mkdir -p .make

export GOBIN := ${CURDIR}/bin

GO     ?= go
BUF    ?= $(GO) tool buf
BUN    ?= bun
DEVCTL ?= $(GO) tool devctl
DOCKER ?= docker
GINKGO ?= $(GO) tool ginkgo

GO_SRC      != $(DEVCTL) list --go
TS_SRC      != $(DEVCTL) list --ts
PROTO_SRC   != $(BUF) ls-files
GO_PB_SRC   := ${PROTO_SRC:proto/%.proto=gen/%.pb.go}
GO_GRPC_SRC := gen/dev/unmango/ouranosis/v1alpha1/server_grpc.pb.go

build: bin/client bin/server bin/wui .make/buf-build
gen generate: ${GO_PB_SRC}
test: .make/ginkgo-run
fmt format: .make/buf-fmt .make/go-fmt
lint: .make/buf-lint .make/go-vet
tidy: go.sum buf.lock
docker: bin/wui.tar

frontend: bin/wui
start-frontend:
	$(BUN) run --cwd cmd/wui start
world: bin/world

${GO_PB_SRC} ${GO_GRPC_SRC} &: buf.gen.yaml ${PROTO_SRC}
	$(BUF) generate $(addprefix --path ,$(filter ${PROTO_SRC},$?))

%_suite_test.go:
	cd $(dir $@) && $(GINKGO) bootstrap
%_test.go:
	cd $(dir $@) && $(GINKGO) generate $(notdir $@)

bin/wui: cmd/wui/dist/index.html
$(addprefix bin/,client server world wui): bin/%: go.mod ${GO_SRC}
	$(GO) build -o $@ ./cmd/$*

bin/wui.tar: cmd/wui/Dockerfile cmd/wui/main.go ${TS_SRC}
	$(DOCKER) build ${CURDIR} --file $< --output type=tar,dest=$@
	$(DOCKER) import $@ ouranosis-wui

cmd/wui/dist/index.html: .make/bun-install
	$(BUN) run --cwd cmd/wui build

bin/buf: go.mod ## Optional bin install for editor integration
	$(GO) install github.com/bufbuild/buf/cmd/buf

bin/devctl: go.mod ## Optional bin install
	$(GO) install github.com/unmango/devctl

bin/ginkgo: go.mod ## Optional bin install
	$(GO) install github.com/onsi/ginkgo/v2/ginkgo

buf.lock: buf.yaml ${PROTO_SRC}
	$(BUF) dep update

go.sum: go.mod ${GO_SRC}
	$(GO) mod tidy

.envrc: hack/example.envrc
	cp $< $@ && chmod u+r $@

.make/buf-build: ${PROTO_SRC}
	$(BUF) build $(addprefix --path ,$?)
	@touch $@

.make/buf-fmt: ${PROTO_SRC}
	$(BUF) format --write $(addprefix --path ,$?)
	@touch $@

.make/buf-lint: ${PROTO_SRC}
	$(BUF) lint $(addprefix --path ,$?)
	@touch $@

.make/bun-install: cmd/wui/package.json
	$(BUN) install --cwd cmd/wui
	@touch $@

.make/go-fmt: ${GO_SRC}
	$(GO) fmt $(addprefix ./,$(sort $(dir $?)))
	@touch $@

.make/ginkgo-run: ${GO_SRC}
	$(GINKGO) $(sort $(dir $?))
	@touch $@

.make/go-vet: ${GO_SRC}
	$(GO) vet $(addprefix ./,$(sort $(dir $?)))
	@touch $@
