# GO parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt

DEPCMD=dep
DEPRUN=$(DEPCMD) ensure

BINARY_NAME=alexa_local_server
BUILD_DIRECTORY=bin
BINARY_PATH=$(BUILD_DIRECTORY)/$(BINARY_NAME)

all: clean build test run
run: build
	./$(BINARY_PATH)
full: full_clean full_build


full_build: deps build 
deps:
	$(DEPRUN)
build: format
	$(GOBUILD) -o $(BINARY_PATH) .
# Build static image with no outwards deps. 
build_static:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -a -installsuffix nocgo -ldflags '-w -extldflags "-static"' -o $(BINARY_PATH) -v .
docker_build:
	docker build -t aa/alexa-local-server .

test:
	$(GOTEST) -v ./...
test_local:
	$(GOTEST) -v ./... -tags=local

full_clean: dep_clean clean
dep_clean: 
	rm -rf ./vendor
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)

format:
	$(GOFMT) ./...

