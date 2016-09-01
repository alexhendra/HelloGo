package main

import "fmt"

func main() {
	var angka = 4
	fmt.Println("Before :", angka)

	change(&angka,10)
	fmt.Println("After :", angka)
}

func change(original *int,value int) {
	*original = value
}