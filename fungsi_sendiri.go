package main

import (
	"strings"
	"fmt"
	"math/rand"
	"time"
	"math"
)

func main() {
	//var names = []string{"John","wick"}
	//printMessage("Halo",names)

	// Fungsi rand.Seed() digunakan untuk memastikan bahwa angka random yang akan di-generate benar2 acak
	// Fungsi rand.Seed() berada di package "math/rand"
	rand.Seed(time.Now().Unix())
	var randomValue int
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	fmt.Println(time.Now())

	fmt.Println()

	// Implementasi fungsi calculate
	var diameter float64 = 15
	// variable penampung return value dari fungsi calculate
	var area, circumference = calculate(diameter)
	fmt.Printf("luas lingkaran\t\t:%.2f\n", area)
	fmt.Printf("keliling lingkaran\t:%.2f\n", circumference)

	// fungsi math.Pow() untuk melakukan pemangkatan angka
	fmt.Println("3 pangkat 2:",math.Pow(3,2))
	// Konstanta math.Pi merepresentasikan nilai pi atau 22/7
	fmt.Println("Pi (22/7):", math.Pi)
}

// Fungsi yang mengembalikan nilai lebih dari 1 (multiple return)
func calculate(d float64) (float64, float64) {
	// Menghitung luas keliling
	var area = math.Pi * math.Pow(d / 2, 2)

	// Hitung keliling
	var circumReference = math.Pi * d

	return area, circumReference
}

// Fungsi yang berparameter string dan slice string
// Dan merupakan fungsi yang tidak memiliki return value
func printMessage(message string, arr []string) {
	// strings.Join() digunakan untuk menggabungkan item string yang ada di slice "arr"
	// strings.Join() berada di package "strings"
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}


// Fungsi yang memiliki return value
func randomWithRange(min, max int) int {
	var value = rand.Int() % (max - min + 1) + min
	return value
}
