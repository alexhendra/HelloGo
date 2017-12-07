package main

import (
	"fmt"
	"time"
)

func main() {
	// Switch pada golang, berbeda dengan bahasa pemgrograman lain
	// Apabila kondisi telah terpenuhi, maka tdk akan dilanjutkan ke pengecekan case selanjutnya

	var point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Pemanfaatan 1 case pada switch bisa memiliki lebih dari 1 kondisi
	var point2 = 8
	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")

	}

	t:=time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t:=i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Alex")

	// Keyword fallthrough dalam switch
	// Gunanya untuk meneruskan pemeriksaan ke case selanjutnya
	var point3 = 7

	switch {
	case point3 == 8:
		fmt.Println("Perfect")
	case (point3 > 3):
		fmt.Println("Awesome")
		fallthrough
	case point3 < 5:
		fmt.Println("You need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need learn more")
		}
	}
}
