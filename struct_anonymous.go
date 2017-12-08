package main
import (
  "fmt"
)

// Anonymous struct adalah struct yang langsung dibuakan object kepada variable
type manusia struct {
  nama string
  umur int
  alamat string
}

// Nested struct di dalam Anonymous struct yang diembed ke dalam struct
type programmer struct {
    people struct{
      name string
      age int
    }
    
    grade int
    skills []string
}

func main(){
  // Anonymous struct
  // *** cara 1 ***
  var p1=struct{
    manusia
    rangking int
  }{}
  
  var p2=struct{
    manusia
    kampus string
  }{
    manusia:manusia{"Budi", 30, "Jl.Raya"},
  }
  
  p1.manusia=manusia{"Alex",21,"Jl.Raya"}
  p1.rangking=1
  fmt.Println(p1)
  
  fmt.Println(p2)
  
  // Anonymous struct
  // *** cara 2 ***
  var tukang struct {
    manusia
    keahlian string
  }
  
  // Anonymous struct
  // *** cara 3 ***
  var a1 = struct { name string; age int } { age: 22, name: "wick" }
  fmt.Println(a1)
  
  tukang.manusia=manusia{"Yuyo", 33, "Jl.Sue"}
  tukang.keahlian="Pasang Bata"
  fmt.Println(tukang)
  
  // Array (Slice) bertipe struct
  allManusia:=[]manusia{
    {nama:"Abdul", umur:22, alamat:"Jl.Sukosuko"},
    {nama:"Zendar", umur:43, alamat:"Jl.Angkasa"},
  }
  
  fmt.Println(allManusia)
  
  // Array (Slice) bertipe embedded struct
  embedAllPelajar:=[]struct{
    manusia
    rangking int
  }{
    {manusia:manusia{"Tole", 50, "Jl.Abadi"}, rangking:1},
    {manusia:manusia{"Eko", 43, "Jl.Abadi2"}, rangking:3},
    {manusia:manusia{"Dodo", 26, "Jl.Abadi3"}, rangking:4},
  }
  
  fmt.Println(embedAllPelajar)
  
  // menggunakan Nested struct
  var prog programmer
  prog.people.name="Alex"
  prog.people.age=29
  prog.skills=[]string {"C++","Java","Golang","Ruby"}
  fmt.Println(prog)
  
}