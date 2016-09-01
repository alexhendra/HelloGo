package main

import (
	"belajargolang/belajar-golang-level-akses/library"
	"fmt"
)

func main() {
	library.SayHello()

	var mhs1 = library.Mahasiswa{"Gisella",1}
	fmt.Println(mhs1.Name)
}
