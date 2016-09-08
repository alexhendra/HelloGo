package main

import (
	"encoding/json"
	"fmt"
)

type User2 struct {
	// `json:"Name"` adalah tag, digunakan utk mapping data JSON ke property "FullName"
	// Property "FullName" ditugaskan utk menampung data JSON property "Name"
	FullName string `json:"Name"`

	Age int
}


func main() {
	var object = []User2{{"john wick", 27}, {"ethan hunt", 32}}

	// Fungsi json.Marshal() digunakan utk decoding data ke JSON.
	// Data tersebut bisa berupa variable objek cetakan struct, map[string]interface{}, bisa jga bertipe array
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
	}
	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
