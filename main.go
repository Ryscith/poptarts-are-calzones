package main

import (
	"fmt"
	//	"log"
	"net/http"
	"text/template"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	homepage, _ := template.ParseFiles("index.html", "header.html")
	
	// sandwichFiles, _ := os.ReadDir("./img/sandwiches")
	// var sandwichFilenames []string
	// for _, file := range sandwichFiles {
	// 	fullPath := "img/sandwiches/" + file.Name()
	// 	sandwichFilenames = append(sandwichFilenames, fullPath)
	// }

	homepageData := struct {
		Reasons []string
		Sandwiches []string
	}{
		Reasons: []string{
			"Sandwiches are based",
			"God is based",
			"Therefore God Exists",
		},
		Sandwiches: []string{
			"https://upload.wikimedia.org/wikipedia/commons/2/24/Bologna_sandwich.jpg",
		},
	}

	homepage.Execute(w, homepageData)

}

func main() {

	http.HandleFunc("/", helloHandler)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
