package main

import "fmt"

func main() {

	// cara 1
	i:=1
	for i <= 3 {
		fmt.Println("Cara 1:", i)
		i++
	}

	// cara 2	
	for j:=7; j<=9; j++{
		fmt.Println("Cara 2:", j)
	}

	// cara 3	
	for {
		fmt.Println("loop")
		break
	}
}
