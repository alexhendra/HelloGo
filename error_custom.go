package main

import (
	"strings"
	"errors"
	"fmt"
)


// Untuk membuat custom error dapat menggunakan fungsi errors.New


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
		fmt.Println(err.Error())
	}
}
