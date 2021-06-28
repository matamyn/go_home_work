package main

import (
	"fmt"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, m)
}
func main() {
	msgHandler := msg("Hello from Web Server in GO")
	fmt.Println("Server is listening ...")
	http.ListenAndServe("localhost:8181", msgHandler)
}
