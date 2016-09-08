package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `{"Name":"Alex Hendra", "Age":27}`
	var jsonData = []byte(jsonString)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData,&data1)

	fmt.Println(data1["Name"])

	// Decode JSON ke Interface{}
	var data2 interface{}
	json.Unmarshal(jsonData,&data2)

	// Utk dapat mengakses nilai dari properti maka harus dilakukan casting terlebih dahulu ke map[string]interface{}
	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :",decodedData["Name"])
	fmt.Println("age :",decodedData["Age"])
}
