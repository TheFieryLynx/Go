package main

import (
	"encoding/json"
	"io/ioutil"

	"fmt"
	"net/http"
	"net/http/httptest"
    "testing"
    
)

type TestHelloHTTPHandler struct{}
func (h *TestHelloHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	HelloHandler(w, r)
}
func TestHelloHandler(t *testing.T) {
    h := &TestHelloHTTPHandler{}
    srv := httptest.NewServer(h)
    defer srv.Close()

    res, err := http.Get(fmt.Sprintf("%s/", srv.URL))
    if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
    defer res.Body.Close()
    b,_ := ioutil.ReadAll(res.Body)
   // fmt.Println(string(b))

    if status := res.StatusCode; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    } else {
        fmt.Println("1. Status is", res.StatusCode)
    }

	expected := `Hello from Server!`
    if string(b) != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            string(b), expected)
    } else {
        fmt.Println("2. Output text is OK")
    }
}

type TestInfoHTTPHandler struct{}
func (h *TestInfoHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	InfoHandler(w, r)
}
func TestInfoHandler1(t *testing.T) {
    h := &TestInfoHTTPHandler{}
    srv := httptest.NewServer(h)
    defer srv.Close()

    res, err := http.Get(fmt.Sprintf("%s/user?name=smb", srv.URL))
    if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
    defer res.Body.Close()
    b,_ := ioutil.ReadAll(res.Body)
    fmt.Println(string(b))

    if status := res.StatusCode; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    } else {
        fmt.Println("1. Status is", res.StatusCode)
    }
    
    var man person
    json.Unmarshal(b, &man)
    fmt.Println(man.Name)
}


/*
func TestHelloHandler1(t *testing.T) {
   
    req, _ := http.NewRequest("GET", "http://localhost:8080/user?name=andrew", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)
	handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    } else {
        fmt.Println("1. Status is", rr.Code)
    }

	expected := `Hello from Server!`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    } else {
        fmt.Println("2. Output text is OK")
    }
}
*/