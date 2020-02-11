//Package database :
// describing relation with database
package database

import (
	"fmt"

	"iam/core/model"

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

func (idb *IAMDatabase) openConnection() {
	db, err := gorm.Open("sqlite3", idb.pathToDB)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	idb.db = db
}

func (idb *IAMDatabase) closeConnection() {
	err := idb.db.Close()
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to close database")
	}

	idb.db = nil
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
		fmt.Printf("Create Table %T\n", elem)
		idb.db.AutoMigrate(elem)
	}

	fmt.Printf("Create Empty Domain\n")

	emptyModel, _ := model.NewDomain("")

	idb.db.Create(&emptyModel)
}

func (idb IAMDatabase) dbDrop() {
	for _, elem := range idb.getListOfObjects() {
		fmt.Printf("Dropping Table %T\n", elem)
		idb.db.DropTableIfExists(elem)
	}
}

//Setup :
func (idb IAMDatabase) Setup() {
	idb.openConnection()
	defer idb.closeConnection() //nolint: errcheck

	idb.dbDrop()
	idb.dbInit()
}
