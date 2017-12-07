package main

import (
	"fmt"
	"strings"
)

// *** Function sebagai parameter ***

// *** cara 1 ***
func filterAgain(data []string, callback func(string) bool) []string{
	var result []string
	for _,each := range data{
		if filtered := callback(each); filtered{
			result=append(result, each)
		}
	}
	return result
}

// *** cara 2 ***
// Membuatkan alias dari function yang sebagai parameter tersebut, dengan membuat alias,
// menjadikan skema function tersebut menjadi tipe data baru
// Menggunakan statement "type"
type FilterCallback func(string) bool

func filterOne(data []string, callback FilterCallback) string{
	result:=""
	for _,item:= range data {
		if filtered:=callback(item); filtered{
			result=item
			break
		}
	}
	return result
}


func main() {
	data:=[]string {"wick","jason","ethan"}

	dataContainsO:=filterAgain(data,func(each string)bool {
		return strings.Contains(each,"o")
	})

	var dataLength5=filterAgain(data, func(each string) bool{
		return len(each) == 5
	})

	dataFilterOne:=filterOne(data,func(each string) bool{
		return strings.Contains(each, "o")
	})

	fmt.Println("Data asli \t\t:", data)
	fmt.Println("Filter ada huruf \"o\"\t:", dataContainsO)
	fmt.Println("Filter jumlah huruf sama dengan\"5\"\t:", dataLength5)
	fmt.Println("Filter ada huruf \"o\" 1 item saja\t:", dataFilterOne)
}
