package handler

import (
	"HappyHomes/handler/projects"
	"HappyHomes/handler/settings"
	"HappyHomes/handler/tasks"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

var DB *sql.DB

func RegisDB(db *sql.DB) {
	DB = db
}

const (
	task    = "task"
	setting = "setting"
	project = "project"
)

func API(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST, DELETE")

	url := r.URL.Path
	dataURL := strings.Split(fmt.Sprintf("%v", url), "/")

	switch dataURL[2] {
	case task:
		tasks.Main(DB, w, r)

	case setting:
		settings.Main(DB, w, r)

	case project:
		projects.Main(DB, w, r)

	default:
		fmt.Println("Wrong Path")

	}
}