WORKING_DIRS=tmp
SRC=$(shell find . -name "*.go")
BIN=tmp/$(shell basename $(CURDIR))
FMT=tmp/fmt

.PHONY: all clean cover test

all: $(WORKING_DIRS) $(FMT) $(BIN)

clean:
	rm -rf $(WORKING_DIRS)

$(WORKING_DIRS):
	mkdir -p $(WORKING_DIRS)

$(FMT): $(SRC)
	go fmt ./... > $(FMT) 2>&1 || true

$(BIN): $(SRC)
	go build -o $(BIN)
