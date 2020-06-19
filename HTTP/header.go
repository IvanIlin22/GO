package main

import (
	"fmt"
	"net/http"
)

func header(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("RequestID", "d41d8cd98f00b204")

	fmt.Fprintln(w, "You browser is", r.UserAgent())
	fmt.Fprintln(w, "You accept", r.Header.Get("Accept"))
}

func main() {
	http.HandleFunc("/", header)

	fmt.Println("Server is listening")
	http.ListenAndServe(":8080", nil)
}
