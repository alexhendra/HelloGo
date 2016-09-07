package main

import (
	"net/url"
	"fmt"
)

func main() {
	var urlString = "http://localhost:8080/hello?name=Alex&age=27"

	// url.Parse(stringUrl) digunakan utk parsing string ke bentuk url (url.URL)
	var u, err = url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("url: %s\n", urlString)
	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	// Query yg ada pada url akan otomatis diparsing menjadi bentuk map[string][]string
	// dengan Key adalah nama elemen query, dan value berupa array string yg berisikan value dari elemen query
	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name,age)

}
