package main

import (
	"math"
	"fmt"
)

type hitung interface {
	luas() float64
	keliling() float64
}

// ==== Struct lingkaran  ====

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64  {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(),2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

// ===========================

// ==== Struct persegi ====
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi,2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

// ===========================

func main() {

	var bangunDatar hitung

	// variable objek interface hanya bisa menampung objek yang minimal memiliki semua method yang ada di interfacenya
	bangunDatar = persegi{10.0}
	fmt.Println("======= Persegi")
	fmt.Println("Luas	:", bangunDatar.luas())
	fmt.Println("Keliling	:", bangunDatar.keliling())
	fmt.Println("Bangun data (Persegi) :",bangunDatar)

	bangunDatar = lingkaran{14.0}
	fmt.Println("======= Lingkaran")
	fmt.Println("Luas	:", bangunDatar.luas())
	fmt.Println("Keliling	:", bangunDatar.keliling())

	// Karena method jariJari() tidak ada di interface "hitung"
	// maka dilakukan casting terhadap objek interface "bangunDatar"
	// caranya : namaObjekInterface.(tipeData)
	fmt.Println("Jari-jari	:", bangunDatar.(lingkaran).jariJari())
}
