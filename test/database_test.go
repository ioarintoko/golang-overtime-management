package test

import (
	"HappyHomes/lib"
	"HappyHomes/model"
	"testing"
)

var database, databaseDefaultMysql string

func init() {
	database = "happyhomes"
	databaseDefaultMysql = "Mysql"
}

func TestDatabaseMysql(t *testing.T) {
	t.Run("MySQL Connection Testing", func(t *testing.T) {
		db, err := lib.ConnectMySql(databaseDefaultMysql)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()
	})

	t.Run("Drop Table Testing", func(t *testing.T) {
		db, err := lib.ConnectMySql(databaseDefaultMysql)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.DropDB(db, database)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create DB Testing", func(t *testing.T) {
		db, err := lib.ConnectMySql(databaseDefaultMysql)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateDB(db, database)

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create Table Setting", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		err = lib.CreateTable(db, model.TableSetting)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Create Table Project", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.CreateTable(db, model.TableProject)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Create Table Task", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.CreateTable(db, model.TableTask)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Add Foreign Key Task - Project", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		err = lib.AddForeignKey(db, model.ForeignKeyTaskProject)
		if err != nil {
			t.Fatal(err)
		}
	})
}