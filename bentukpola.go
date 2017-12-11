package main

import (
	"fmt"

)

/*
  x
   x
x o x
 x
  x

 */

func bentukPola(i int) {
	total:=i*2
	for x:=1;x<=total;x++ {
		for j := 1; j <= i; j++ {
			for k := 1; k < i; k++ {
				fmt.Printf(" ")
			}
			fmt.Printf("x")
			//i++
		}
		fmt.Printf("\n")
	}
}

func main() {
	bentukPola(3)
}
