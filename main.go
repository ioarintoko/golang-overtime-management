package main

import (
	"HappyHomes/handler"
	"HappyHomes/lib"
	"fmt"
	"net/http"
)

var database string
func init() {
	database = "happyhomes"
}

func main() {
	db, err := lib.ConnectMySql(database)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	handler.RegisDB(db)
	http.HandleFunc("/api/", handler.API)
	http.ListenAndServe(":8087", nil)
}