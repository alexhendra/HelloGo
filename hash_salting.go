package main

import (
	"fmt"
	"time"
	"crypto/sha1"
)

/*
	Salt dalam konteks kriptografi adalah data random yg digabungkan pada data asli sebelum proses hash dilakukan
	Hash merupakan enkripsi satu arah dgn lebar data yg sudah pasti, menjadikan sangat mungkin sekali kalau hasil hash -
	untuk beberapa data adalah sama.
	Salt sangat berguna utk mencegah serangan menggunakan metode pencocokan data-data yg hasil hash-nya adalah sama (dictionary attach)
*/

func doHashUsingSalt(text string) (string, string) {

	// Fungsi UnixNano() mengembalikan nano second atau nano detik dari waktu sekarang
	var salt = fmt.Sprint("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("Text : '%s', Salt : %s", text, salt)
	fmt.Println(saltedText)

	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func main() {
	var text = "Pesan rahasia saya"
	fmt.Printf("original : %s\n\n", text)
	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("hashed 1: %s\n\n", hashed1)

	var hashed2, salt2 = doHashUsingSalt(text)
	fmt.Printf("hashed 2: %s\n\n", hashed2)

	var hashed3, salt3 = doHashUsingSalt(text)
	fmt.Printf("hashed 3: %s\n\n", hashed3)

	//a, b, c := salt1, salt2, salt3

	fmt.Printf("a : %s, b : %s, c : %s\n",salt1,salt2,salt3)
}
