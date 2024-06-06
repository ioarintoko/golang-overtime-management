package test

import (
	"HappyHomes/lib"
	"HappyHomes/model"
	"fmt"
	"testing"
)

var dataset = []*model.Setting{
	{
		Name: "Bram 1",
		Rate: 13000,
	},
	{
		Name: "Bram 2",
		Rate: 13000,
	},
}

func TestSetting(t *testing.T) {
	t.Run("Test Insert Table Setting", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		for _, val := range dataset {
			err := val.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Test Update Table Setting", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		dataUpdate := map[string]interface{} {
			"name" : "Bramantio 1",
		}

		err = dataset[0].Update(db, dataUpdate)

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Delete Table Setting", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		dataset := model.Setting{IdSetting: 2}
		err = dataset.Delete(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get Table Setting", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		dataset := model.Setting{IdSetting: 1}
		err = dataset.Get(db)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(dataset)
	})

	t.Run("Test Gets Table Setting", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		dataset, err := model.GetsSetting(db)

		if err != nil {
			t.Fatal(err)
		}

		for _, val := range dataset {
			fmt.Println(*val)
		}
	})
}