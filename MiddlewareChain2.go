package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello , you ahve hit %s \n", r.URL.Path)
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

type headerSetter struct {
	val     string
	handler http.Handler
}

func (hs headerSetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Hello", hs.val)
	hs.handler.ServeHTTP(w, r)
}

func newHeaderSetter(val string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return headerSetter{val, h}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	lh := logger(mux)
	hs := newHeaderSetter("krishna")(lh)
	http.ListenAndServe(":80", hs)
}
