package main

import (
	"encoding/json"
	"fmt"
)

type Usernya struct {
	// `json:"Name"` adalah tag, digunakan utk mapping data JSON ke property "FullName"
	// Property "FullName" ditugaskan utk menampung data JSON property "Name"
	FullName string `json:"Name"`

	Age int
}

func main() {
	var jsonString = `[
		{"Name":"Alex Hendra", "Age": 28},
		{"Name":"Bambang Budi", "Age": 30}
	]`

	var datanya []*Usernya

	var err = json.Unmarshal([]byte(jsonString), &datanya)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("User 1:", datanya[0].FullName)
	fmt.Println("User 2:", datanya[1].FullName)

}
