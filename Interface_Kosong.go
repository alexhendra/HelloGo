package main

import (
	"fmt"
	"strings"
)

func main() {
	// Interface kosong bisa menampung segala jenis data (array, pointer, dll)
	// Dengan cara penulisan seperti ini maka interface kosong dianggap sebagai tipe data
	var secret interface{}

	// variable map dengan key = string, dan value = interface kosong
	var data map[string]interface{}
	data = map[string]interface{}{
		"name": "Budi Gunawan",
		"grade": 2,
		"breakfast": []string{"Bulberry", "Chocolate", "Orange"},
	}
	// =========================================

	secret = "Alex Hendra"
	fmt.Println(secret)

	fmt.Println()

	secret = []string{"Apple", "Berry", "Chery"}
	fmt.Println(secret)

	fmt.Println()

	secret = 12.4
	fmt.Println(secret)

	fmt.Println()
	fmt.Println(data)
	fmt.Println(data["breakfast"])

	fmt.Println()

	// ====  Melakukan casting variable interface  ====

	var secret2 interface{}
	secret2 = 2

	// casting ke int
	var number = secret2.(int) * 10
	fmt.Println(secret2, "multiplied by 10 is :", number)

	fmt.Println()

	secret = []string{"apple", "manggo", "banana"}

	// casting ke slice string
	var gruits = strings.Join(secret.([]string), ",	")
	fmt.Println(gruits, "is	my favorite fruits")
	// ==============================================

	fmt.Println()

	// ====  Casting variable interface kosong ke Objek Pointer  ====

	// variable secret dideklarasikan bertipe interface kosong yang menampung alamat memory (referensi) objek struct "person"
	var secret3 interface{} = &person{name:"Sugeng", age:32}

	// Casting dilakukan dengan menuliskan tanda asteriks "*", untuk mendapatkan nilai aslinya
	var name = secret3.(*person).name
	fmt.Println(name)

	// ========================================
}

type person struct{
	name string
	age int
}
