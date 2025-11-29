package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{'Hello' : '%s'", r.URL.Path)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{'Hello' : 'Test Endpoint'}")
}

func main() {
	fmt.Println("Hello!")
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8080", nil)
}
