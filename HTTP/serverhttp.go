package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
	Name string
}

func (t *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Name: ", t.Name, "URL: ", r.URL.Path)
}

func main() {
	testHandler := &Handler{Name: "test"}
	http.Handle("/test/", testHandler)
	rootHandler := &Handler{"root"}
	http.Handle("/", rootHandler)

	fmt.Println("Server is listening")
	http.ListenAndServe(":8080", nil)
}
