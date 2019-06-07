PACKAGE = github.com/issmirnov/soteria
TARGET = soteria

SHELL := /bin/bash

# figure out current platform for deps.
UNAME := $(shell uname | tr A-Z a-z )

LDFLAGS = -s -w 

.DEFAULT_GOAL := all

.PHONY: \
	all \
	build \
	deps \
	debug \
	init

## Default action: debug.
all: debug

## Compile and run with sample file.
debug:
	go build -o "$(TARGET)" .
	# Contributors: provde the TOKEN and CHATID here if you wish. Just make sure not to commit and push!
	./$(TARGET) -f tests/hello_world.txt

## Compile and compress binary.
build:
	mkdir -p out
	GOOS=linux go build -ldflags "$(LDFLAGS)" -o "out/$(TARGET)-linux" .
	upx --brute out/$(TARGET)-linux
	GOOS=darwin go build -ldflags "$(LDFLAGS)" -o "out/$(TARGET)-darwin" .
	upx --brute out/$(TARGET)-darwin

## "go get" deps
deps:
	go get ./...

## Install system deps
init: deps
ifeq ($(UNAME), linux)
	sudo apt install upx-ucl 
endif
ifeq ($(UNAME), darwin)
	brew install upx 
endif


# Fancy help message
# Source: https://gist.github.com/prwhite/8168133
# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
TARGET_MAX_CHAR_NUM=20

## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
