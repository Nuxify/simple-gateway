# set default shell
SHELL = bash -e -o pipefail

# variables
VERSION                  ?= $(shell cat ./VERSION)

default: run

.PHONY:	install
install:
	go mod tidy
	go mod vendor

.PHONY:	lint
lint:
	golangci-lint run 

.PHONY:	build
build:
	mkdir -p bin
	go build -race -o bin/api-gateway \
	    main.go

.PHONY:	test
test:
	go test -race -v -p 1 ./...
	
.PHONY:	run
run:	build
	./bin/api-gateway

.PHONY: up
up:
	docker-compose down
	docker-compose up -d --build

