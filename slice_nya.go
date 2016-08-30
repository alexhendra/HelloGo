package main

import "fmt"

func main() {
	// Slice adalah referensi elemen array

	// Cara inisialisasi slice
	// Mirip dengan array, bedanya slice tidak ada tanda "..." di bagian []
	/*var dataSlice = []string{"Andi","Budi","Sule"}

	fmt.Println(dataSlice)*/

	//var arrayku = [...]string{"Pepaya","Anggur","Melon","Jeruk","Durian"}

	// Karena slice merupakan referensi elemen array, maka datanya bisa bersumber dari array
	// kode arrayku[0:3] maksudnya adalah mengambil item array dari index 0 sampai < 3
	// dimana kembaliannya berupa slice atau elemen referensi
	/*var newSliceData = arrayku[0:3]
	fmt.Print("Data Slice baru : ", newSliceData)*/

	// Slice merupakan tipe data reference
	// Artinya jika ada slice baru yang terbentuk dari slice lama,
	// maka elemen slice baru memiliki referensi yang sama dengan slice lama.
	//var dataSlice = []string{"Andi", "Budi", "Coki", "Dedi", "Eko", "Fandil", "Gonju", "Hurje"}
	//var mhsA = dataSlice[0:3]
	//var mhsB = dataSlice[2:5]
	//
	//var detailMhsA = mhsA[0:1]
	//var detailMhsB = mhsB[3:4]
	//
	//fmt.Println(dataSlice)
	//fmt.Println(mhsA)
	//fmt.Println(mhsB)
	//fmt.Println(detailMhsA)
	//fmt.Println(detailMhsB)

	// sebelumnya detailMhsA[0] = "Andi", diubah menjadi "Zulham"
	//detailMhsA[0] = "Zulham"
	//
	//fmt.Println()
	//fmt.Println(dataSlice)
	//fmt.Println(mhsA)
	//fmt.Println(mhsB)
	//fmt.Println(detailMhsA)
	//fmt.Println(detailMhsB)

	// setelah code di-run, maka perubahan terhadap detailMhsA[0] berpengaruh terhadap slice lama

	// Fungsi len() digunakan untuk menghitung lebar slice yang ada
	//fmt.Println("Lebar dataSlice : ", len(dataSlice))

	// Fungsi cap() digunakan untuk menghitung lebar maksimum/kapasitas slice.
	// Nilai kembalian fungsi ini awalnya sama dengan len(),
	// tapi bisa berubah tergantung dari operasi slice yang dilakukan
	//fmt.Println("Contoh penggunaan cap()")
	//fmt.Println("len(dataSlice) : ", len(dataSlice))
	//fmt.Println("cap(dataSlice) : ", cap(dataSlice))
	//
	//fmt.Println()
	//
	//fmt.Println("len(mhsA) : ", len(mhsA))
	//fmt.Println("cap(mhsA) : ", cap(mhsA))
	//
	//fmt.Println()
	//
	//fmt.Println("len(mhsB) : ", len(mhsB))
	//fmt.Println("cap(mhsB) : ", cap(mhsB))

	// apabila pengambilan nilai array dimulai dari index > 0
	// maka total kapasitas slice dimulai dari index tersebut sampai index terakhir dari array

	// ====== cara melakukan slice (slicing) dengan langsung menentukan kapasitasnya ======
	// mahasiswa[index_startnya,index_batas_atasnya,jumlah_kapasitasnya]

	var mahasiswa = []string {"Anto","Bambang","Erlang","Duko","Fanta"}
	var mhs_A = mahasiswa[1:4:4]

	fmt.Println(mahasiswa)
	fmt.Println(mhs_A)

	fmt.Println()

	fmt.Println(len(mhs_A))
	fmt.Println(cap(mhs_A))
	// ====================================================================

}
