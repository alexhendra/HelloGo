package main

import (
	"encoding/json"
	"fmt"
)

/*
	Data JSON tipenya adalah []byte, bisa didapat dari file ataupun string (hasil casting)
	dengan json.Unmarshal data tersebut bisa dikonversi menjadi bentuk objek, entah itu dlm bentuk map[string]interface{}
	ataupun objek hasil struct

	Untuk decode data JSON ke variable objek struct, semua level akses propertinya harus public
*/


type User struct {
	// `json:"Name"` adalah tag, digunakan utk mapping data JSON ke property "FullName"
	// Property "FullName" ditugaskan utk menampung data JSON property "Name"
	FullName string `json:"Name"`

	Age int
}

func main() {
	var jsonString = `{"Name":"Alex Hendra", "Age":27}`
	var jsonData = []byte(jsonString)

	var data User

	// json.Unmarshal() melakukan konversi JSON Data ke bentuk Struct
	// variable yg menampung hasil decode harus di-passing sbg pointer (&data)
	var err = json.Unmarshal(jsonData,&data)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("User :",data.FullName)
	fmt.Println("Age :",data.Age)
}
