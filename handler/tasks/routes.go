package tasks

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

func Routes(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	dataURL := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataURL[len(dataURL)-1]

	switch r.Method {
	case http.MethodGet:
		if lastIndex == "task" {
			Gets(db, w, r)
		} else {
			Get(db, w, r)
		}

	case http.MethodDelete:
		Delete(db, w, r)

	case http.MethodPost:
		Insert(db, w, r)

	case http.MethodPut:
		Update(db, w, r)

	default:
		fmt.Println("Wrong Method")
	}
}