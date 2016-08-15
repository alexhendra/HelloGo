package main

import "fmt"

func main() {
	// Switch pada golang, berbeda dengan bahasa pemgrograman lain
	// Apabila kondisi telah terpenuhi, maka tdk akan dilanjutkan ke pengecekan case selanjutnya

	var point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Pemanfaatan 1 case pada switch bisa memiliki lebih dari 1 kondisi
	var point2 = 8
	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")

	}

}
