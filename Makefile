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
APP_PATH=./cmd/alexa_local_server
BUILD_CONFIG_DIRECTORY=build
DOCKER_BUILD_FILE=$(BUILD_CONFIG_DIRECTORY)/Dockerfile

all: clean build test run
run: build
	./$(BINARY_PATH)
docker_deploy:
	docker run -d -p 8000:8000 aa/alexa_local_server
full: full_clean full_build

full_build: deps build 
deps:
	$(DEPRUN)
build: format
	$(GOBUILD) -o $(BINARY_PATH) $(APP_PATH)
# Build static image with no outwards deps. 
build_static:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -a -installsuffix nocgo -ldflags '-w -extldflags "-static"' -o $(BINARY_PATH) -v $(APP_PATH)
docker_build:
	docker build -t aa/alexa-local-server -f $(DOCKER_BUILD_FILE) .

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

