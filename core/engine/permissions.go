// Package engine :
// Handling permissions queries from MIA
package engine

import (
	"errors"
	"mia/core/constant"
	"mia/core/database"
	"mia/core/model"

	"github.com/jinzhu/gorm"
)

func askDBForPermissions(
	idb database.MIADatabase,
	haveToOpenConnection bool,
	roleName string,
	domainName string,
	objName string,
	act constant.Action,
	fIfFound func(*gorm.DB, model.Permission) error,
	fIfNotFound func(*gorm.DB, model.Permission) error,
) error {
	var (
		permission model.Permission
		role       model.Role
		domain     model.Item
		object     model.Item
	)

	if haveToOpenConnection {
		idb.OpenConnection()
		defer idb.CloseConnection() //nolint: errcheck
	}

	role, err := GetRole(idb, false, roleName)

	if err != nil {
		return err
	}

	domain, err = GetItem(idb, false, model.ITEM_TYPE_DOMAIN, domainName)

	if err != nil {
		return err
	}

	object, err = GetItem(idb, false, model.ITEM_TYPE_OBJ, objName)

	if err != nil {
		return err
	}

	query := idb.DB().Where("id_role = ?", role.ID)
	query = query.Where("id_domain = ?", domain.ID)
	query = query.Where("id_object = ?", object.ID)
	query = query.Where("action = ?", act)
	res := query.Take(&permission)

	if res.Error != nil && !res.RecordNotFound() {
		return errors.New("unknown error occurred")
	}

	permission.IDRole = role.ID
	permission.IDDomain = domain.ID
	permission.IDObject = object.ID
	permission.Action = act

	if res.RecordNotFound() {
		return fIfNotFound(res, permission)
	}

	return fIfFound(res, permission)
}

//AddPermission :
// adds an permission in the mia
// returns an error if
//	- one of the role, the subject or the domain does not exist
//	- the permission already exists
func AddPermission(
	idb database.MIADatabase,
	haveToOpenConnection bool,
	roleName string,
	domainName string,
	objName string,
	act constant.Action,
	eff bool,
) error {
	return askDBForPermissions(idb, haveToOpenConnection, roleName, domainName, objName, act,
		func(_ *gorm.DB, _ model.Permission) error {
			return errors.New("the assignation already exists")
		},
		func(db *gorm.DB, permission model.Permission) error {
			db.Error = nil
			permission.Effect = eff
			res := db.Create(&permission)
			return res.Error
		},
	)
}

//RemovePermission :
// remove the permission from the mia
// returns an error if
//	- the permission does not exist
func RemovePermission(
	idb database.MIADatabase,
	haveToOpenConnection bool,
	roleName string,
	domainName string,
	objName string,
	act constant.Action,
) error {
	return askDBForPermissions(idb, haveToOpenConnection, roleName, domainName, objName, act,
		func(db *gorm.DB, permission model.Permission) error {
			res := db.Delete(&permission)
			return res.Error
		},
		func(_ *gorm.DB, _ model.Permission) error {
			return errors.New("the assignation does not exist")
		},
	)
}
