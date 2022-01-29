package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := DBClient.Query("SELECT id, isbn, title, author FROM dbo.Books")
	var id, isbn, title, author string
	books := []book{}

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		rows.Scan(&id, &isbn, &title, &author)
		books = append(books, book{id, isbn, title, author})
	}

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Gets any params
	var paramId string = string(params["id"])
	querryString := fmt.Sprintf("SELECT id, isbn, title, author FROM dbo.Books WHERE id = %v", paramId)

	rows, err := DBClient.Query(querryString)

	if err != nil {
		panic(err.Error())
	}

	var id, isbn, title, author string

	for rows.Next() {
		rows.Scan(&id, &isbn, &title, &author)
		json.NewEncoder(w).Encode(book{id, isbn, title, author})
		return
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bookVar book
	_ = json.NewDecoder(r.Body).Decode(&bookVar)

	querryString := fmt.Sprintf("INSERT INTO [dbo].[books] OUTPUT INSERTED.* VALUES ('%v','%v','%v')", bookVar.Isbn, bookVar.Title, bookVar.Author)

	rows, err := DBClient.Query(querryString)

	if err != nil {
		panic(err.Error())
	}

	var id, isbn, title, author string

	for rows.Next() {
		rows.Scan(&id, &isbn, &title, &author)
		json.NewEncoder(w).Encode(book{id, isbn, title, author})
		return
	}
}
