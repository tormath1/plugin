TOOL_BIN := $(PWD)/bin
ENUMER := $(TOOL_BIN)/enumer

$(ENUMER):
	@GOBIN=$(TOOL_BIN) go install github.com/dmarkham/enumer

.PHONY: generate
generate: $(ENUMER)
	@go generate ./...
