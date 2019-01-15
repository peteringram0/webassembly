package main

import (
	"fmt"
	"syscall/js"
)

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
}

func main() {

	fmt.Println("Hello, WebAssembly!")

	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c

}

func add(i []js.Value) {
	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
	data := js.ValueOf(i[0].Int() + i[1].Int()).String()
	println(data)
}
