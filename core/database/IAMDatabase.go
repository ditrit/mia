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

//OpenConnection :
// Starts a connection with the database
// Must call CloseConnection after manipulating
// The default pattern is :
//		idb.OpenConnection()
//		defer idb.CloseConnection() //nolint: errcheck
func (idb *IAMDatabase) OpenConnection() {
	db, err := gorm.Open("sqlite3", idb.pathToDB)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	idb.db = db
}

//CloseConnection :
// Ends a connection
// Must be called after OpenConnection
func (idb *IAMDatabase) CloseConnection() {
	err := idb.db.Close()
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to close database")
	}

	idb.db = nil
}

//DB :
// Getter return the gorm object
// This function can be used in calls between OpenConnection and Close Connection
// OpenConnection() --> DB() --> CloseConnection
func (idb IAMDatabase) DB() *gorm.DB {
	return idb.db
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

	idb.db.Create(model.GetRootDomain())
}

func (idb IAMDatabase) dbDrop() {
	for _, elem := range idb.getListOfObjects() {
		fmt.Printf("Dropping Table %T\n", elem)
		idb.db.DropTableIfExists(elem)
	}
}

//Setup :
func (idb *IAMDatabase) Setup(dropTables bool) {
	idb.OpenConnection()
	defer idb.CloseConnection() //nolint: errcheck

	if dropTables {
		idb.dbDrop()
	}

	idb.dbInit()
}
