package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "http://google.com"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	var decodedByte,_ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println(decodedString)
}
