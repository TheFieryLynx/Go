package main

import (
	"time"
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
        fmt.Println("Status is", res.StatusCode)
    }

	expected := `Hello from Server!`
    if string(b) != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            string(b), expected)
    } else {
        fmt.Println("Output text is", string(b))
    }
}

type TestInfoHTTPHandler struct{}
func (h *TestInfoHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	InfoHandler(w, r)
}
func TestInfoHandler(t *testing.T) {
    h := &TestInfoHTTPHandler{}
    srv := httptest.NewServer(h)
    defer srv.Close()

    t.Run("test1", func(t *testing.T) {  
        res, err := http.Get(fmt.Sprintf("%s/user", srv.URL))
        if err != nil {
            t.Fatalf("could not send GET request: %v", err)
        }
        defer res.Body.Close()
        b,_ := ioutil.ReadAll(res.Body)
        //fmt.Println(string(b))

        if status := res.StatusCode; status != http.StatusOK {
            t.Errorf("handler returned wrong status code: got %v want %v",
                status, http.StatusOK)
        } else {
            fmt.Println("Status is", res.StatusCode)
        }
        var man person
        json.Unmarshal(b, &man)
        //fmt.Println(man.Name)
        expectedName := `anon`
        if man.Name != expectedName {
            t.Errorf("handler returned unexpected Name: got %v want %v",
                man.Name, expectedName)
        } else {
            fmt.Println("Name is ", man.Name)
        }
        expectedDate :=string(time.Now().Format("02.01.2006"))
        if man.Date != expectedDate {
            t.Errorf("handler returned unexpected Date: got %v want %v",
                man.Date, expectedDate)
        } else {
            fmt.Println("Date is ", man.Date)
        }

        expectedLenDate := 20;
        if len(man.Data) != expectedLenDate {
            t.Errorf("handler returned unexpected Len of Data: got %v want %v",
                len(man.Data), expectedLenDate)
        } else {
            fmt.Println("Data is ", man.Data)
        }
    })

    t.Run("test2", func(t *testing.T) {  
        res, err := http.Get(fmt.Sprintf("%s/user?name=Ivan", srv.URL))
        if err != nil {
            t.Fatalf("could not send GET request: %v", err)
        }
        defer res.Body.Close()
        b,_ := ioutil.ReadAll(res.Body)
        //fmt.Println(string(b))

        if status := res.StatusCode; status != http.StatusOK {
            t.Errorf("handler returned wrong status code: got %v want %v",
                status, http.StatusOK)
        } else {
            fmt.Println("Status is", res.StatusCode)
        }
        var man person
        json.Unmarshal(b, &man)
        //fmt.Println(man.Name)
        expectedName := `Ivan`
        if man.Name != expectedName {
            t.Errorf("handler returned unexpected Name: got %v want %v",
                man.Name, expectedName)
        } else {
            fmt.Println("Name is ", man.Name)
        }
        expectedDate :=string(time.Now().Format("02.01.2006"))
        if man.Date != expectedDate {
            t.Errorf("handler returned unexpected Date: got %v want %v",
                man.Date, expectedDate)
        } else {
            fmt.Println("Date is ", man.Date)
        }

        expectedLenDate := 20;
        if len(man.Data) != expectedLenDate {
            t.Errorf("handler returned unexpected Len of Data: got %v want %v",
                len(man.Data), expectedLenDate)
        } else {
            fmt.Println("Data is ", man.Data)
        }
    })
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