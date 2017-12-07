package main

import "fmt"

// Struct adalah kumpulan definisi variable (atau property) dan atau fungsi(atau method), yang dibungkus dengan nama tertentu.
// sama dengan konsep class pada OOP, Struct bisa menjadi variable objek (di instans menjadi objek)

// Cara vertikal
type mahasiswa struct {
	name string
	grade string
}

// Cara horizontal
type orang struct {nama string; umur int; hobbies []string}


func main() {
	// ==== Inisialisasi Objek ====
	var std1 = mahasiswa{}
	std1.name="Budi"
	std1.grade="3"

	var std2 = mahasiswa{"Charlie","4"}

	var std3 = mahasiswa{name:"Dadang"}


	var mhs1 mahasiswa
	mhs1.grade="2"
	mhs1.name="Alex"

	fmt.Println("Name :",mhs1.name)
	fmt.Println("Grade :",mhs1.grade)

	fmt.Println("Name :",std1.name)
	fmt.Println("Grade :",std1.grade)

	fmt.Println("Name :",std2.name)
	fmt.Println("Grade :",std2.grade)

	fmt.Println("Name :",std3.name)
	fmt.Println("Grade :",std3.grade)

	fmt.Println()

	// Objek dari struct bisa diambil nilai Pointer-nya, dan disimpan pada variable yang bertipe struct pointer
	var mhs2 *mahasiswa = &mhs1
	fmt.Println("Mahasiswa 1 :", mhs1.name)
	fmt.Println("Mahasiswa 2 :", mhs2.name)

	mhs2.name = "JOKO"
	fmt.Println("Mahasiswa 1 :", mhs1.name)
	fmt.Println("Mahasiswa 2 :", mhs2.name)

	// ==================================

	// ==== Deklarasi dan Inisialisasi struct secara horizontal ====
	var org1 = struct {nama string; umur int}{nama:"Alex", umur:28}
	var org2 = struct {nama string; umur int}{nama:"Hamham", umur:35}

	fmt.Println()
	fmt.Println("Org1: " ,org1)
	fmt.Println("Org2: " ,org2)
	// ============================================
}

