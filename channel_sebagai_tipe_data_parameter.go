package main

import (
	"fmt"
	"runtime"
)

// Channel sebagai tipe data di parameter fungsi "printMessages"
// Data yg dikirimkan lewat channel di paramater bersifat "Passing by Reference"
func printMessages(what chan string)  {
	fmt.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _,each := range []string{"Doli","Badrul","Sumimar"} {
		go func(who string) {
			var data = fmt.Sprintf("Hello %s", who)
			messages <- data
		}(each) // each merupakan nilai yang mengisi fungsi closure yang digunakan GOROUTINE
	}

	for i:=0;i<3;i++ {
		printMessages(messages)
	}
}
