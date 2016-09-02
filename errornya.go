package main

import (
	"fmt"
	"strconv"
)

// Error merupakan sebuah tipe. Error memiliki beberapa property yg menampung informasi yg berhubungan dengan Error yg bersangkutan
func main() {
	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	var number int
	var err error

	// Fungsi strconv.Atoi() digunakan untuk mengkonversi string menjadi int, nilai kembalian ada 2
	// nilai kembalian bertipe int & error
	number,err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	}else{
		fmt.Println(input,"is not number")

		// err.Error() menampilkan pesan error-nya
		fmt.Println(err.Error())
	}
}
