package main

import (
	"runtime"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(2)

	// variable messages di deklarasikan sebagai channel
	// formatnya : make(chan tipedatanya)
	var messages = make(chan string)

	// Closure sayHelloTo
	var sayHelloTo = func(who string) {
		var data =  fmt.Sprintf("Hello %s", who)

		// variable data dikirimkan ke channel "messages"
		// karena tanda panah ke kiri " <- "
		messages <- data
	}

	go sayHelloTo("Alex Hendra")
	go sayHelloTo("Bambang Joyo")
	go sayHelloTo("Lala Lili")

	// data yg melalui channel "messages" disimpan ke variable "msg1"
	var msg1 = <- messages
	fmt.Println(msg1)

	var msg2 = <- messages
	fmt.Println(msg2)

	var msg3 = <- messages
	fmt.Println(msg3)
}
