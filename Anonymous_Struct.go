package main

import "fmt"

type orang struct {
	nama string
	umur int
}

func main() {
	// === Anonymous Struct ===
	// Struct dideklarasi di awal dan dijadikan objek diawal
	// Cocok untuk variable objek struct yang dipakai hanya sekali
	var mhs1 = struct {
		orang
		grade int
	}{}

	// Anonymous struct dengan inisialisasi
	var mhs2 = struct {
		orang
		grade int
	}{
		orang:orang{"Bambang",35},
		grade:4,
	}

	// Anonymous struct dijadikan sebagai array
	var allMhs = []struct{
		orang
		grade int
	}{
		{orang: orang{"Andi",20}, grade:1},
		{orang: orang{"Fadil",25}, grade:3},
		{orang: orang{"Suminto",40}, grade:2},
	}

	for _,mhs := range allMhs {
		fmt.Println(mhs)
	}

	// ========================

	// ==== Anonymous Struct yang tidak bisa diinisialisasi pada saat deklarasi ====
	var dosen struct{
		grade int
	}

	fmt.Println()
	fmt.Println("Dosen :",dosen)
	// =======================================


	fmt.Println()

	mhs1.orang = orang{"Alex",30}
	mhs1.grade = 2

	fmt.Println("name :", mhs1.orang.nama)
	fmt.Println("age :", mhs1.orang.umur)
	fmt.Println("grade :", mhs1.grade)

	fmt.Println()

	fmt.Println("name :", mhs2.orang.nama)
	fmt.Println("age :", mhs2.orang.umur)
	fmt.Println("grade :", mhs2.grade)
	fmt.Println(mhs2)

}
