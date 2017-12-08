package main
import (
  "fmt"
)

// Embedded struct adalah struct yang menjadi bagian dari struct lain
type manusia struct {
  nama string
  umur int
  alamat string
}

type pelajar struct {
  kelas int
  rangking int
  manusia
}

type mahasiswa struct{
  nama string
  semester string
  kampus string
  manusia
}

func main(){
  var p1= pelajar{}
  p1.nama="Alex"
  p1.umur=30
  p1.kelas=1
  p1.rangking=1
  
  fmt.Println(p1)
  
  m1:=mahasiswa{}
  // cara menggunakan property yang namanya sama dengan milik parent dari mahasiswa
  m1.manusia.nama="Budi" 
  m1.nama="Sujanto"
  m1.umur=20
  m1.kampus="UIIIIS"
  
  fmt.Println(m1)
  
  // pengisian nilai sub-struct
  p2:=manusia{nama:"Lili", umur:27}
  m2:=mahasiswa{manusia:p2, nama:"Silo", kampus:"ITIS"}
  fmt.Println(m2)
}