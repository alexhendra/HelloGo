package main

import "fmt"

// Defer digunakan untuk menyatakan bahwa baris program itu dijalankan diakhir setelah seluruh baris kode dijalankan
// Ketika ada banyak statement yg di defer, maka statement tersebut akan dieksekusi diakhir secara berurutan
func main() {
	// Baris kode ini dijalankan terakhir, setelah baris kode dibawahnya dijalankan
	defer fmt.Println("Haiii.....")
	fmt.Println("Good Luck!!")
}
