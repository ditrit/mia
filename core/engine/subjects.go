// Package engine :
// Handling subjects queries from IAM
package engine

import (
	"errors"
	"iam/core/database"
	"iam/core/model"

	"github.com/jinzhu/gorm"
)

// use this function to have the benefit of abstraction and closures
func askDBForSubjects(
	idb database.IAMDatabase,
	name string,
	fIfFound func(*gorm.DB, model.Subject) (model.Subject, error),
	fIfNotFound func(*gorm.DB, model.Subject) (model.Subject, error),
) (model.Subject, error) {
	var querySubject model.Subject

	idb.OpenConnection()
	defer idb.CloseConnection() //nolint: errcheck

	if len(name) == 0 {
		return querySubject, errors.New("the name cannot be empty")
	}

	res := idb.DB().Where("name = ?", name).First(&querySubject)

	if res.Error != nil && !res.RecordNotFound() {
		return querySubject, errors.New("unknown error occurred")
	}

	if res.RecordNotFound() {
		return fIfNotFound(res, querySubject)
	}

	return fIfFound(res, querySubject)
}

// AddSubject :
// add subject in the IAM
// returns an error if :
//	- the subject already exists in the iam
//	- the subject has an empty name
func AddSubject(
	idb database.IAMDatabase,
	s model.Subject,
) error {
	_, err := askDBForSubjects(idb, s.Name,
		func(_ *gorm.DB, qs model.Subject) (model.Subject, error) {
			return qs, errors.New("the subject already exists in the iam")
		},
		func(db *gorm.DB, qs model.Subject) (model.Subject, error) {
			db.Error = nil
			res := db.Create(&s)
			return qs, res.Error
		})

	return err
}

// RemoveSubject :
// remove subject in the IAM
// returns an error if :
//	- the subject does not exist in the iam
func RemoveSubject(
	idb database.IAMDatabase,
	s model.Subject,
) error {
	_, err := askDBForSubjects(idb, s.Name,
		func(_ *gorm.DB, qs model.Subject) (model.Subject, error) {
			idb.DB().Delete(&s)
			return qs, idb.DB().Error
		},
		func(db *gorm.DB, qs model.Subject) (model.Subject, error) {
			return qs, errors.New("the subject does not exist in the iam")
		})

	return err
}

// RenameSubject :
// rename subject in the IAM
// returns an error if :
//	- the subject does not exist in the iam
//	- the new name given is empty
func RenameSubject(
	idb database.IAMDatabase,
	s model.Subject,
	newName string,
) error {
	_, err := askDBForSubjects(idb, s.Name,
		func(_ *gorm.DB, qs model.Subject) (model.Subject, error) {
			idb.DB().Model(&s).Update("name", newName)
			return qs, idb.DB().Error
		},
		func(db *gorm.DB, qs model.Subject) (model.Subject, error) {
			return qs, errors.New("the subject does not exist in the iam")
		})

	return err
}

// GetSubject :
// get subject in the IAM
// returns an error if :
//	- the subject does not exist in the iam
func GetSubject(
	idb database.IAMDatabase,
	name string,
) (model.Subject, error) {
	return askDBForSubjects(idb, name,
		func(db *gorm.DB, qs model.Subject) (model.Subject, error) {
			return qs, db.Error
		},
		func(db *gorm.DB, qs model.Subject) (model.Subject, error) {
			return qs, errors.New("the subject does not exist in the iam")
		})
}

// AddSubjectLink :
// add a relation between two subjects in the IAM
// returns an error if :
//	- one of the subjects already exists in the iam
//	- one of the subjects has an empty name
//	- the link already exists
func AddSubjectLink(
	idb database.IAMDatabase,
	sParent model.Subject,
	sChild model.Subject,
) error {
	return nil
}

// RemoveSubjectLink :
// remove a relation between two subjects in the IAM
// returns an error if :
//	- the link does not exist
func RemoveSubjectLink(
	idb database.IAMDatabase,
	sParent model.Subject,
	sChild model.Subject,
) error {
	return nil
}

// AddSubjectArchitecture :
// add an architecture to the IAM
// returns an error if :
//	- tabs given have not the same size
//	- one of the subjects already exists in the iam
//	- one of the subjects has an empty name
//	- one of the links alrady exists
// We can ignore some of this error with the other parameters
func AddSubjectArchitecture(
	idb database.IAMDatabase,
	parents []model.Subject,
	childs []model.Subject,
	ignoreAlreadyExistsSubject bool,
	ignoreAlreadyExistsLinks bool,
) error {
	return nil
}
