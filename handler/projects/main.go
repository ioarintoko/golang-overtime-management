package projects

import (
	"database/sql"
	"net/http"
)

func Main(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	Routes(db, w, r)
}