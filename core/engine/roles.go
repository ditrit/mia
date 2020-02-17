// Package engine :
// Handling roles queries from IAM
package engine

import (
	"errors"
	"iam/core/database"
	"iam/core/model"

	"github.com/jinzhu/gorm"
)

// use this function to have the benefir of abstraction and closures
func askDBForRoles(
	idb database.IAMDatabase,
	name string,
	haveToOpenConnection bool,
	fIfFound func(*gorm.DB, model.Role) error,
	fIfNotFound func(*gorm.DB) error,
) error {
	var role model.Role

	if haveToOpenConnection {
		idb.OpenConnection()
		defer idb.CloseConnection() //nolint: errcheck
	}

	res := idb.DB().Where("name = ?", name).Take(&role)

	if res.Error != nil && !res.RecordNotFound() {
		return errors.New("unknown error occurred")
	}

	if res.RecordNotFound() {
		return fIfNotFound(res)
	}

	return fIfFound(res, role)
}

//AddRole :
// add role in the IAM
// returns an error if :
//	- the role already exists
//	- the name is not conform
func AddRole(
	idb database.IAMDatabase,
	name string,
) error {
	role, err := model.NewRole(name)

	if err != nil {
		return err
	}

	err = askDBForRoles(idb, name, true,
		func(db *gorm.DB, _ model.Role) error {
			return errors.New("the role already exists")
		},
		func(db *gorm.DB) error {
			db.Error = nil
			res := db.Create(role)
			return res.Error
		},
	)

	return err
}

//RemoveRole :
// remove role in the IAM
// returns an error if :
//	- the role does not exist in the iam
//	- the role is present in a assignment TODO
func RemoveRole(
	idb database.IAMDatabase,
	name string,
) error {
	_, err := model.NewRole(name)

	if err != nil {
		return err
	}

	err = askDBForRoles(idb, name, true,
		func(db *gorm.DB, role model.Role) error {
			res := db.Delete(&role)
			return res.Error
		},
		func(db *gorm.DB) error {
			return errors.New("the role doesn't exist")
		},
	)

	return err
}

// GetAllRoles :
// returns all roles in the IAM
func GetAllRoles(
	idb database.IAMDatabase,
) ([]model.Role, error) {
	var roles []model.Role

	idb.OpenConnection()
	defer idb.CloseConnection() //nolint: errcheck

	res := idb.DB().Find(&roles)

	return roles, res.Error
}
