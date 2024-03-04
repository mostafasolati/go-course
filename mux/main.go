package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func helloHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World! from mux2")
}

type mymux struct {
	pageVisit int
}

type Middleware func(http.Handler) http.Handler

func (m *mymux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.pageVisit++
	switch r.URL.Path {
	case "/a":
		fmt.Fprint(w, "you are in a page")
	case "/b":
		http.Redirect(w, r, "/c", http.StatusTemporaryRedirect)
	case "/c":
		fmt.Fprint(w, "you are in c page")
	case "/visit":
		fmt.Fprintf(w, "page view: %d", m.pageVisit)
	default:
		fmt.Fprint(w, "400000004 not found")
	}
}

func Before(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before:", r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func After(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		fmt.Println("After:", r.URL.Path)
	})
}

func LogTime(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		//handler.ServeHTTP(w, r)
		if r.URL.Path == "/a" {
			fmt.Fprint(w, "Invalid")
		} else {
			handler.ServeHTTP(w, r)
		}
		fmt.Println(r.URL.Path, time.Since(t))
	})
}

func main() {
	mux := http.NewServeMux()
	mux2 := http.NewServeMux()
	mux3 := new(mymux)

	var handler http.Handler = mux3
	handler = Before(handler)
	handler = After(handler)
	handler = LogTime(handler)

	mux.HandleFunc("/hello", helloHandler)
	mux2.HandleFunc("/hello", helloHandler2)

	go http.ListenAndServe("localhost:6060", handler)
	go http.ListenAndServe("localhost:5050", mux)
	http.ListenAndServe("localhost:4040", mux2)
}
