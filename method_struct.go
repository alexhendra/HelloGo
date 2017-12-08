package main

import (
	"fmt"
	"strings"
)

// Method adalah bagian dari struct
// Dimana method tersebut dapat diakses melalui instants objek

type orang struct {
	nama string
	umur int
}

//Fungsi "sayHai" adalah milik struct "orang" dilihat dari "func(nama_variable nama_struct)"
func (o orang) sayHai()  {
	fmt.Println("Hai :", o.nama)
}

func (s orang) getNameAt(i int) string {
	return strings.Split(s.nama," ")[i-1]
}

// == Method Pointer ==
func (o *orang) sayHaiPointer()  {
	fmt.Println("Hai :", o.nama)
}
// ==============

func main() {
	var org1 = orang{"Lusi Bongot",24}
	org1.sayHai()

	var nama = org1.getNameAt(2)
	fmt.Println("Nama Panggilan",nama)

	fmt.Println()

	// variable objek pointer
	var org2 = &orang{"Bambang Rio",24}
	org2.sayHaiPointer()

	var nama2 = org2.getNameAt(2)
	fmt.Println("Nama Panggilan",nama2)

	fmt.Println()

}
