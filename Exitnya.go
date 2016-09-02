package main

import (
	"fmt"
	"os"
)

// os.Exit() digunakan untuk menghentikan program secara paksa.
// Semua statement setelah exit tidak akan diproses
// Fungsi ini memiliki sebuah parameter bertipe numerik yg wajib diisi.
// Angka yang dimasukkan akan muncul sebagai exit status ketika program berhenti


func main() {
	defer fmt.Println("Halo...")
	os.Exit(1)
	fmt.Println("Selamat Datang")
}
