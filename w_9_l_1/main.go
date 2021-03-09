package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.FileServer(http.Dir("/usr/share/doc"))
	http.Handle("/", http.FileServer(http.Dir("/home/shant/Pictures/")))
	// http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8888", nil)

}

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Welcome to our first GoLang HOME webpage")
}

func about(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Welcome to our first about webpage")
}

func contact(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Welcome to our first Contact webpage")
}
