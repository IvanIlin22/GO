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

func (u *User) PrivatActive() string {
	if !u.Active {
		return ""
	}
	return "method says user " + u.Name + " active"
}
func main() {
	tmpl := template.Must(template.ParseFiles("users.html"))
	users := []User{
		User{1, "Vasily", true},
		User{2, "Ivan", false},
		User{ID: 3, Name: "Dmitry", Active: true},
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		err := tmpl.ExecuteTemplate(writer, "users.html", struct{Users []User}{users})
		if err != nil {
			panic(err)
		}
	})

	fmt.Println("Server is listening")
	http.ListenAndServe(":8080", nil)
}
