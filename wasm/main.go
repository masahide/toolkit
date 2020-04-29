package main

import (
	"encoding/base64"
	"fmt"
	"html"
	"syscall/js"
)

const (
	base64LineLength = 76
)

type base64conv struct {
	input     js.Value
	output    js.Value
	splitLine js.Value
}

func splitLine(s string, size int) string {
	res := ""
	for ; size > 0; s = s[size:] {
		if len(s) < size {
			size = len(s)
		}
		res += s[:size] + "\n"
	}
	return res
}

func (b *base64conv) encode(this js.Value, args []js.Value) interface{} {
	s := b.input.Get("value").String()
	base64str := base64.StdEncoding.EncodeToString([]byte(s))
	if b.splitLine.Get("checked").Bool() {
		base64str = splitLine(base64str, base64LineLength)
	}
	b.output.Set("innerHTML", base64str)
	return nil
}
func decode(in string) string {
	text, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return err.Error()
	}
	return string(text)
}
func (b *base64conv) decode(this js.Value, args []js.Value) interface{} {
	input := b.input.Get("value").String()
	value := html.EscapeString(decode(input))
	b.output.Set("innerHTML", value)
	fmt.Printf("value:%s\n", value)
	return nil
}

func registerCallbacks() {
	var document = js.Global().Get("document")
	getElementByID := func(targetID string) js.Value {
		return document.Call("getElementById", targetID)
	}
	dec := &base64conv{
		input:  getElementByID("base64_text"),
		output: getElementByID("decode_text"),
	}
	js.Global().Set("base64decode", js.FuncOf(dec.decode))
	enc := &base64conv{
		input:     getElementByID("input_text"),
		output:    getElementByID("base64_text"),
		splitLine: getElementByID("split_line"),
	}
	js.Global().Set("base64encode", js.FuncOf(enc.encode))
}

func main() {
	c := make(chan struct{}, 0)
	println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
