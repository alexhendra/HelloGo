package main

import "fmt"

func main() {
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%1.f%s not bad\n", percent, "%")
	}

	if 7%2 == 0{
		fmt.Println("7 is even")
	}else{
		fmt.Println("7 is odd")
	}
}
