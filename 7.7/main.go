package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc("/", rootHandler)
	mux := http.NewServeMux()
	fmt.Errorf
	http.ListenAndServe("localhost:8889", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	result := "this is result"
	w.Write([]byte(result))
}
