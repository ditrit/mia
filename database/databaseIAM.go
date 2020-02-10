//Package database :
// describing relation with database
package database

import (
	"fmt"

	model "iam/model"

	"github.com/jinzhu/gorm"
	// necessary to gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func dbInit(db *gorm.DB) {
	db.AutoMigrate(&model.Domain{})
	db.AutoMigrate(&model.DomainLink{})
	db.AutoMigrate(&model.Object{})
	db.AutoMigrate(&model.ObjectLink{})
	db.AutoMigrate(&model.Subject{})
	db.AutoMigrate(&model.SubjectLink{})
}

func dbReset(db *gorm.DB) {
	db.DropTableIfExists(&model.Domain{})
	db.DropTableIfExists(&model.DomainLink{})
	db.DropTableIfExists(&model.Object{})
	db.DropTableIfExists(&model.ObjectLink{})
	db.DropTableIfExists(&model.Subject{})
	db.DropTableIfExists(&model.SubjectLink{})
}

//Setup :
func Setup() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close() //nolint: errcheck

	dbReset(db)
	dbInit(db)
}
