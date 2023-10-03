package main

import (
	"fmt"
	//	"log"
	"net/http"
	"text/template"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Reilly Is Really Gay</h1>")
	homepage, _ := template.ParseFiles("index.html")

	homepageData := struct {
		Reasons []string
	}{
		Reasons: []string{
			"Sandwiches are based",
			"God is based",
			"Therefore God Exists",
		},
	}

	homepage.Execute(w, homepageData)

}

func main() {

	http.HandleFunc("/", helloHandler)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
