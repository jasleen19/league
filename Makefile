BINARY_DIR=bin
BINARY_NAME=server

all: build run

build:
	@mkdir -p $(BINARY_DIR)
	go build -o ./$(BINARY_DIR)/$(BINARY_NAME) .

run:
	./$(BINARY_DIR)/$(BINARY_NAME)

fmt:
	gofmt

test:
	go test ./...

clean:
	go mod tidy
	go clean
	rm -f $(BINARY_NAME)