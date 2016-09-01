package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string)  {
	for i:=0; i<till;i++ {
		fmt.Println((i+1), message)
	}
}	

// GOROUTINE = mini thread, sangat ringan hanya butuh 2kB memori saja untuk 1 GOROUTINE.
// Eksekusinya bersifat asynchronous
func main() {
	// Code ini digunakan untuk menentukan jumlah prosesor yg aktif
	// Apabila angka prosesor yang diinputkan lebih dari jumlah asli logical processor,
	// maka dianggap menggunakan semua prosesor yang ada
	runtime.GOMAXPROCS(2)

	// Code ini adalah GOROUTINE, ditandai dengan kata "go"
	go print(5,"GoRoutine")


	print(5,"Print Biasa")


	var input string
	// Digunakan utk menerima input keyboard
	fmt.Scanln(&input)
}
