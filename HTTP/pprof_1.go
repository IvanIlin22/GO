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

func handle(w http.ResponseWriter, r *http.Request) {
	s := ""
	for i := 0; i < 1000; i++ {
		p := &Post{ID: i, Text: "new post"}
		s += fmt.Sprintf("%#v", p)
		w.Write([]byte(s))
	}
}

//curl http://127.0.0.1:8080/debug/pprof/heap -o mem_out.txt
//curl http://127.0.0.1:8080/debug/pprof/profile?seconds=5 -o cpu_out.txt

//Построение при помощи graviz
//go tool pprof -svg -alloc_objects pprof_1.exe mem_out.txt > mem_ao.svg
//go tool pprof -svg pprof_1.exe cpu_out.txt > cpu.svg
func main() {
	http.HandleFunc("/", handle)

	fmt.Println("server is listening")
	http.ListenAndServe(":8080", nil)
}
