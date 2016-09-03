package main

import (
	"encoding/base64"
	"fmt"
)

// Golang memiliki package "encoding/base64", yg berisikan fungsi2 utk kebutuhan encode dan decode data
// ke base64 atau sebaliknya. Data yg akan di-encode harus bertipe []byte, perlu dilakukan casting utk
// data2 yg belum bertipe []byte


func main() {
	var data = "Joko Suleso"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded :", encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	fmt.Printf("%v\n",decodedByte)

	var decodedString = string(decodedByte)
	fmt.Println("Decoded: ",decodedString)
}
