package main

import (
	"regexp"
	"fmt"
)

// Regex (Regular Expression) adalah suatu teknik yang digunakan untuk pencocokan string dgn pola tertentu

func main() {
	var text = "banana burger soup"

	// Fungsi regexp.Compile() digunakan untuk mengkompilasi ekspresi regex yg dimasukkan.
	// Fungsi tersebut mengembalikan objek bertipe regexp.*Regexp
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Fungsi FindAllString() berfungsi untuk pencarian semua string yg sesuai dgn ekspresi regex
	// Dengan kembalian berupa array string
	// Nilai kembalian bisa ditentukan, contoh dibawah mengembalikan 2 data
	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)


	// Jika nilai kembalian di set -1, maka akan mengembalikan semua data
	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)

	// Fungsi MatchString() digunakan utk mengecek apakah string memenuhi sebuah pola regexp
	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)

	// Fungsi FindString() digunakan utk mencari string yg memenuhi kriteria regexp yg telah ditentukan
	// Fungsi tersebut hanya mengembalikan 1 buah hasil saja
	var str = regex.FindString(text)
	fmt.Println(str)

	// Fungsi FindStringIndex(), method ini sama dengan FindString() hanya saja yg dikembalikan index saja
	var idx = regex.FindStringIndex(text)
	fmt.Printf("%v\n", idx)

	// Fungsi regex.ReplaceAllString() berguna untuk me-replace semua string yg
	// memenuhi kriteria regexp dgn string lain
	var strRep = regex.ReplaceAllString(text, "potato")
	fmt.Println(strRep)

	// Fungsi regex.Split() digunakan utk memisah string dgn pemisah adalah substring yg
	// memenuhi kriteria regexp yg telah ditentukan
	var strSplit = regex.Split(text, -1)
	fmt.Printf("%#v	\n", strSplit)
}
