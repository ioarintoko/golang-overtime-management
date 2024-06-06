package model

import (
	"HappyHomes/lib"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Task struct {
	IdTask 			int    			`json:"idtask"`
	Name      		string 			`json:"name"`
	IdProject      	int    			`json:"idproject"`
	ProjectName		string			`json:"projectname"`
	StartDate		time.Time		`json:"startdate"`
	EndDate			time.Time		`json:"enddate"`
	StartTime		string			`json:"starttime"`
	EndTime			string			`json:"endtime"`
	Duration		time.Duration	`json:"duration"`
}

var TableTask = lib.Table{
	Name: "Task",
	Field: []string{
		"IdTask INT(2) PRIMARY KEY AUTO_INCREMENT",
		"Name VARCHAR(50)",
		"IdProject INT(2)",
		"StartDate DATETIME",
		"EndDate DATETIME",
		"Duration INT(50)",
	},
}

var ForeignKeyTaskProject = lib.ForeignKey{
	Name: "Task",
	ForeignName: "Project",
	Field: "IdProject",
	ForeignField: "IdProject",
}

func(t *Task) Insert(db *sql.DB) error {
	query := "INSERT INTO Task (Name, IdProject, StartDate, EndDate, Duration) VALUES (?,?,?,?,?)"
	_, err := db.Exec(query, t.Name, t.IdProject, t.StartDate, t.EndDate, t.Duration)
	return err
}

func(t *Task) Delete(db *sql.DB) error {
	query := "DELETE FROM Task WHERE IdTask = ?"
	_, err := db.Exec(query, t.IdTask)
	return err
}

func (t *Task) Update(db *sql.DB, datatask map[string]interface{}) error{
	var kolom = []string{}
	var args []interface{}

	for key, value := range datatask {
		if value == "" {
			continue
		}
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
	}

	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE Task SET %s WHERE IdTask = %d",
	dataUpdate, t.IdTask)
	_, err := db.Exec(query, args...)
	fmt.Println(query)
	return err
}

func (t *Task) Get(db *sql.DB) error {
	query := "SELECT * FROM Task WHERE IdTask = ?"
	err := db.QueryRow(query, t.IdTask).Scan(&t.IdTask, &t.Name, &t.IdProject, &t.StartDate, &t.EndDate, &t.Duration)
	return err
}

func GetsTask(db *sql.DB, params ...string) ([]*Task, error) {
	var kolom = []string{}
	var args []interface{}
	if len(params) != 0 {
		if params[0] != "" {
			dataParams := strings.Split(params[len(params)-1], ";")
			for _, v := range dataParams {
				temp := strings.Split(fmt.Sprintf("%s", v), ",")
				where := fmt.Sprintf("%s %s ?", strings.ToLower(temp[0]), temp[1])
				kolom = append(kolom, where)
				args = append(args, temp[2])
			}			
		}
	}

	dataKondisi := strings.Join(kolom, " AND ")
	var query string
	query = `SELECT t.*, p.Name, TIME(t.StartDate) AS StartTime, TIME(t.EndDate) AS EndTime FROM Task t
	LEFT JOIN Project p ON t.IdProject = p.IdProject
	`
	if dataKondisi != "" {
		query += " WHERE " + dataKondisi
	}

	datatask, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer datatask.Close()

	var result []*Task
	for datatask.Next() {
		each := &Task{}
		err := datatask.Scan(&each.IdTask, &each.Name, &each.IdProject, &each.StartDate, &each.EndDate, &each.Duration, &each.ProjectName, &each.StartTime, &each.EndTime)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}