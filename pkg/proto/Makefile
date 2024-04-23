PROTO_DIR := .
GO_OUT_DIR := $(PROTO_DIR)
PROTOC := protoc

# Source and target files
PROTO_FILES := $(shell find $(PROTO_DIR) -name '*.proto')
GO_OUT_FILES := $(patsubst %.proto,%.pb.go,$(PROTO_FILES))

# Default target
all: $(GO_OUT_FILES)

# Rule for generating Go files from proto files
%.pb.go: %.proto
	$(PROTOC) --go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
	          --go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative \
	          -I$(PROTO_DIR) $<

# Clean generated files
clean:
	rm -f $(GO_OUT_FILES)

# Help
help:
	@echo "Usage:"
	@echo "  make          # Compile all proto files"
	@echo "  make clean    # Remove all generated files"
