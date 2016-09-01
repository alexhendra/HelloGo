package main

import "fmt"

func main() {
	// Pointer adalah referensi atau alamat memory

	// variable nilaiA adalah variable biasa
	var nilaiA int = 4

	// var nilaiB adalah variable pointer
	// yang ditandai dengan asteriks "*" sebelum penulisan tipe data
	// dan diisikan dengan nilai dari alamat memory dari variable nilaiA
	var nilaiB *int = &nilaiA

	fmt.Println("Nilai A (value):", nilaiA)
	fmt.Println("Nilai A (address):", &nilaiA)

	fmt.Println("Nilai B (value):", *nilaiB)
	fmt.Println("Nilai B (address):", nilaiB)

	fmt.Println()

	nilaiA = 5

	fmt.Println("Nilai B (value):", *nilaiB)
	fmt.Println("Nilai B (address):", nilaiB)
}
