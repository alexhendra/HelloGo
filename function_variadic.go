package main

import (
  "fmt"
  "strings"
)

// variadic functions, adalah functions yang memiliki jumlah arguments yg bebas tanpa dideklarasikan jumlah arguments nya
// *** cara 1 ***
func sum(nums ...int) {
  fmt.Print(nums, " ")
  total:=0
  for _,num:=range nums{
    //fmt.Println(i,num)
    total+=num
  }
  fmt.Println(total)
}

// *** cara 2 ***
func tampilkan(nama string, nilai ...string) {
  nilaiConcat:=strings.Join(nilai,",")
  fmt.Printf("Nama saya: %s\n", nama)
  fmt.Printf("Daftar nilai saya: %s\n", nilaiConcat)
}

func main() {
  sum(1,2)
  sum(1,2,3)
  
  nums:=[]int{1,2,3,4,5}
  
  // penulisan parameter yg akan dikirimkan harus menggunakan tanda titik 3x
  sum(nums...)
  
  tampilkan("Alex","89","70")
  tampilkan("Boy") // argument variadic bisa tidak dikirimkan
}


