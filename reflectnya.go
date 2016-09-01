package main

import (
	"reflect"
	"fmt"
)

func main() {
	// Reflection merupakan teknik untuk inspeksi sebuah variable, mengambil informasi dari variable
	// atau bahkan memanipulasinya
	// Dapat melihat struktur variable, tipe, nilai pointer, dan banyak lagi.

	var number = 23
	var reflectValue =  reflect.ValueOf(number)
	fmt.Println("Tipe variable :", reflectValue.Type())

	// Dilakukan pembandingan dengan konstanta reflect.Int
	if reflectValue.Kind() == reflect.Int {
		fmt.Println(reflectValue.Kind())
		fmt.Println("Nilai variable :", reflectValue.Int())
		// Fungsi Interface() mengembalikan nilai interface kosong
		fmt.Println("Nilai Interface :", reflectValue.Interface())

		// Nilai aslinya sendiri bisa diakses dengan meng-casting interface kosong
		var nilai = reflectValue.Interface().(int)
		fmt.Println("Nilai Asli :", nilai)

		fmt.Println()

		// ====  Pengaksesan Informasi Property Variable Objek  ====
		var s1 = &student{Name:"Joko", Grade:2}
		s1.getPropertyInfo()
		// =============================================
	}
	// ===========================================

	// ==== Mengakses informasi method dari variable objek ====
	var s1 = &student{Name: "Karsono Joyo", Grade:4}
	fmt.Println("Nama :", s1.Name)

	var reflectValue2 = reflect.ValueOf(s1)

	// Mencari method "SetName"
	var method = reflectValue2.MethodByName("SetName")

	// Mengisi parameter dari method
	// Parameternya diisi dengan array []reflect.Value
	// Dan array tersebut harus dalam bentuk Value dengan cara reflect.ValueOf
	method.Call([]reflect.Value {
		// Mengembalikan nilai value "Joyo"
		reflect.ValueOf("Joyo"),
	})
	fmt.Println("Nama :", s1.Name)
	// ============================================
}

type student struct {
	Name string
	Grade int
}

// ====  Pengaksesan Informasi Property Variable Objek  ====
func (s *student) getPropertyInfo() {
	var reflValue = reflect.ValueOf(s)

	// Check apakah pointer atau tidak
	if reflValue.Kind() == reflect.Ptr {
		reflValue = reflValue.Elem()
	}

	var reflType = reflValue.Type()

	// Perulangan sebanyak jumlah property yang ada pada struct "student"
	for i:=0; i < reflValue.NumField(); i++ {

		// reflType.Field(i).Name , untuk mengambil nama property
		fmt.Println("Nama :", reflType.Field(i).Name)

		// reflType.Field(i).Type, untuk mengambil tipe data property
		fmt.Println("Tipe data :", reflType.Field(i).Type)

		// reflValue.Field(i).Interface() untuk mengambil nilai property dalam bentuk interface kosong "{}"
		fmt.Println("Nilai :", reflValue.Field(i).Interface())
		fmt.Println("")

	}
}

func (s *student) SetName(Nama string)  {
	s.Name = Nama
}

// ====================================================
