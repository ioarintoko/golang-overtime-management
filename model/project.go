package model

import (
	"HappyHomes/lib"
	"database/sql"
	"fmt"
	"strings"
)

type Project struct {
	IdProject int    `json:"idproject"`
	Name      string `json:"name"`
}

var TableProject = lib.Table{
	Name: "Project",
	Field: []string{
		"IdProject INT(2) PRIMARY KEY AUTO_INCREMENT",
		"Name VARCHAR(50)",
	},
}

func(p *Project) Insert(db *sql.DB) error {
	query := "INSERT INTO Project (Name) VALUES (?)"
	_, err := db.Exec(query, p.Name)
	return err
}

func(p *Project) Delete(db *sql.DB) error {
	query := "DELETE FROM Project WHERE IdProject = ?"
	_, err := db.Exec(query, p.IdProject)
	return err
}

func (p *Project) Update(db *sql.DB, datapro map[string]interface{}) error{
	var kolom = []string{}
	var args []interface{}

	for key, value := range datapro {
		if value == "" {
			continue
		}
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
	}

	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE Project SET %s WHERE IdProject = %d",
	dataUpdate, p.IdProject)
	_, err := db.Exec(query, args...)
	fmt.Println(query)
	return err
}

func (p *Project) Get(db *sql.DB) error {
	query := "SELECT * FROM Project WHERE IdProject = ?"
	err := db.QueryRow(query, p.IdProject).Scan(&p.IdProject, &p.Name)
	return err
}

func GetsProject(db *sql.DB, params ...string) ([]*Project, error) {
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
	query = "SELECT * FROM Project"
	if dataKondisi != "" {
		query += " WHERE " + dataKondisi
	}

	datapro, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer datapro.Close()

	var result []*Project
	for datapro.Next() {
		each := &Project{}
		err := datapro.Scan(&each.IdProject, &each.Name)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}