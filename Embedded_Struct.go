package main

import "fmt"

// Embedded struct adalah penurunan properti dari satu struct ke struct lain,
// sehingga properti struct yang diturunkan bisa digunakan

type orang struct {
	nama string
	umur int
}

type mahasiswa struct {
	grade int

	// cara embedded-nya cukup tuliskan nama struct yg mau di embedded
	orang
}

func main() {
	var mhs1 mahasiswa
	mhs1.nama = "Alex"
	mhs1.umur = 30
	mhs1.grade = 8

	fmt.Println("name :", mhs1.nama)
	fmt.Println("age :", mhs1.umur)
	fmt.Println("age :", mhs1.orang.umur)
	fmt.Println("grade :", mhs1.grade)

	fmt.Println()

	var org1 = orang{nama: "Joko", umur: 21}
	var mhs2 = mahasiswa{orang: org1, grade: 2}

	fmt.Println("name :", mhs2.nama)
	fmt.Println("age :", mhs2.umur)
	fmt.Println("grade :", mhs2.grade)
}
