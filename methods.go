package main

import (
	"fmt"
	"strings"
)

// di Golang, method itu menempel pada struct, sehingga hanya bisa diakses lewat instance object dari struct
type hewan struct{
	jeniskelamin string
	name string
}

// method milik struct hewan, karena tertulis "s hewan"
func(s hewan) sayHello() {
	fmt.Println("halo ", s.name)
}

func(s hewan) getNameAt(i int) string{
	return strings.Split(s.name, " ")[i-1]
}

func(s hewan) changeName1(n string) {
	s.name=n
}

// *** Method Pointer ***
// karena bersifat reference ke memory, maka apabila nilai property name pada struct ikut berubah
func(s *hewan) changeName2(n string) {
	s.name=n
}


func main() {
	var h1 = hewan{"Jantan", "Gajah Sumatra"}
	h1.sayHello()

	name:=h1.getNameAt(1)
	fmt.Println("Nama: ", name)

	h1.changeName1("Harimau")
	fmt.Println("hasil dari method changeName1:", h1.name)

	h1.changeName2("Singa")
	fmt.Println("hasil dari method changeName2:", h1.name)

	// pengaksesan method pointer bisa dari instance object biasa atau instance object pointer
	var h2=hewan{"Betina", "Tupai"}
	h2.changeName2("Lutung")
	fmt.Println("h2:", h2.name)

	var h3= &hewan{"Betina", "Unta"}
	h3.changeName2("Gorila")
	fmt.Println("h3:", h3.name)

}
