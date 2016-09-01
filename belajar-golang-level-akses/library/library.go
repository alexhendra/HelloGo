package library

import "fmt"

// Struct yang access PUBLIC
type Mahasiswa struct {
	// Property PUBLIC
	Name string

	// Property PRIVATE
	//grade int

	Grade int
}

// fungsi SayHello() memiliki access PUBLIC ditandai dengan awalan huruf BESAR pada nama fungsi
func SayHello() {
	fmt.Println("Hello")
}

// fungsi introduce() memiliki access PRIVATE ditandai dengan awalan huruf KECIL pada nama fungsi
func introduce(name string) {
	fmt.Println("Nama saya :", name)
}
