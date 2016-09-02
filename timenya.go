package main

import (
	"time"
	"fmt"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("Time1 %v\n", time1)

	var time2 = time.Date(2016,7,30,2,40,8,6, time.UTC)
	// format %v, menampilkan bentuk asli dari nilai variable
	fmt.Printf("Time2 %v\n", time2)
}
