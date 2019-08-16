package main

import (
	"fmt"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a test")
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "hellow world")

	})

	http.HandleFunc("/test/", viewHandler)
	http.ListenAndServe(":80", nil)

}
