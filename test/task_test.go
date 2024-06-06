package test

import (
	"HappyHomes/lib"
	"HappyHomes/model"
	"fmt"
	"testing"
	"time"
)

var datatask = []*model.Task{
	{
		Name: "Testing 1",
		IdProject: 1,
		StartDate: time.Now(),
		EndDate: time.Now(),
		Duration: 3600,
	},
	{
		Name: "Testing 2",
		IdProject: 1,
		StartDate: time.Now(),
		EndDate: time.Now(),
		Duration: 3600,
	},
}

func TestTask(t *testing.T) {
	t.Run("Test Insert Table Task", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		for _, val := range datatask {
			err := val.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Test Update Table Task", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		dataUpdate := map[string]interface{} {
			"name" : "Bramantio 1",
		}

		err = datatask[0].Update(db, dataUpdate)

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Delete Table Task", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		datatask := model.Task{IdTask: 2}
		err = datatask.Delete(db)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get Table Task", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)

		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		datatask := model.Task{IdTask: 1}
		err = datatask.Get(db)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(datatask)
	})

	t.Run("Test Gets Table Task", func(t *testing.T) {
		db, err := lib.ConnectMySql(database)
		if err != nil {
			t.Fatal(err)
		}

		defer db.Close()

		datatask, err := model.GetsTask(db)

		if err != nil {
			t.Fatal(err)
		}

		for _, val := range datatask {
			fmt.Println(*val)
		}
	})
}