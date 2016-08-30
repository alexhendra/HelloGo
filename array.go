package main

import "fmt"

func main() {
	// array adalah kumpulan data bertipe sama
	/*var data [4]int
	 for i := 0; i < len(data); i++ {
	 	data[i] = i
	 }

	 for i := 0; i < len(data); i++ {
	 	fmt.Println("Data [", i, "] =", i)
	 }*/

	// cara horizontal
	/* var fruits = [4]string{"apple", "grape", "banana", "melon"}
	 fmt.Println("Jumlah	element	\t\t", len(fruits))
	 fmt.Println("Isi semua element \t", fruits)*/

	// cara	vertikal
	/* var fruits = [4]string{
	 	"apple",
	 	"grape",
	 	"banana",
	 	"melon",
	 }*/

	// fmt.Println("Jumlah	element	\t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)

	// Deklarasi array boleh tidak dituliskan jumlah item array-nya
	// Jumlah item akan dikalkulasi otomatis menyesuaikan data elemen yang diisikan
	/*var data = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("isi dari array :", data)*/

	// Membuat dan menentukkan jumlah elemen array dengan "make"
	// make([]tipedatanya,jumlahelemennya)
	var arrayku = make([]string,2)
	arrayku[0] = "alex"
	arrayku[1] = "hendra"
	fmt.Println(arrayku)
}
