package main

import (
	"net/http"
	"fmt"
)

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w,"Apa kabar !")
}

func main() {
	// Fungsi http.HandleFunc() digunakan untuk routing aplikasi web
	// Parameter pertama adalah route yg diinginkan, Parameter kedua adalah callback atau aksi ketika route diakses
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Halo !")
	})

	http.HandleFunc("/index",index)

	fmt.Println("Starting web server at http://localhost:9090")

	// Fungsi http.ListenAndServe() digunakan utk menghidupkan server sekaligus menjalankan aplikasi di server tersebut
	// Di Golang, 1 web aplikasi adalah 1 buah server berbeda
	http.ListenAndServe(":9090",nil)
}
