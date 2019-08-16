package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello , you ahve hit %s \n", r.URL.Path)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":80", mux)
}
