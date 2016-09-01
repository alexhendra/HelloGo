package main

import (
	"math"
	"fmt"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	// Embedded Interface
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

// ====  Method pointer  ====
func (k *kubus) volume() float64  {
	return math.Pow(k.sisi,3)
}

func (k *kubus) keliling() float64  {
	return k.sisi * 12
}

func (k *kubus) luas() float64  {
	return math.Pow(k.sisi,2) * 6
}
// =========================

func main() {
	// karena struct kubus memiliki method pointer
	// maka apabila ditampung oleh variable dgn tipe interface, yang diambil adalah alamat memory (referensi) dari kubus
	var ruangBangun hitung = &kubus{4}
	fmt.Println("=========== kubus")
	fmt.Println("Luas :", ruangBangun.luas())
	fmt.Println("Keliling :", ruangBangun.keliling())
	fmt.Println("Volume :", ruangBangun.volume())

}