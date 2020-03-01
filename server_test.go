package main

import (
	"io/ioutil"
	"fmt"
	"testing"
	"net/http/httptest"
    "net/http"
	
)
//HealthCheckHandler аваыа


//TestHealthCheckHandler тест
func TestHealthCheckHandler(t *testing.T) {
    srv := httptest.NewServer(handler())
    defer srv.Close()
    res, err := http.Get(fmt.Sprintf("%s/info?name=andrew", srv.URL))
    if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
    defer res.Body.Close()
    b,_ := ioutil.ReadAll(res.Body)
    fmt.Println(string(b))
    
}