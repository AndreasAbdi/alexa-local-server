# GO parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
BINARY_NAME=alexa_local_server
BUILD_DIRECTORY=bin
BINARY_PATH=$(BUILD_DIRECTORY)/$(BINARY_NAME)

all: clean build test 
run: build
	./$(BINARY_PATH)
build: format
	$(GOBUILD) -o $(BINARY_PATH) -v ./...
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)
format:
	$(GOFMT) ./...