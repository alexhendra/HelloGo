package main

import (
  "fmt"
  "strings"
  "math"
)

// penulisan functions
// *** cara 1 ***
func cara1(a int, b int)int {
  return a+b
}

// *** cara 2 ***
func cara2(a,b,c,d int)int{
  return (a+b)+c*d
}

// *** cara 3 ***
// tanpa return value
func cara3(arr []string) {
  for _,val:=range arr {
    fmt.Println(val)
  }
  
  concatStr:=strings.Join(arr," ")
  fmt.Println("concatStr:", concatStr)
}

// *** cara 4 ***
// multiple return value
func cara4()(string,int) {
  return "Alex",1
}

// *** cara 5 ***
// return value dideclare di variable
func calculate(x float64) (area float64, circumference float64) {
  area=math.Pi * math.Pow(x/2,2)
  circumference=math.Pi * x
  
  // cukup dituliskan retun tanpa disertai nama variablenya
  return
}

func main() {
  hasil:=cara1(3,9)
  hasil2:=cara2(1,2,2,3)
  fmt.Println(hasil)
  fmt.Println(hasil2)
  cara3([]string{"Jhon","Richard"})
  
  strVal,nilai:=cara4()
  fmt.Println(strVal,nilai)
  
  a,c:=calculate(5.0)
  fmt.Println("calculate:",a,c)
  
}


