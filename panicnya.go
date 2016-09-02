package main

import (
	"strings"
	"errors"
	"fmt"
)

// Panic digunakan untuk menampilkan stack trace dari error. Informasi yg ditampilkan lebih detail.


func validate(input string)(bool,error)  {
	if strings.TrimSpace(input) == "" {
		// mengisi pesan error buatan sendiri
		return false, errors.New("Cannot be empty")
	}
	return true,nil
}

func main() {
	var name string
	fmt.Print("Type your name :")
	fmt.Scanln(&name)

	if valid,err := validate(name); valid {
		fmt.Println("Halo :",name)
	}else{
		panic(err.Error())
	}
}
