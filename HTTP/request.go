package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "getHandler: incominr request %#v\n", r)
		fmt.Fprintf(w, "getHandler: r.Url %#v\n", r.URL)
	})
	http.HandleFunc("/raw_body", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintf(w, "postHandler: raw body %s\n", string(body))
	})

	fmt.Println("Server is listening")
	http.ListenAndServe(":8080", nil)
}

func runGet() {
	url := "http://127.0.0.1:8080/?param=123&param2=test"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error happened", err)
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("http.Get body %#v\n\n\n", string(respBody))
	//val := resp.Request.FormValue("param")
	//val2 := resp.Request.URL.Query().Get("param2")
	//fmt.Println(val, val2)
}

func runGetFullReq() {
	req := &http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"User-Agent": {"coursera/golang"},
		},
	}
	req.URL, _ = url.Parse("http://127.0.0.1:8080/?id=42")
	req.URL.Query().Set("user", "rvasily")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error happened", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error happened", err)
	}
	fmt.Printf("testGetFullReq resp %#v\n\n\n", string(body))
}

func runTransportAndPost() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			DualStack: true,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}
	data := `{"id": 42, "user": "rvasily"}`
	body := bytes.NewBufferString(data)

	url := "http://127.0.0.1:8080/raw_body"
	req, _ := http.NewRequest(http.MethodPost, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error happened", err)
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("runTransport %#v\n\n\n", string(respBody))
}

func main() {
	go startServer()
	time.Sleep(100 * time.Millisecond)

	runGet()
	runGetFullReq()
	runTransportAndPost()
}
