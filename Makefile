GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=x2j

all: test build
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
test:
		$(GOCMD) test