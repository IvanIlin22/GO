package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type CheckoutResult struct {
	Success bool
	ErrCode string
	Value   int
}

func CheckoutCalc(w http.ResponseWriter, r *http.Request) {
	var operation string

	operation = r.URL.Path[5:]
	r.ParseForm()
	params := r.Form
	var data CheckoutResult
	if len(params) > 2 {
		//Проверка на два числа, если больше ошибка
		data = CheckoutResult{Success: false, ErrCode: "More than two numbers", Value: 0}
		res, _ := json.Marshal(data)
		fmt.Fprintln(w, string(res))
	} else if a, ok := params["a"]; !ok {
		//Проверка на то, что правильно названы и их два
		data = CheckoutResult{Success: false, ErrCode: "misnamed", Value: 0}
		res, _ := json.Marshal(data)
		fmt.Fprintln(w, string(res))
	} else if b, ok := params["b"]; !ok {
		//Проверка на то, что правильно названы и их два
		data = CheckoutResult{Success: false, ErrCode: "misnamed", Value: 0}
		res, _ := json.Marshal(data)
		fmt.Fprintln(w, string(res))
	} else if a, erra := strconv.Atoi(a[0]); erra != nil {
		//Проверка на то, что валидные
		data = CheckoutResult{Success: false, ErrCode: "invalid value", Value: 0}
		res, _ := json.Marshal(data)
		fmt.Fprintln(w, string(res))
	} else if b, errb := strconv.Atoi(b[0]); errb != nil {
		//Проверка на то, что валидные
		data = CheckoutResult{Success: false, ErrCode: "invalid value", Value: 0}
		res, _ := json.Marshal(data)
		fmt.Fprintln(w, string(res))
	} else {
		if operation == "add" {
			add := a + b
			data = CheckoutResult{Success: true, ErrCode: "", Value: add}
			res, _ := json.Marshal(data)
			fmt.Fprintln(w, string(res))
		} else if operation == "sub" {
			sub := a - b
			data = CheckoutResult{Success: true, ErrCode: "", Value: sub}
			res, _ := json.Marshal(data)
			fmt.Fprintln(w, string(res))
		} else if operation == "mul" {
			mul := a * b
			data = CheckoutResult{Success: true, ErrCode: "", Value: mul}
			res, _ := json.Marshal(data)
			fmt.Fprintln(w, string(res))
		} else if operation == "div" {
			if b == 0 {
				data = CheckoutResult{Success: false, ErrCode: "div to zero", Value: 0}
				res, _ := json.Marshal(data)
				fmt.Fprintln(w, string(res))
			} else {
				div := a / b
				data = CheckoutResult{Success: true, ErrCode: "", Value: div}
				res, _ := json.Marshal(data)
				fmt.Fprintln(w, string(res))
			}
		}
	}
}

func main() {
	http.HandleFunc("/api/", CheckoutCalc)

	fmt.Println("Server is listening")
	http.ListenAndServe(":8080", nil)
}
