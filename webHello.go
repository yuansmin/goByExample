// a simple web hello world
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe("localhost:5000", nil)
	// or just below
	// http.ListenAndServe("localhost:5000", http.HandlerFunc(Hello))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world~\n"))
}
