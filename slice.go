package main

import "fmt"

func main() {
  s:=make([]string, 3)
  
  for i:=0;i<len(s);i++ {
    s[i]="1"
  }
  
  for index, num := range s{
    fmt.Println(index,":", num)
  }
  
  // clone array
  c:=make([]string,len(s))
  copy(c,s)
  
  fmt.Println(c)
  
  // menambah data baru
  s=append(s, "d")
  s=append(s,"e","f")
  fmt.Println(s)
  
  // membuat array baru dari array lain, dimulai dari index 2 sampai index sebelum 5
  l:=s[2:5]
  fmt.Println(l)
  
  // mengambil item dari index awal(0) sampai ke index sebelum 5
  l=s[:5]
  fmt.Println(l)
  
  // mengambil item dari index 2 sampai ke index terakhir
  l=s[2:]
  fmt.Println(l)
}


