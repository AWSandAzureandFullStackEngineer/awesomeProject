package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "New Hello 3!"
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %q want %q",
			body, expected)
	}
}
