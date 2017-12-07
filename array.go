package main

import "fmt"

func main() {
	// array adalah kumpulan data bertipe sama

	// *** Cara 1 ***
	var data [4]int
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	for i := 0; i < len(data); i++ {
		fmt.Println("Data [", i, "] =", data[i])
	}

	// *** Cara 2 (horizontal) ***
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	 fmt.Println("Jumlah	element	\t\t", len(fruits))
	 fmt.Println("Isi semua element \t", fruits)

	// cara	vertikal
	/* var fruits = [4]string{
	 	"apple",
	 	"grape",
	 	"banana",
	 	"melon",
	 }*/

	// fmt.Println("Jumlah	element	\t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)

	// *** Cara 3 ***
	// Deklarasi array boleh tidak dituliskan jumlah item array-nya
	// Jumlah item akan dikalkulasi otomatis menyesuaikan data elemen yang diisikan
	var bilangan = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("isi dari array :", bilangan)

	// *** Cara 4 ***
	// Membuat dan menentukkan jumlah elemen array dengan statement "make"
	// make([]tipedatanya,jumlahelemennya)
	var arrayku = make([]string,2)
	arrayku[0] = "alex"
	arrayku[1] = "hendra"
	fmt.Println(arrayku)


	// *** Array 2 dimensi ***
	arr2:=[2][3] int {}
	for i:=0; i<2; i++ {
		for j:=0; j<3; j++ {
			arr2[i][j]=i+j
		}
	}

	fmt.Println("2d: ", arr2)
}
