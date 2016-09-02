package main

import (
	"fmt"
	"runtime"
)

// Penerimaan data lewat channel yang dipakai oleh banya GOROUTINE, akan lebih mudah dengan memanfaatkan keyword "for-range"
// For-range apabila diterapkan di channel, akan melakukan perulangan tanpa henti.
// Perulangan tersebut tetap berjalan meski tidak ada transaksi pada channel, dan hanya akan berhenti jika status channel -
// berubah menjadi closed atau sudah ditutup.
// Fungsi close() digunakan untuk menutup channel

// Channel yg sudah ditutup tidak bisa digunakan lagi untuk menerima ataupun mengirim data.

// ============ Arah panah Channel pada parameter =============
// -----------------------------------------------------------------------------------------
// | Sintaks		| Penjelasan							   |
// |---------------------------------------------------------------------------------------|
// | ch chan string	| Parameter "ch" bisa digunakan untuk mengirim dan menerima data   |
// |---------------------------------------------------------------------------------------|
// | ch chan <- string	| Parameter "ch" hanya bisa digunakan untuk mengirim data	   |
// |---------------------------------------------------------------------------------------|
// | ch <- chan string	| Parameter "ch" hanya bisa digunakan untuk menerima data	   |
// |----------------------------------------------------------------------------------------

// Fungsi untuk mengirim data via channel
func sendPesan(ch chan <- string)  {
	for i:=0; i<20;i++ {
		ch <- fmt.Sprintf("Data %d",i)
	}

	// setelah semua data sudah dikirimkan maka channel di-close
	close(ch)
}

// Fungsi untuk menerima data via channel
func printPesan(ch <- chan string)  {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	go sendPesan(messages)

	printPesan(messages)
}
