package main

import (
	"os"
	"fmt"
)

/*
	Arguments adalah data opsional yg disisipkan ketika eksekusi program.
*/


func main() {
	// Data arguments didapat dari os.Args
	// Data tersebut disimpan dalam bentuk array dgn pemisah adalah tanda spasi
	// os.Args mengembalikan tak hanya arguments saja, tapi juga path file
	var argsRaw = os.Args
	fmt.Printf("-> %#v\n",argsRaw)

	// argsRaw[1:] gunanya untuk mengambil slice arguments-nya saja
	var args = argsRaw[1:]
	fmt.Printf("-> %#v\n",args)

	// Cara eksekusinya : go run argumentsnya.go "Jeruk" "Apel" Semangka

}
