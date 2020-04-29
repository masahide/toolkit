package main

import (
	"encoding/base64"
	"fmt"
	"html"
	"syscall/js"
)

var document = js.Global().Get("document")

func getElementByID(targetID string) js.Value {
	return document.Call("getElementById", targetID)
}

func base64encode(this js.Value, args []js.Value) interface{} {
	value := getElementByID(args[1].String()).Get("value").String()
	getElementByID(args[0].String()).Set("innerHTML", value)
	return nil
}
func decode(in string) string {
	text, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return err.Error()
	}
	return string(text)
}
func base64decode(this js.Value, args []js.Value) interface{} {
	input := getElementByID(args[1].String()).Get("value").String()
	value := html.EscapeString(decode(input))
	getElementByID(args[0].String()).Set("innerHTML", value)
	fmt.Printf("value:%s\n", value)
	return nil
}
func add(this js.Value, args []js.Value) interface{} {
	js.Global().Set("output", js.ValueOf(args[0].Int()+args[1].Int()))
	fmt.Printf("%v\n", js.ValueOf(args[0].Int()+args[1].Int()))
	return nil
}

func registerCallbacks() {
	//obj := getElementByID("base64_text")
	//obj.Call("addEventListener", "keyup", base64decode)
	js.Global().Set("base64encode", js.FuncOf(base64encode))
	js.Global().Set("base64decode", js.FuncOf(base64decode))
	js.Global().Set("add", js.FuncOf(add))
}

func main() {
	c := make(chan struct{}, 0)
	println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
