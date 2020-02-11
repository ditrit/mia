//Package database :
// describing relation with database
package database

import (
	"fmt"

	model "iam/core/model"

	"github.com/jinzhu/gorm"
	// necessary to gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//IAMDatabase :
// The main structure that will allow us to interact with the database
type IAMDatabase struct {
	db       *gorm.DB
	pathToDB string
}

//NewIAMDatabase :
// Initialize the structure
func NewIAMDatabase(p string) IAMDatabase {
	res := IAMDatabase{
		db:       nil,
		pathToDB: p,
	}

	return res
}

//getListOfObjects :
// returns the list of objects that will be represented in the database
func (IAMDatabase) getListOfObjects() []interface{} {
	return []interface{}{
		&model.Domain{},
		&model.DomainLink{},
		&model.Object{},
		&model.ObjectLink{},
		&model.Subject{},
		&model.SubjectLink{},
		&model.Role{},
		&model.Assignment{},
		&model.Permission{},
	}
}

func (idb IAMDatabase) dbInit() {
	for _, elem := range idb.getListOfObjects() {
		idb.db.AutoMigrate(elem)
	}

	emptyModel, _ := model.NewDomain("")
	idb.db.Create(&emptyModel)
}

func (idb IAMDatabase) dbReset() {
	for _, elem := range idb.getListOfObjects() {
		idb.db.DropTableIfExists(elem)
	}
}

//Setup :
func (idb IAMDatabase) Setup() {
	db, err := gorm.Open("sqlite3", idb.pathToDB)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close() //nolint: errcheck

	idb.dbReset()
	idb.dbInit()
}
