package main

import (
  "fmt"
)

// Closure adalah function yg bisa disimpan dalam variable atau biasa disebut anonymous function
// dengan konsep ini, maka dapat dibuatkan function di dalam function,
// atau function mengembalikan function

// *** cara 1 ***
func intSeq() func()int{
  i:=0
  return func() int{
    i++
    return i
  }
}

// *** cara 2 ***
var getMinMax=func(n []int)(min int, max int) {
  for i,e:=range n {
    switch{
      case i==0:
        min,max=e,e
      case e > max:
        max=e
      case e<min:
       min=e
    }
  }
  return
}

func main() {
  nextIn:=intSeq()
  
  fmt.Println("nextIn():",nextIn())
  fmt.Println("nextIn():",nextIn())
  fmt.Println("nextIn():",nextIn())
  
  othIn:=intSeq()
  fmt.Println("othIn:",othIn())
  
  
  fmt.Println(getMinMax([] int{2,4,5,7,8,22}))
}


