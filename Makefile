
dist   ?= ./public
src    ?= ./wasm/
goroot  = $(shell go env GOROOT)

build:
	GOOS=js GOARCH=wasm go get ./...
	GOOS=js GOARCH=wasm go build -o $(dist)/main.wasm $(src)
	cp "$(goroot)/misc/wasm/wasm_exec.js" $(dist)/


run:
	go run httpsv.go
