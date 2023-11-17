package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	fmt.Println("here")
	http.HandleFunc("/create", CreateHandler(db))
	// http.HandleFunc("/update", UpdateHandler(db))
}
