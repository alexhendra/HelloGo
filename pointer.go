package main
import (
  "fmt"
)

func zeroVal(ival int){
  ival=0
}

// Pointer adalah reference atau alamat memory
// Variable pointer akan berisikan alamat memory dari suatu nilai


// variable bertipe pointer ditandai dengan asteriks "*"
func zeroPtr(iptr *int) {
  *iptr=0
}

func main(){
  i:=1
  fmt.Println("initial:", i)
  
  zeroVal(i)
  fmt.Println("zeroVal:", i)
  
  // tanda "&" digunakan untuk mengambil alamat memory dari variable biasa
  // biasa disebut metode referencing
  zeroPtr(&i)
  fmt.Println("zeroPtr:", i)
  
  fmt.Println("pointer:", &i)
  
  a:=8
  var b *int=&a
  
  fmt.Println("a", a)
  fmt.Println("b", b)
  fmt.Println("b", *b)
  
  // karena a & b memiliki alamat memory yang sama, maka ketika salah satu variable diubah
  // akan mengakibatkan perubahan juga di variable lainnya
  *b=3
  
  fmt.Println("a", a)
  fmt.Println("b", b)
  fmt.Println("b", *b)
}