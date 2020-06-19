package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("users.html"))
	users := []User{
		User{1, "Vasily", true},
		User{2, "<h2>Ivan</h2>", false},
		User{ID: 3, Name: "Dmitry", Active: true},
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmpl.Execute(writer, struct {
			Users []User
		}{users})
	})

	fmt.Println("Server is listening")
	http.ListenAndServe(":8080", nil)
}
