package main

import (
	"fmt"

)

func kDifference(a []int32, k int32) int32 {
	fmt.Println(a)
	var temp int32=0

	for i:=len(a)-1; i>=1; i-- {
		fmt.Println("i",i)
		for j:=0; j<i;j++ {
			fmt.Println("j",j)
			K:=a[j]-a[i]
			fmt.Println("K",K)
			if(k==K) {
				temp++
			}else if K<k {
				break
			}
		}
	}
	fmt.Println("Temp", temp)
	return temp
}

func main() {
   var hasil int32 =kDifference([]int32 {5,1,5,3,4,2},2)
   fmt.Println(hasil)
}
