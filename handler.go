package main

import (
	"fmt"
	"log"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello , you ahve hit %s \n", r.URL.Path)
}

func main() {
	err := http.ListenAndServe(":80", HelloHandler{})
	log.Fatal(err)

}
