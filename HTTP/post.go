package main

import (
	"fmt"
	"net/http"
)

var loginFromTmpl = []byte(`
<html>
	<body>
	<form atcion = "/" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type ="submit" value="login">
	</form>
	</body>
</html>`)

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write(loginFromTmpl)
		return
	}
	inputLogin := r.FormValue("login")
	inputPassword := r.FormValue("password")
	fmt.Fprintln(w, "you enter: ", inputLogin, inputPassword)
}

func main() {
	http.HandleFunc("/", mainPage)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
