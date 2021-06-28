package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About Page")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "hello.html")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Index Page")
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}
