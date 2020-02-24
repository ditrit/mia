//Package database :
// describing relation with database
package database

import (
	"fmt"
	"sync"

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
	m        *sync.Mutex
}

//NewIAMDatabase :
// Initialize the structure
func NewIAMDatabase(p string) IAMDatabase {
	res := IAMDatabase{
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
func (idb *IAMDatabase) OpenConnection() {
	idb.m.Lock()
	db, err := gorm.Open("sqlite3", idb.pathToDB)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "iam_" + defaultTableName
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
func (idb *IAMDatabase) CloseConnection() {
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
func (idb IAMDatabase) DB() *gorm.DB {
	return idb.db
}

//getListOfObjects :
// returns the list of objects that will be represented in the database
func (IAMDatabase) getListOfObjects() []interface{} {
	return []interface{}{
		&model.Role{},
		&model.Item{},
		&model.ItemLink{},
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

	for _, root := range model.GetRoots() {
		idb.db.Create(root)
	}
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
