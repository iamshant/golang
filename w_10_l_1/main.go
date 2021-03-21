package main

import (
	"fmt"
	"net/http"
	"text/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/request", request)
	http.HandleFunc("/docs", docs)
	http.HandleFunc("/features", features)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8888", nil)

}

var db *sql.DB
var err error
func init(){

	fmt.Println("Go MySQL Tutorial")

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err = sql.Open("mysql", "root:a.12345@tcp(127.0.0.1:3306)/hosting_db")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    // defer db.Close()

 
    fmt.Println("db connection successful")
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

	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")

	// fmt.Println(name, company, email)
	// fmt.Fprintf(rw, `recieved`) //response

	
	// method 2:
	// r.ParseForm()

	// for key, val := range r.Form {
	// 	fmt.Println(key, val)
	// }

	// perform a db.Query insert
    // insert, err := db.Query("INSERT INTO `request` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, 'Shant', 'BAIUST', 'naimshant@gmail.com', '1');")
    qs := "INSERT INTO `request` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, '%s', '%s', '%s', '1');"
    sql := fmt.Sprintf(qs, name, company, email)
    
    fmt.Println(sql)
    insert, err := db.Query(sql)

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()

	fmt.Fprintf(rw, `recieved`)
}
