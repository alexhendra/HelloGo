package main

import (
	"flag"
	"fmt"
)

/*
	Cara ke 2 untuk menggunakan flag adalah dgn passing reference, dimana nilai dari flag diambil lewat parameter pointer
*/

func main() {
	var name string
	// flag.StringVar() melakukan penyimpanan dgn metode referensi, dimana variabel yg digunakan adalah "name"
	flag.StringVar(&name,"name","any","type your name")
	var age = flag.Int64("age",25,"type your age")

	flag.Parse()

	// Pengaksesan nilai asli tdk perlu di-deference lagi
	fmt.Printf("Name\t: %s\n",name)

	fmt.Printf("Age\t: %d\n",*age)

	// Cara eksekusinya : go run flagnya.go -name="Alex Hendra" -age=28
}
