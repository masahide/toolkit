
dist ?= ./public

src  ?= ./wasm/

export GOOS = js
export GOARCH = wasm

GOROOT=$(shell go env GOROOT)

build:
	go get ./...
	go build -o $(dist)/main.wasm $(src)
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" $(dist)/


