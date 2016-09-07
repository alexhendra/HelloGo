package main

import (
	"os/exec"
	"fmt"
)

/*
	Exec digunakan untuk eksekusi perintah command line (CLI) lewat kode program
	cara eksekusi command adalah dengan menuliskan command dalam bentuk string, diikuti arguments-nya(jika ada) -
	sebagai parameter variadic pada fungsi
*/

func main() {
	var output1, _ = exec.Command("ls").Output()
	fmt.Printf("-> ls\n%s\n", string(output1))
	var output2, _ = exec.Command("pwd").Output()
	fmt.Printf("-> pwd\n%s\n", string(output2))

	// Eksekusi perintah command line (CLI) dgn parameter
	var output3, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Printf("-> git config user.name\n%s\n", string(output3))
}
