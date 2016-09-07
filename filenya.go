package main

import (
	"fmt"
	"os"
	"io"
)

var path = "/home/alex/Documents/testbuatfiledarigolang.txt"

func checkError(err error)  {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func deleteFile() {
	var err = os.Remove(path)
	checkError(err)
}

func readFile() {
	// Buka file, file akses hanya "Read"
	var file,err = os.OpenFile(path,os.O_RDONLY,0644)
	checkError(err)
	defer file.Close()

	// Baca file
	var text = make([]byte, 1024)
	for{
		// file.Read() akan menyimpan isi file yg dibaca ke variable "text"
		n,err := file.Read(text)

		// apabila error bukan dari io.EOF
		// io.EOF menandakan bahwa file yg sedang dibaca adalah baris terakhir isi file
		if err != io.EOF {
			checkError(err)
		}

		if n==0 {
			break
		}
	}
	fmt.Println(string(text))
	checkError(err)
}

func writeFile()  {
	// Buka file dgn level akses Read & Write
	var file,err = os.OpenFile(path,os.O_RDWR,0644)
	checkError(err)
	defer file.Close()

	// Tulis data ke file
	_,err = file.WriteString("halo\n")
	checkError(err)
	_,err = file.WriteString("mari belajar golang\n")
	checkError(err)

	// Simpan perubahan
	err =  file.Sync()
	checkError(err)
}

func createFile() {
	// Memeriksa apakah file sudah ada
	var _,err = os.Stat(path)

	// Jika belum ada maka buat file
	if os.IsNotExist(err) {
		var file,err = os.Create(path)
		checkError(err)
		defer file.Close() // baris kode ini dieksekusi terakhir, karena ditulis defer
	}

	
}

func main() {
	createFile()
	writeFile()
	readFile()
}
