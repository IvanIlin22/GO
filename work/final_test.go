package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type TestCase struct {
	URL    string
	Result *CheckoutResult
}

func Checkout(ur string, operation string) (*CheckoutResult, error) {
	url := ur + operation
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &CheckoutResult{}

	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func TestCheckout(t *testing.T) {
	cases := []TestCase{
		{
			URL: "/api/add?a=1&b=5",
			Result: &CheckoutResult{
				Success: true,
				ErrCode: "",
				Value:   6,
			},
		},
		{
			URL: "/api/mul?a=2&b=3",
			Result: &CheckoutResult{
				Success: true,
				ErrCode: "",
				Value:   6,
			},
		},
		{
			URL: "/api/div?a=6&b=2",
			Result: &CheckoutResult{
				Success: true,
				ErrCode: "",
				Value:   3,
			},
		},
		{
			URL: "/api/sub?a=6&b=1",
			Result: &CheckoutResult{
				Success: true,
				ErrCode: "",
				Value:   5,
			},
		},
		{
			URL: "/api/sub?a=c&b=3",
			Result: &CheckoutResult{
				Success: false,
				ErrCode: "invalid value",
				Value:   0,
			},
		},
		{
			URL: "/api/div?a=2&b=0",
			Result: &CheckoutResult{
				Success: false,
				ErrCode: "div to zero",
				Value:   0,
			},
		},
		{
			URL: "/api/mul?a=2",
			Result: &CheckoutResult{
				Success: false,
				ErrCode: "misnamed",
				Value:   0,
			},
		},
		{
			URL: "/api/mul?a=2&c=1",
			Result: &CheckoutResult{
				Success: false,
				ErrCode: "misnamed",
				Value:   0,
			},
		},
		{
			URL: "/api/mul?a=2&b=1&c=3",
			Result: &CheckoutResult{
				Success: false,
				ErrCode: "More than two numbers",
				Value:   0,
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(CheckoutCalc))

	for index, value := range cases {
		result, err := Checkout(ts.URL, value.URL)
		if err != nil {
			t.Errorf("[%d] unexpected error: %#v", index, err)
		}
		if !reflect.DeepEqual(value.Result, result) {
			t.Errorf("[%d] wrong result, expected %#v, got %#v", index, value.Result, result)
		}
	}
	ts.Close()
}
