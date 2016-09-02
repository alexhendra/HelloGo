package main

import (
	"time"
	"math/rand"
	"fmt"
	"runtime"
)

// Timeout digunakan untuk mengontrol penerimaan data dari channel berdasarkan waktu diterimanya, dengan -
// durasi timeout bisa ditentukan sendiri

// Ketika tidak ada aktivitas penerimaan data selama durasi tersebut, akan memicu callback yg isinya juga ditentukan sendiri

// chan parameter hanya utk mengirim
func sendData(ch chan <- int)  {
	for i:=0; true; i++ {
		ch <- i
		detik := time.Duration(rand.Int() % 10 + 1)
		fmt.Println(detik)
		time.Sleep(detik * time.Second)
	}
}

// chan parameter hanya utk menerima
func retreiveData(ch <- chan int)  {
	loop:

	for{
		select {
		case data:= <- ch:
			fmt.Println(`Receive data "`, data, `"`,"\n")

		// untuk mengatur apabila setelah 5 detik, maka statement-nya dijalankan
		case <- time.After(time.Second * 5) :
			fmt.Println(time.Second)
			fmt.Println("Timeout. no activities under 5 seconds")
			break loop
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var messages = make(chan int)

	go sendData(messages)
	retreiveData(messages)
}
