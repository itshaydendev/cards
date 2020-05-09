GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=cards

all: build
build:
	$(GOBUILD) -o build/$(BINARY_NAME) -v github.com/itshaydendev/cards/cmd/cards
