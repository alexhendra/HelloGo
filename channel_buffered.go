package main

import (
	"runtime"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(2)

	// Angka 2 merupakan nilai jumlah buffer-nya, 2 merupakan angka index slot buffer yg terbesar
	// Yang berarti jumlahnya ada 3 karena dimulai dari index 0 (0,1,2)
	messages := make(chan int,2)

	go func() {
		for{
			i:=<-messages
			fmt.Println("Receive data :",i)
		}
	}()

	for i:=0;i<5;i++ {
		fmt.Println("Send data :",i)
		messages <- i
	}
}
