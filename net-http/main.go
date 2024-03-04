package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "<h1>Hello World</h1>")
	w.Header().Set("location", "/bye")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bye Bye")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	c := r.PathValue("color")
	size := r.URL.Query().Get("size")
	fmt.Fprintf(w, "<h1 style='color:%s;font-size:%s'>time: %d:%d:%d</h1>", c, size, t.Hour(), t.Minute(), t.Second())
	fmt.Fprintf(w, "Host: %s<br>", r.Host)
	fmt.Fprintf(w, "Method: %s<br>", r.Method)
	fmt.Fprintf(w, "URL Raw Query: %s<br>", r.URL.RawQuery)
	fmt.Fprintf(w, "URL Path: %s<br>", r.URL.Path)
	fmt.Fprintf(w, "RequestURI: %s<br>", r.RequestURI)

}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	image, header, err := r.FormFile("image")
	fmt.Println(err)
	dest, _ := os.Create("public/" + header.Filename)
	defer image.Close()
	defer dest.Close()
	io.Copy(dest, image)

	w.Header().Set("content-type", "text/html")
	fmt.Fprintf(w, "Thank you %s, your inviation sent to your email:%s", name, email)
	fmt.Fprintf(w, "<img src='%s'>", "/"+header.Filename)
	fmt.Println("password:", password)

}

// localhos:5050/time/red?size=20&name=ali&age=22

func main() {
	fileHandler := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fileHandler))
	http.HandleFunc("/time/{color}", timeHandler)
	http.HandleFunc("/bye", byeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("POST /register", registerHandler)
	http.ListenAndServe("localhost:5050", nil)

}
