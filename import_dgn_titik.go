package main

// Import dgn tanda titik "." membuat method/property/fungsi milik package yg diimport tdk perlu dituliskan nama packagenya
import (
	. "belajargolang/belajar-golang-level-akses/library"
	"fmt"
)

func main() {
	// Tdk perlu dituliskan nama package-nya lagi : library.Mahasiswa{}
	var mhs = Mahasiswa{"Yuya",30}
	fmt.Println("Nama :",mhs.Name)
	fmt.Println("Grade :",mhs.Grade)

}
