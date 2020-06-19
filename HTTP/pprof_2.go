package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type Post struct {
	ID      int
	Text    string
	Author  string
	Comment int
	Time    time.Time
}

func getPost(out chan []Post) {
	posts := []Post{}
	for i := 0; i < 10; i++ {
		post := Post{ID: i, Text: "text"}
		posts = append(posts, post)
	}
	out <- posts
}

func handleLeak(w http.ResponseWriter, r *http.Request) {
	res := make(chan []Post)
	go getPost(res)
}

//ab -n 1000 -c 10 http://127.0.0.1:8080/
//curl http://localhost:8080/debug/pprof/goroutine?debug=2 -o goroutines.txt

//Построение при помощи graviz
//

func main() {
	http.HandleFunc("/", handleLeak)

	fmt.Println("server is listening")
	http.ListenAndServe(":8080", nil)
}
