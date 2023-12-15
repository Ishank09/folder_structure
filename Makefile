TARGET = "main.go"
GOPATH = $(shell go env GOPATH)

ARGS ?= ${GOPATH}/src/snap_roles


run:
	go run main.go $(ARGS)

clean: 
	go clean

clean_run: run clean