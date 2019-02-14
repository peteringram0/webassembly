package main

import (
	"fmt"
	"syscall/js"
)

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
	js.Global().Set("fib", js.NewCallback(FibonacciRecursion))
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

func FibonacciRecursion(n int) {
	if n <= 1 {
		data := js.ValueOf(n.Int())
		println(data)
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
