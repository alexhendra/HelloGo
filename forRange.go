package main

import "fmt"

func main() {
  nums:=[]int {2,3,4}
  sum:=0
  for _,num:= range nums{
    sum+=num
  }
  fmt.Println(sum)
  
  // looping dengan menggunakan 'for range' akan menghasilkan index & value dari array
  for idx,val:= range nums{
    fmt.Printf("index-%d=%d\n",idx,val)
  }
  
  temp:=map[string]string{}
  temp["A"]="Alex"
  temp["B"]="Boy"
  temp["C"]="Charlie"
  
  // looping dengan menggunakan 'for range' akan menghasilkan key & value dari Map
  for ky,val:=range temp{
    fmt.Printf("keys:%s, value:%s\n",ky,val)
  }
  fmt.Println("temp:", temp)
  
  // mengambil key dari Map
  for k:=range temp{
    fmt.Println("Key:", k)
  }
  
  // i akan berisikan index dari string
  // c akan berisikan nilai konversi karakter ke Unicode
  for i,c:=range "BOBBI"{
    fmt.Println(i,c)
  }
}


