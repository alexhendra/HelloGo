package main

import "fmt"

func main() {
	// Map adalah tipe data asosiatif yang ada di golang
	// Bentuknya key-value, artinya setiap datanya terdiri dari key dan value
	// contoh : var barang map[tipe_data_key]tipe_data_value
	//var barang map[string]int
	//barang = map[string]int{}

	/*var barang = map[string]int{}

	barang["Ban"]=1000
	barang["Oli"]=2000
	barang["Gardang"]=3000

	fmt.Println("Ban : ", barang["Ban"])
	fmt.Println("Oli : ", barang["Oli"])
	fmt.Println("Gardang : ", barang["Gardang"])*/

	// Pengaksesan data map yang tidak ada di dalam map atau mengakses data yang value-nya tidak ada di map -
	// Akan mengembalikan nilai 0, karena nilai default map = 0
	//fmt.Println("Kopling : ", barang["Kopling"])

	// ====== Inisialisasi Nilai MAP =====

	// cara vertikal
	var barang = map[string]int{"Ban":1000,"Oli":2000,"Gardang":3000}

	// cara horizontal
	// di akhir item harus dituliskan tanda koma ","
	var mahasiswa = map[string]string{
		"Budi":"8888",
		"Tono":"9999",
	}

	fmt.Println("Barang: ", barang)
	fmt.Println("Mahasiswa: ", mahasiswa)

	// ==== Dengan make() atau new() ====
	//var gudang = new(map[string]int)
	//
	//var toko = make(map[string]int)
	//
	//fmt.Println(gudang)
	//fmt.Println(toko)
	// ===============================


	// ===== Iterasi dengan for-range =====

	for key, val := range barang {
		fmt.Println(key,"\t:", val)
	}

	fmt.Println()

	// ===== Menghapus item Map =====
	// Fungsi delete() digunakan untuk menghapus item dengan key tertentu pada map

	// Menghapus elemen "Gardang" pada map barang
	delete(barang,"Gardang")

	for key, val := range barang {
		fmt.Println(key,"\t:", val)
	}
	// =============================

	// ==== Menghitung jumlah item map ====
	fmt.Println("Jumlah elemen Barang:", len(barang))
	fmt.Println()
	// ================================

	// ==== Mengecek apakah elemen dengan key tertentu ada atau tidak =====
	var value,isExists = barang["Gardang"]

	if isExists {
		fmt.Println(value)
	} else {
		fmt.Println("Item is not exists !")
	}
	// ==============================================

}
