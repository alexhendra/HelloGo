package main

import "fmt"

func main() {

	// Penggunaan label, label yang dibuat adalah outerLoop
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop //Target yang akan di break, adalah outerloop
			}
			fmt.Print("matriks	[", i, "][", j, "]", "\n")
		}
	}
}
