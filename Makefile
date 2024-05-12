GO := go
TARGET := cmd/mapz/main.go
OUTPUT := bin/mapz

all: build

build:
	$(GO) build -o $(OUTPUT) $(TARGET)

clean:
	rm -f $(OUTPUT)

install:
	@echo "Installing $(OUTPUT) to /usr/local/bin..."
	cp $(OUTPUT) /usr/local/bin

.PHONY: all build clean install
