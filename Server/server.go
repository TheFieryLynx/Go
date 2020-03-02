package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type person struct {
	Name string `json:"Name"`
	Date string `json:"Date"`
	Data string `json:"Data"`
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) string {
	sym := make([]byte, len)
	for i := 0; i < len; i++ {
		sym[i] = byte(randInt(97, 122))
	}
	return string(sym)
}



//HelloHandler для корневого запроса
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello from Server!")
	io.WriteString(w, `Hello from Server!`)
}

//InfoHandler для info
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	var man person
	man.Name = r.URL.Query().Get("name")
	if len(man.Name) == 0 {
		man.Name = "anon"
	}
	man.Date = string(time.Now().Format("02.01.2006 15:04:05"))
	man.Data = randomString(20)
	//fmt.Fprintf(w, "Name: %s\nDate: %s\nData: %s\n\n\n", man.Name, man.Date, man.Data)

	var jsonData []byte
	jsonData, _ = json.Marshal(man)
	//fmt.Fprintf(w, "%s", string(jsonData))
	io.WriteString(w, string(jsonData))
}

// func handler() http.Handler {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", HelloHandler)
// 	mux.HandleFunc("/info", InfoHandler)
// 	return mux
// }

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/user", InfoHandler)
	fmt.Println("Starting...")
	http.ListenAndServe(":8080", nil)
}
