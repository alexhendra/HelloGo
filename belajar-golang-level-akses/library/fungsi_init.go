package library

import "fmt"

var Student= struct {
	Name string
	Grade int
}{}


// Fungsi init() akan dijalankan pertama kali ketika package library dipanggil dari code lain
// 1 package bisa memiliki banyak fungsi init()
// Urutan eksekusinya adalah dimulai dari file mana yg terlebih dulu dipanggil
func init() {
	Student.Name = "John Wick"
	Student.Grade = 2

	fmt.Println("-->  library/fungsi_init.go imported")
}
