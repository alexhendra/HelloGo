package main

import "fmt"

func main() {
  
  // *** cara 1 ***
  var chicken map[string]int
  chicken=map[string]int{}
  fmt.Println("chicken:", chicken)
  
  chicken["januari"]=50
  chicken["februari"]=100
  
  fmt.Println(chicken)
  
  // *** cara 2 ***
  temp:=map[string]string{}
  temp["A"]="Alex"
  temp["B"]="Boy"
  temp["C"]="Charlie"
  
  temp2:=map[string]int{"foo":1,"bar":2}
  
  fmt.Println("temp:", temp)
  fmt.Println("temp2:", temp2)
  
  // *** cara 3 ***
  map1:=make(map[string]bool)
  fmt.Println("map1:",map1)
  
  // *** cara 4 ***
  // statement new harus diawali karakter *
  // agar nilai aslinya yg disimpan, bukan data pointer
  map2:=*new(map[string]int)
  fmt.Println("map2:",map2)
  
  // *** delete item ***
  delete(temp,"B")
  fmt.Println("temp:",temp)
  
  // pada saat mengambil data dengan key, sebenarnya menghasilkan 2 retun valuenya
  // yaitu value dari map, dan kondisi apakah exist atau tidak seperti berikut
  valuenya,prs:=temp["B"]
  fmt.Println("prs:",prs,"valuenya:",valuenya)
}


