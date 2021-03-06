package main

import (
	"strings"
	"fmt"
)

func main() {
	var data = []string{"Sule","Aziz","Andre","Nunung"}
	var dataContains = filter(data,func(sumberString string) bool {
		return strings.Contains(sumberString,"u")
	})

	var dataLength = filter(data,func(sumberString string) bool {
		return len(sumberString) == 5
	})

	fmt.Println("Data asli \t\t:",data)

	fmt.Println("Filter ada huruf \"u\"\t:", dataContains)
	fmt.Println("Filter jumlah huruf \"5\"\t:", dataLength)
}

/*func filter(data []string, callback func(string) bool) []string{
	var result []string
	for _, valuenya := range data {
		if filtered := callback(valuenya); filtered {
			result = append(result, valuenya)
		}
	}

	return result
}*/

// Bisa juga parameter yang berupa closure di fungsi filter dibuatkan alias
// Dengan alias maka skema fungsi tersebut menjadi tipe data baru
// Pembuatan alias dengan keyword "type"
type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string{
	var result []string
	for _, valuenya := range data {
		if filtered := callback(valuenya); filtered {
			result = append(result, valuenya)
		}
	}

	return result
}

