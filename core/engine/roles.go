// Package engine :
// Handling roles queries from MIA
package engine

import (
	"errors"
	"mia/core/database"
	"mia/core/model"

	"github.com/jinzhu/gorm"
)

// use this function to have the benefir of abstraction and closures
func askDBForRoles(
	idb database.MIADatabase,
	name string,
	haveToOpenConnection bool,
	fIfFound func(*gorm.DB, model.Role) (model.Role, error),
	fIfNotFound func(*gorm.DB, model.Role) (model.Role, error),
) (model.Role, error) {
	var role model.Role

	if haveToOpenConnection {
		idb.OpenConnection()
		defer idb.CloseConnection() //nolint: errcheck
	}

	res := idb.DB().Where("name = ?", name).Take(&role)

	if res.Error != nil && !res.RecordNotFound() {
		return role, errors.New("unknown error occurred")
	}

	if res.RecordNotFound() {
		return fIfNotFound(res, role)
	}

	return fIfFound(res, role)
}

//AddRole :
// add role in the MIA
// returns an error if :
//	- the role already exists
//	- the name is not conform
func AddRole(
	idb database.MIADatabase,
	haveToOpenConnection bool,
	name string,
) error {
	role, err := model.NewRole(name)

	if err != nil {
		return err
	}

	_, err = askDBForRoles(idb, name, haveToOpenConnection,
		func(db *gorm.DB, r model.Role) (model.Role, error) {
			return r, errors.New("the role already exists")
		},
		func(db *gorm.DB, _ model.Role) (model.Role, error) {
			db.Error = nil
			res := db.Create(role)
			return *role, res.Error
		},
	)

	return err
}

//RemoveRole :
// remove role in the MIA
// returns an error if :
//	- the role does not exist in the mia
//	- the role is present in a assignment TODO
func RemoveRole(
	idb database.MIADatabase,
	haveToOpenConnection bool,
	name string,
) error {
	role, err := model.NewRole(name)

	if err != nil {
		return err
	}

	_, err = askDBForRoles(idb, name, haveToOpenConnection,
		func(db *gorm.DB, r model.Role) (model.Role, error) {
			res := db.Delete(&role)
			return r, res.Error
		},
		func(db *gorm.DB, r model.Role) (model.Role, error) {
			return r, errors.New("the role doesn't exist")
		},
	)

	return err
}

// GetRole :
// get role in the MIA
// returns an error if :
//	- the role does not exist in the mia
func GetRole(
	idb database.MIADatabase,
	haveToOpenConnection bool,
	name string,
) (model.Role, error) {
	return askDBForRoles(idb, name, haveToOpenConnection,
		func(db *gorm.DB, r model.Role) (model.Role, error) {
			return r, db.Error
		},
		func(db *gorm.DB, r model.Role) (model.Role, error) {
			return r, errors.New("the role doesn't exist")
		},
	)
}

// IsRoleExists :
// Does the role exists
// returns an error if :
//	- a strange error occurred
func IsRoleExists(
	idb database.MIADatabase,
	haveToOpenConnection bool,
	name string,
) (bool, error) {
	notFoundErr := errors.New("the role doesn't exist")

	_, err := askDBForRoles(idb, name, haveToOpenConnection,
		func(db *gorm.DB, r model.Role) (model.Role, error) {
			return r, db.Error
		},
		func(db *gorm.DB, r model.Role) (model.Role, error) {
			return r, notFoundErr
		},
	)

	if err == nil {
		return true, nil
	}

	if err == notFoundErr {
		return false, nil
	}

	return false, err
}

// GetAllRoles :
// returns all roles in the MIA
func GetAllRoles(
	idb database.MIADatabase,
	haveToOpenConnection bool,
) ([]string, error) {
	var roles []model.Role

	if haveToOpenConnection {
		idb.OpenConnection()
		defer idb.CloseConnection() //nolint: errcheck
	}

	db := idb.DB().Find(&roles)

	res := make([]string, len(roles))

	for i, role := range roles {
		res[i] = role.Name
	}

	return res, db.Error
}
