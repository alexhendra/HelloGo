package main

import "fmt"

func main() {
	var s1,s2,s3 string

	// Skema fungsi Scanln "func Scanln(a ...interface{})(n int, err error)
	// Fungsi scanln bisa menampung parameter bertipe interface{} dalam jumlah tak terbatas "..."
	// Tiap parameter akan menampung karakter2 inputan user yg sudah dipisah dengan tanda spasi
	fmt.Scanln(&s1,&s2,&s3)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
