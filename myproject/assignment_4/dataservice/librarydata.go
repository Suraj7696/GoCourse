package dataservice

import (
	"database/sql"
	"encoding/json"
	"fmt"
	model "myproject/assignment_4/model"
	"net/http"
)

func CreateBook(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}

	fmt.Println(book)

	query := "INSERT INTO books (title, author, publication_year) VALUES (?,?,?)"
	_, err := db.Exec(query, book.Title, book.Author, book.Publication_year)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
	return nil
}

func BookExists(db *sql.DB, title string, author string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM books WHERE title=? and author=?)"
	err := db.QueryRow(query, title, author).Scan(&exists)
	return exists, err
}
