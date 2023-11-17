package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	dataservice "myproject/assignment_4/dataservice"
	"myproject/assignment_4/model"
	"net/http"
)

func CreateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}
	fmt.Println(book.Author)
	if exists, err := dataservice.BookExists(db, book.Title, book.Author); err != nil {
		fmt.Println(err)
		return err
	} else if exists {
		fmt.Println(exists)
		http.Error(w, "Book already exists", http.StatusBadRequest)
		return errors.New("book already exists")
	}
	return dataservice.CreateBook(db, w, r)
}
