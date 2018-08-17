package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Hello(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:5000", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	Hello(res, req)

	exp := "Hello Web"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected: %s; Actually Got: %s", exp, act)
	}
}
