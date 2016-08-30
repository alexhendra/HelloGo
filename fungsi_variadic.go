package main

import (
	"fmt"
	"strings"
)

func main() {
	//var avg = calculate2(2,3,4,5,6,7,8,9,10)

	// fmt.Sprintf() tidak menampilkan nilai, tetapi mengembalikan nilai
	//var msg = fmt.Sprintf("Rata-rata : %.2f",avg)
	//fmt.Println(msg)

	/*// Mengirimkan slice sebagai passing parameter ke fungsi variadic
	var numbers = []int{2,3,4,5,6,7,8,9,10}
	// cara mengirimkannya, nama variable diikuti tanda "..."
	var avg = calculate2(numbers...)

	var msg = fmt.Sprintf("Rata-rata : %.2f",avg)
	fmt.Println(msg)*/

	//yourHobbies("Alex","Football","Basket","Swimming")
	var myHobbies = []string{"Football","Basket","Swimming"}
	// mengirimkan data slice ke parameter variadic
	yourHobbies("Alex",myHobbies...)

}

// Fungsi variadic
// Dimana parameter nya tidak terbatas dan memiliki tipe data yang sama untuk setiap parameter
// ditandai dengan "..."
func calculate2(numbers ...int) float64 {
	var total int = 0
	for _,number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

// Fungsi dengan parameter biasa & variadic sekaligus
// syaratnya, parameter variadic harus diletakkan diposisi terakhir
func yourHobbies(name string,hobbies ...string) {
	var hobhiesString = strings.Join(hobbies,", ")
	fmt.Printf("Hello, my name is: %s\n",name)
	fmt.Printf("My hobbies are: %s\n",hobhiesString)
}