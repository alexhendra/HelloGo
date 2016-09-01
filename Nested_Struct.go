package main

import "fmt"

// Nested struct adalah struct yang embedded dan dituliskan langsung di dalam struct
type mahasiswa struct {
	person struct{
		name string
		age int
	       }
	grade int
	hobbies []string
}

func main() {
	var mhs1 mahasiswa
	mhs1.person = struct {
		name string
		age int
	}{name : "Bambang", age : 30}
	mhs1.person.name="Alex"
	mhs1.grade = 10
	mhs1.hobbies =[]string{"Renang","Sepak Bola"}

	fmt.Println(mhs1)
}
