//Package database :
// describing relation with database
package database

import (
	"fmt"
	"sync"

	"mia/core/model"

	"github.com/jinzhu/gorm"
	// necessary to gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//MIADatabase :
// The main structure that will allow us to interact with the database
type MIADatabase struct {
	db       *gorm.DB
	pathToDB string
	m        *sync.Mutex
}

//NewMIADatabase :
// Initialize the structure
func NewMIADatabase(p string) MIADatabase {
	res := MIADatabase{
		db:       nil,
		pathToDB: p,
		m:        &sync.Mutex{},
	}

	return res
}

//OpenConnection :
// Starts a connection with the database
// Must call CloseConnection after manipulating
// The default pattern is :
//		idb.OpenConnection()
//		defer idb.CloseConnection() //nolint: errcheck
func (idb *MIADatabase) OpenConnection() {
	idb.m.Lock()
	db, err := gorm.Open("sqlite3", idb.pathToDB)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "mia_" + defaultTableName
	}

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	// db.LogMode(true)
	idb.db = db
}

//CloseConnection :
// Ends a connection
// Must be called after OpenConnection
func (idb *MIADatabase) CloseConnection() {
	if idb.db == nil {
		panic("failed to close database, wasn't opened")
	}

	err := idb.db.Close()

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to close database")
	}

	idb.db = nil
	idb.m.Unlock()
}

//DB :
// Getter return the gorm object
// This function can be used in calls between OpenConnection and Close Connection
// OpenConnection() --> DB() --> CloseConnection
func (idb MIADatabase) DB() *gorm.DB {
	return idb.db
}

//getListOfObjects :
// returns the list of objects that will be represented in the database
func (MIADatabase) getListOfObjects() []interface{} {
	return []interface{}{
		&model.Role{},
		&model.Item{},
		&model.ItemLink{},
		&model.Assignment{},
		&model.Permission{},
	}
}

func (idb MIADatabase) dbInit() {
	for _, elem := range idb.getListOfObjects() {
		fmt.Printf("Create Table %T\n", elem)
		idb.db.AutoMigrate(elem)
	}

	fmt.Printf("Create Empty Domain\n")

	for _, root := range model.GetRoots() {
		idb.db.Create(root)
	}
}

func (idb MIADatabase) dbDrop() {
	for _, elem := range idb.getListOfObjects() {
		fmt.Printf("Dropping Table %T\n", elem)
		idb.db.DropTableIfExists(elem)
	}
}

//Setup :
func (idb *MIADatabase) Setup(dropTables bool) {
	idb.OpenConnection()
	defer idb.CloseConnection() //nolint: errcheck

	if dropTables {
		idb.dbDrop()
	}

	idb.dbInit()
}
