BUILD_DIR=build
BIN_NAME=dotme
OUT=$(BUILD_DIR)/$(BIN_NAME)

GO=go
GOBUILD=$(GO) build -o $(OUT)
GOCLEAN=$(GO) clean

all: build

build: cmd/dotme
	$(GOBUILD) -o $(OUT) -v ./$<

clean:
	rm -rf $(BUILD_DIR)
	$(GOCLEAN)

.PHONY: build clean
