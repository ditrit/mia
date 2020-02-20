package database_test

import (
	"fmt"
	"iam/core/database"
	"iam/core/model"
	"os"
	"testing"
)

func testClosingWithoutOpening(t *testing.T, db database.IAMDatabase) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("good behavior")
		}
	}()

	db.CloseConnection()
	t.Errorf("close without opening should panic")
}

func testConcurrency(t *testing.T, db database.IAMDatabase) {
	role, _ := model.NewRole("CTO")
	roleCand, _ := model.NewRole("CTO")

	db.OpenConnection()

	db.DB().AutoMigrate(&role)
	db.DB().Create(&role)
	db.DB().Where("name = ?", "CTO").Take(&role)

	db.CloseConnection()

	c := make(chan int)

	for i := 0; i < 100; i++ {
		go func(c chan int, i int) {
			db.OpenConnection()
			defer db.CloseConnection() //nolint: errcheck

			db.DB().Where("name = ?", "CTO").Take(&roleCand)

			if roleCand.ID != role.ID {
				t.Errorf("paniced")
			}
			c <- i
		}(c, i)
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", <-c)
	}
	fmt.Printf("\n")
}

func TestDatabase(t *testing.T) {
	db := database.NewIAMDatabase("test.db")

	defer os.Remove("test.db") //nolint: errcheck

	db.OpenConnection()
	db.CloseConnection()

	testClosingWithoutOpening(t, db)
	testConcurrency(t, db)
}
