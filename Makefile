.DEFAULT_GOAL := run

.PHONY:fmt vet build run govulncheck staticcheck revive

.SILENT:

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build main.go -o app

run:
	go run main.go