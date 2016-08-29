package main

import "fmt"

func main() {

	// cara 1
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// cara 2
	var i = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	// cara 3
	j := 0
	for {
		fmt.Println("Angka", j)
		j++
		if j == 5 {
			break
		}
	}
}
