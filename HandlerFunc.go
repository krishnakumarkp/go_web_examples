package main

import (
	"fmt"
	"net/http"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello , you have hit %s \n", r.URL.Path)
	})

	http.ListenAndServe(":80", h)

}
