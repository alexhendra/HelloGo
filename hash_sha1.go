package main

import (
	"crypto/sha1"
	"fmt"
)


/*
	Hash merupakan sebuah algoritma enkripsi untuk mengubah text menjadi deretan karakter acak.
	Jumlah karakter hasil hash selalu sama. Hash termasuk one-way encription,
	membuat hasil dari hash tidak bisa dikembalikan ke text asli.
	Sha1 atau Secure Hash Algorithm 1 merupakan salah satu algoritma hashing yang sering digunakan untuk enkripsi data.
	Hasil dari sha1 adalah data dengan lebar 20 byte atau 160 bit,
	biasa ditampilkan dalam bentuk bilangan heksadesimal 40 digit.
*/


func main() {
	var pesan = "Ini pesan rahasia"

	// Fungsi sha1.New() menghasilkan nilai hash.Hash
	var sha = sha1.New()

	// Fungsi Write() digunakan untuk men-set data yg akan di-Hash
	sha.Write([]byte(pesan))

	// Fungsi Sum() digunakan untuk eksekusi proses hash, menghasilkan data yg sudah di-hash dalam bentuk []byte
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x",encrypted)
	fmt.Println(encryptedString)
}
