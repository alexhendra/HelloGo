package main

import "fmt"

func main() {
	// Fungsi closure, adalah fungsi yang disimpan ke dalam variable atau bisa juga disebut anonymous function
	var getMinMax = func(n []int) (int, int) {
		var min, max int

		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 8, 5, 1, 6, 7, 4}
	var min, max = getMinMax(numbers)

	// template "%v" digunakan untuk menampilkan segala jenis data
	fmt.Printf("data:%v \n min:%v \n max:%v \n", numbers, min, max)
}
