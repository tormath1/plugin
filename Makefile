TOOL_BIN := $(PWD)/bin
ENUMER := $(TOOL_BIN)/enumer

$(ENUMER):
	@GOBIN=$(TOOL_BIN) go install github.com/dmarkham/enumer

.PHONY: generate
generate: $(ENUMER)
	@go generate ./...

.PHONY: test
test: generate
	@go build -buildmode=plugin -o testdata/plugin.so testdata/plugin.go
	@go test -v ./...
	@rm -f testdata/plugin.so

