package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/request", request)
	http.HandleFunc("/docs", docs)
	http.HandleFunc("/features", features)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8888", nil)

}

func home(rw http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(rw, nil)

	// fmt.Fprintf(rw, "Welcome to our first GoLang HOME webpage")
}

func docs(rw http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(rw, nil)

	// fmt.Fprintf(rw, "Welcome to our first about webpage")
}

func features(rw http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(rw, nil)
	// 	fmt.Fprintf(rw, "Welcome to our first Contact webpage")
}

func request(rw http.ResponseWriter, r *http.Request) {

	// method 1:

	// name := r.FormValue("name")
	// company := r.FormValue("company")
	// email := r.FormValue("email")

	// fmt.Println(name, company, email)
	// fmt.Fprintf(rw, `recieved`) //response

	// method 2:

	r.ParseForm()

	for key, val := range r.Form {
		fmt.Println(key, val)
	}

	fmt.Fprintf(rw, `recieved`)
}
