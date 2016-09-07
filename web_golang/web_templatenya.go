package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// variable data berisikan nilai yg dikirimkan ke template.html
		var data = map[string]string{
			"Name":"Alex Hendra",
			"Message":"Ini isi message alex",
		}

		// Fungsi template.ParseFiles() melakukan parsing dari template.html
		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
		}

		// Fungsi Execute() akan membuat hasil parsing template ditampilkan ke web browser
		t.Execute(w,data)
	})

	fmt.Println("Starting web server at http://localhost:8181")
	http.ListenAndServe(":8181",	nil)
}
