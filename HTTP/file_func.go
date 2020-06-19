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

func IsUserOdd(u *User) bool {
	return u.ID % 2 != 0
}

func Calc(u *User) int {
	/*if u.ID % 2 != 0 {
		return u.ID + 10
	}*/
	return u.ID + 10
}

func main() {
	tmplFunc := template.FuncMap{
		"OddUser": IsUserOdd,
		"Calculate": Calc,
	}
	tmpl, err := template.New("").Funcs(tmplFunc).ParseFiles("func.html")
	if err != nil {
		panic(err)
	}

	users := []User{
		User{1, "Vasily", true},
		User{2, "Ivan", false},
		User{ID: 3, Name: "Dmitry", Active: true},
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		err := tmpl.ExecuteTemplate(writer, "func.html", struct {
			Users []User
		}{users})
		if err != nil {
			panic(err)
		}
	})

	fmt.Println("Server is listening")
	http.ListenAndServe(":8080", nil)
}
