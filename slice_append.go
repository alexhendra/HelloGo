package main

import "fmt"

func main() {
	// fungsi append() digunakan untuk menambahkan elemen pada slice
	// Apabila nilai len() = cap() maka bila elemen baru yang ditambahkan akan menjadi referensi baru
	// Namun apabila nilai len() < cap() maka elemen baru akan ditambahkan ke slice dan slice lain yang referensi-nya sama

	var mahasiswa = []string{"Alex", "Budi", "Coky", "Dedi", "Ema", "Fino"}
	var newMahasiswa = mahasiswa[0:2]

	fmt.Println(cap(newMahasiswa))
	fmt.Println(len(newMahasiswa))
	fmt.Println()
	fmt.Println(mahasiswa)
	fmt.Println(newMahasiswa)

	// mhsB merupakan slice baru dan jadi referensi baru
	// sehingga slice lama ikut berubah elemennya. Elemen index ke-2 diubah menjadi "Gomgom"
	// Dan juga karena len() < cap(), maka elemen yang baru ditambahkan terdeteksi sebagai -
	// perubahan nilai pada referensi lama
	var mhsB = append(newMahasiswa, "Gomgom")

	fmt.Println(mahasiswa)
	fmt.Println(newMahasiswa)
	fmt.Println(mhsB)


}
