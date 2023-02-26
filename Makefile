.PHONY: build
build:
	go build ./cmd/app/main.go

.PHONY: run
run:
	go run ./cmd/app/main.go

.DEFAULT_GOAL := build