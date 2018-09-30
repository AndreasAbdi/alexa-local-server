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
	$(GOBUILD) -o $(BINARY_PATH) .
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)
format:
	$(GOFMT) ./...
#Build static image with no outwards deps. 
build_static:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -a -installsuffix nocgo -ldflags '-w -extldflags "-static"' -o $(BINARY_PATH) -v .
docker_build:
	docker build -t AndreasAbdi/alexa-local-server .