package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Main page")
}

func main() {
	http.HandleFunc("/page", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.URL.Path)
	})
	http.HandleFunc("/pages/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.URL.Path)
	})
	http.HandleFunc("/", handler)

	fmt.Println("Server is listening")
	http.ListenAndServe(":8080", nil)
}
