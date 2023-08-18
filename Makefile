BUILD_DIR=build
BIN_NAME=dotme
OUT=$(BUILD_DIR)/$(BIN_NAME)

GO=go
GOBUILD=$(GO) build -o $(OUT)
GOCLEAN=$(GO) clean

all: build

build: clean
	mk +$(BUILD_DIR)
	$(GOBUILD) -v ./cmd/dotme

run:
	./$(OUT)

clean:
	rm -rf $(BUILD_DIR)
	$(GOCLEAN)

.PHONY: build run clean
