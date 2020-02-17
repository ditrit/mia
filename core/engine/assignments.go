// Package engine :
// Handling assignments queries from IAM
package engine

import (
	"errors"
	"iam/core/database"
	"iam/core/model"

	"github.com/jinzhu/gorm"
)

func askDBForAssignments(
	idb database.IAMDatabase,
	haveToOpenConnection bool,
	roleName string,
	subjName string,
	domainName string,
	fIfFound func(*gorm.DB, model.Assignment) error,
	fIfNotFound func(*gorm.DB, model.Assignment) error,
) error {
	var (
		assign  model.Assignment
		role    model.Role
		subject model.Item
		domain  model.Item
	)

	if haveToOpenConnection {
		idb.OpenConnection()
		defer idb.CloseConnection() //nolint: errcheck
	}

	role, err := GetRole(idb, false, roleName)

	if err != nil {
		return err
	}

	subject, err = GetItem(idb, false, model.ITEM_TYPE_SUBJ, subjName)

	if err != nil {
		return err
	}

	domain, err = GetItem(idb, false, model.ITEM_TYPE_DOMAIN, domainName)

	if err != nil {
		return err
	}

	query := idb.DB().Where("id_role = ?", role.ID)
	query = query.Where("id_subject = ?", subject.ID)
	query = query.Where("id_domain = ?", domain.ID)
	res := query.Take(&assign)

	if res.Error != nil && !res.RecordNotFound() {
		return errors.New("unknown error occurred")
	}

	assign.IDRole = role.ID
	assign.IDSubject = subject.ID
	assign.IDDomain = domain.ID

	if res.RecordNotFound() {
		return fIfNotFound(res, assign)
	}

	return fIfFound(res, assign)
}

//AddAssignment :
// adds an assignment in the iam
// returns an error if
//	- one of the role, the subject or the domain does not exist
//	- the assignment already exists
func AddAssignment(
	idb database.IAMDatabase,
	haveToOpenConnection bool,
	roleName string,
	subjName string,
	domainName string,
) error {
	return askDBForAssignments(idb, haveToOpenConnection, roleName, subjName, domainName,
		func(_ *gorm.DB, _ model.Assignment) error {
			return errors.New("the assignation already exists")
		},
		func(db *gorm.DB, assign model.Assignment) error {
			db.Error = nil
			res := db.Create(&assign)
			return res.Error
		},
	)
}

//RemoveAssignment :
// remove the assignment from the iam
// returns an error if
//	- the assignment does not exist
func RemoveAssignment(
	idb database.IAMDatabase,
	haveToOpenConnection bool,
	roleName string,
	subjName string,
	domainName string,
) error {
	return askDBForAssignments(idb, haveToOpenConnection, roleName, subjName, domainName,
		func(db *gorm.DB, assign model.Assignment) error {
			res := db.Delete(&assign)
			return res.Error
		},
		func(_ *gorm.DB, _ model.Assignment) error {
			return errors.New("the assignation does not exist")
		},
	)
}
