package main

import (
	"flag"
	"fmt"
)

/*
	Flag memiliki kegunaan yg sama seperti arguments, yaitu utk parameterize eksekusi program,
	dengan penulisan dalam bentuk key-value
*/

func main() {
	// Disiapkan flag bertipe string, key adalah "name", dgn nilai default "any" dan keterangan "type your name"
	// Nilai balik dari flag.String() adalah string pointer, jadi perlu di-deference terlebih dahulu agar dpt nilai asli
	var name =  flag.String("name","any","type your name")
	var age = flag.Int64("age",25,"type your age")

	flag.Parse()
	// Nilai balik dari flag.String() adalah string pointer, jadi perlu di-deference terlebih
	// dahulu agar dpt nilai asli (*name)
	fmt.Printf("Name\t: %s\n",*name)
	fmt.Printf("Age\t: %d\n",*age)

	// Cara eksekusinya : go run flagnya.go -name="Alex Hendra" -age=28
}
