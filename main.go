package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var books []book

func main() {

	port := flag.String("port", "8000", "Server port")
	user := flag.String("user", "", "DB username")
	password := flag.String("password", "", "DB password")
	address := flag.String("address", "", "DB address")
	dbport := flag.String("dbport", "", "DB port")
	dbname := flag.String("dbname", "", "DB name")

	flag.Parse()

	InitDbConnection(*user, *password, *address, *dbport, *dbname)

	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+*port, r))
}
