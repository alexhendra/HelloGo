package main

import "fmt"

func main() {
	var datas = [2]string{"Alex", "Hendra"}

	// Nilai tiap elemen array "datas" ditampung variable "data"
	// Sedangkan indeks-nya ditampung oleh variable "i"
	/*for i, data := range datas {
		fmt.Printf("elemen %d : %s\n", i, data)
	}*/

	// Apabila indeks-nya tidak ingin ditampilkan, maka gunakan variable underscore _
	for _, data := range datas {
		fmt.Printf("elemen : %s\n", data)
	}
}
