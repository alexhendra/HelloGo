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

// *** cara 3 ***
// Immediately-Invoked Function Expression (IIFE)
// IIFE adalah closure yang dipanggil langsung pada saat deklarasi
var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
var newNumbers=func(min int) []int{
	var r []int
	for _,e:=range numbers{
		if e<min{
			continue
		}
		r=append(r,e)
	}
	return r
}(3)

// *** cara 4 ***
// Closure sebagai nilai kembalian
func findMax2(numbers []int, max int) (int, func()[]int) {
	var res []int
	for _,e:=range numbers {
		if e<=max{
			res=append(res,e)
		}
	}

	return len(res),func() []int{
		return res
	}
}


func main() {
  nextIn:=intSeq()
  
  fmt.Println("nextIn():",nextIn())
  fmt.Println("nextIn():",nextIn())
  fmt.Println("nextIn():",nextIn())
  
  othIn:=intSeq()
  fmt.Println("othIn:",othIn())
  
  
  fmt.Println(getMinMax([] int{2,4,5,7,8,22}))

  fmt.Println("filtered number:",newNumbers)

  max:=3
  howMany,getNumbers:=findMax2(numbers,max)
  theNumbers:=getNumbers()

  fmt.Println("Found \t:",howMany)
  fmt.Println("Value \t:",theNumbers)

}


