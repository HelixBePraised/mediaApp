# Author: Jackson Taylor
# Date: 02/06/2021

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
GOINSTALL=$(GOCMD) install
BINARY_NAME=jama

all: build
build:
		$(GOBUILD) -o $(BINARY_NAME) -v
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
run:
	    $(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
install:
	    $(GOINSTALL)
