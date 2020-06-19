package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
	w.Write([]byte("!!!"))
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server is listening")
	http.ListenAndServe(":8080", nil)
}
