package test

import (
	"HappyHomes/lib"
	"HappyHomes/model"
	"fmt"
	"testing"
)

var datapro = []*model.Project{
	{
		Name: "Online Shop",
	},
	{
		Name: "UI/UX",
	},
}

func TestProject(t *testing.T) {
	t.Run("Test Insert Table Project", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		for _, val := range datapro {
			err := val.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Test Update Table Project", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		dataUpdate := map[string]interface{} {
			"name" : "Online Shop Golang",
		}

		err = datapro[0].Update(db, dataUpdate)

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Delete Table Project", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		datapro := model.Project{IdProject: 2}
		err = datapro.Delete(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get Table Project", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		datapro := model.Project{IdProject: 1}
		err = datapro.Get(db)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(datapro)
	})

	t.Run("Test Gets Table Project", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		datapro, err := model.GetsProject(db)

		if err != nil {
			t.Fatal(err)
		}

		for _, val := range datapro {
			fmt.Println(*val)
		}
	})
}