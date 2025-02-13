BINARY_NAME = mythpromo
BIN_DIR = ../../bin
BINARY = $(BIN_DIR)/$(BINARY_NAME)
GO = go
GOFLAGS = -ldflags="-s -w"
RM = rm -rf

all: build

build:
	@$(GO) build $(GOFLAGS) -o $(BINARY) .

clean:
	@$(RM) $(BINARY)

run: build
	@./$(BINARY)

deps:
	@$(GO) mod download

fmt:
	@$(GO) fmt ./...

test:
	@$(GO) test ./...

$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

build: $(BIN_DIR)

.PHONY: all build clean run deps fmt test
