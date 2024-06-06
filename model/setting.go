package model

import (
	"HappyHomes/lib"
	"database/sql"
	"fmt"
	"strings"
)

type Setting struct {
	IdSetting int    `json:"idsetting"`
	Name      string `json:"name"`
	Rate      int    `json:"rate"`
}

var TableSetting = lib.Table{
	Name: "Setting",
	Field: []string{
		"IdSetting INT(2) PRIMARY KEY AUTO_INCREMENT",
		"Name VARCHAR(50)",
		"Rate INT(12)",
	},
}

func(s *Setting) Insert(db *sql.DB) error {
	query := "INSERT INTO Setting (Name, Rate) VALUES (?,?)"
	_, err := db.Exec(query, s.Name, s.Rate)
	return err
}

func(s *Setting) Delete(db *sql.DB) error {
	query := "DELETE FROM Setting WHERE IdSetting = ?"
	_, err := db.Exec(query, s.IdSetting)
	return err
}

func (s *Setting) Update(db *sql.DB, dataset map[string]interface{}) error{
	var kolom = []string{}
	var args []interface{}

	for key, value := range dataset {
		if value == "" {
			continue
		}
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
	}

	dataUpdate := strings.Join(kolom, ",")
	query := fmt.Sprintf("UPDATE Setting SET %s WHERE IdSetting = %d",
	dataUpdate, s.IdSetting)
	_, err := db.Exec(query, args...)
	fmt.Println(query)
	return err
}

func (s *Setting) Get(db *sql.DB) error {
	query := "SELECT * FROM Setting WHERE IdSetting = ?"
	err := db.QueryRow(query, s.IdSetting).Scan(&s.IdSetting, &s.Name, &s.Rate)
	return err
}

func GetsSetting(db *sql.DB, params ...string) ([]*Setting, error) {
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
	query = "SELECT * FROM Setting"
	if dataKondisi != "" {
		query += " WHERE " + dataKondisi
	}

	dataset, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	defer dataset.Close()

	var result []*Setting
	for dataset.Next() {
		each := &Setting{}
		err := dataset.Scan(&each.IdSetting, &each.Name, &each.Rate)
		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
}