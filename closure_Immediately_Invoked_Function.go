package main

import "fmt"

func main() {
	// closure jenis Immediately-Invoked Function Expression (IIFE)
	// dieksekusi langsung pada saat deklarasi

	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	// Yang ditampung adalah variable newNumbers adalah nilai kembaliannya
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3) // --> Ini adalah parameternya
	fmt.Println("original number :", numbers)
	fmt.Println("filtered number :", newNumbers)

	// Cara penulisan gaya manifest typing
	/*var closure (func(string, int, []string) int)
	closure = func(a string,b int,c []string) int {
		return 1
	}

	fmt.Println(closure("a",10,[]string{"1","2","3","4"}))*/

}
