package main

import (
	"fmt"
	"database/sql"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)
const (
	host     = "gorandb.ca9rjrw9e8wu.eu-central-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "mydb"
  )
var db *sql.DB

func init() {
	tmpDB, err := sql.Open("postgres", "host=gorandb.ca9rjrw9e8wu.eu-central-1.rds.amazonaws.com port=5432 user=xxxxx password=xxxxx dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpDB.Ping()
	if err != nil {
	  panic(err)
	}
	fmt.Println("Successfully connected!")
	fmt.Println(tmpDB)
	db = tmpDB
}

func main(){
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))
	http.HandleFunc("/", handleListBook)
	http.HandleFunc("/book.html", handleViewBook)
	http.HandleFunc("/save", handleSaveBook)
	http.HandleFunc("/delete", handleDeleteBook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}