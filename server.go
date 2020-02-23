package main

import (
	"fmt"
	"net/http"
)

type message string

func (m message) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, m)
}

func main() {
	msgHandler := message("Hello from Server!")
	fmt.Println("Starting...")
	http.ListenAndServe("localhost:8080", msgHandler) //msgHandler - интерфейс type Handler interface
}
