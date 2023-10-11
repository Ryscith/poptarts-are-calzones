package main

import (
	"fmt"
	"log"

	//	"log"
	"net/http"
	"text/template"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	homepage, _ := template.ParseFiles("index.html")

	// sandwichFiles, _ := os.ReadDir("./img/sandwiches")
	// var sandwichFilenames []string
	// for _, file := range sandwichFiles {
	// 	fullPath := "img/sandwiches/" + file.Name()
	// 	sandwichFilenames = append(sandwichFilenames, fullPath)
	// }

	homepageData := struct {
		Reasons    []string
		Sandwiches []string
	}{
		Reasons: []string{
			"Sandwiches are based",
			"God is based",
			"Therefore God Exists",
		},
		Sandwiches: []string{
			"https://upload.wikimedia.org/wikipedia/commons/2/24/Bologna_sandwich.jpg",
			"img/sandwich_showcase.jpg",
			"img/sandwiches/580b57fcd9996e24bc43c1b7.png",
		},
	}

	homepage.Execute(w, homepageData)

}

func main() {

	http.HandleFunc("/", helloHandler)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("."))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
