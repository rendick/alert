# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOGET = $(GOCMD) get
BINARY_NAME = alert
INSTALL_PATH = /usr/local/bin

all: clean build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

install:
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)"
	cp $(BINARY_NAME) $(INSTALL_PATH)

.PHONY: all build clean install
