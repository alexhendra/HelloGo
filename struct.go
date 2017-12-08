package main
import (
  "fmt"
)

// Struct adalah kumpulan definisi variable (atau property), function (atau method)
// mirip seperti class pada konsep OOP, yg bisa dibuatkan instance object-nya
// Property bisa memiliki tipe data yg bervariasi

type pelajar struct {
  nama string
  umur int
  alamat string
}

func main(){
  fmt.Println(pelajar{"Buody", 15, "Jln. Setiajadi"})
  
  fmt.Println(pelajar{nama:"Alex", umur:28})
  
  var p1 pelajar
  p1.nama="Joko"
  p1.umur=25
  fmt.Println(p1)
  
  p2:=pelajar{nama:"Sule", umur:40}
  fmt.Println(p2.nama)
  
  // variable p3 menggunakan alamat memory yg sama dgn variable p2
  p3:=&p2
  p3.nama="Badrul"
  fmt.Println(*p3)
  fmt.Println(p3)
  fmt.Println(p2)
  
}