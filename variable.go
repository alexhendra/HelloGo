package main

import "fmt"

func main() {
	// cara deklarasi variable di go
	var a string="Alex"
	fmt.Println(a)

	var b = 3
	fmt.Println(b)

	c:=true
	fmt.Println(c)

	// deklarasi multi variable di 1 baris code
	var d,e string="Budi", "Joko"
	fmt.Println(d,e)

	var f,g int=8,9
	fmt.Println(f,g)

	h,i:=false,true
	fmt.Println(h,i)
}
