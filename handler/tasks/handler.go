package tasks

import (
	"HappyHomes/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Gets(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	dataParams := r.URL.Query().Get("params")

	dataset, err := model.GetsTask(db, dataParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	jsonData, err := json.Marshal(dataset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write(jsonData)
}

func Get(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	dataURL := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataURL[len(dataURL)-1]

	id, err := strconv.Atoi(lastIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	dataset := &model.Task{IdTask: id}
	data := dataset.Get(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(data)

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write(jsonData)
}

func Delete(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	dataURL := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataURL[len(dataURL)-1]

	id, err := strconv.Atoi(lastIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if lastIndex != "Task" {
		dataset := &model.Task{IdTask: id}
		err := dataset.Delete(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Write([]byte("OK"))
	}
}

func Insert(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	dataURL := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataURL[len(dataURL)-1]

	if lastIndex == "Task" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer r.Body.Close()

		var Task model.Task
		err = json.Unmarshal(body, &Task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		err = Task.Insert(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func Update(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	dataURL := strings.Split(fmt.Sprintf("%v", url), "/")
	lastIndex := dataURL[len(dataURL)-1]

	if lastIndex != "Task" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer r.Body.Close()

		jsonMap := make(map[string]interface{})
		err = json.Unmarshal(body, &jsonMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		fmt.Println(body)
		fmt.Println(jsonMap)

		id, err := strconv.Atoi(lastIndex)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		dataset := model.Task{IdTask: id}
		err = dataset.Update(db, jsonMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.Write([]byte("OK"))
		}
	}
}