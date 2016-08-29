package main

import "fmt"

func main() {
	// var firstName string = "Alex"
	// var midName = "Hendra"
	// lastName := "..."

	// fmt.Println(firstName, midName, lastName)

	// Multi deklarasi
	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"

	// var umur, alamat, tgl_lahir = 23, "Jalan Raya", "08/08/1988"
	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// variable underscore _
	// Digunakan untuk menampung nilai yang tidak pernah dipakai (ibarat blackhole)
	// Biasanya	underscore sering dimanfaatkan untuk menampung nilai balik fungsi yang tidak digunakan.
	// yourName, _ := "Alex Hendra", "Blablabla"

	// Deklarasi dengan keyword new
	// Digunakan untuk mencetak data pointer dengan tipe data tertentu
	myName := new(string)
	fmt.Println(myName)  // Akan menampilkan alamat memory nilai myName
	fmt.Println(*myName) // Menampilkan isi / nilai dari myName

	// Variable konstanta
	const pi = 3.142857143
	fmt.Println(pi)
}
