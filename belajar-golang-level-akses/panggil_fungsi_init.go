package main

import (
	"fmt"
	"belajargolang/belajar-golang-level-akses/library"
)

func main() {
	// Ketika file ini dijalankan maka fungsi init yang ada di package library dijalankan pertama kali
	fmt.Printf("Name : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)
}
