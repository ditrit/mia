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
	haveToOpenConnection bool,
	fIfFound func(*gorm.DB, model.Subject) (model.Subject, error),
	fIfNotFound func(*gorm.DB, model.Subject) (model.Subject, error),
) (model.Subject, error) {
	var querySubject model.Subject

	if len(name) == 0 {
		return querySubject, errors.New("the name cannot be empty")
	}

	if haveToOpenConnection {
		idb.OpenConnection()
		defer idb.CloseConnection() //nolint: errcheck
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

func askDBForSubjectLinks(
	idb database.IAMDatabase,
	parentName string,
	childName string,
	haveToOpenConnection bool,
	fIfFound func(*gorm.DB, model.SubjectLink) error,
	fIfNotFound func(*gorm.DB, model.SubjectLink) error,
) error {
	var (
		sLink    model.SubjectLink
		parentDB model.Subject
		childDB  model.Subject
		err      error
	)

	if haveToOpenConnection {
		idb.OpenConnection()
		defer idb.CloseConnection() //nolint: errcheck
	}

	// Search parent subject
	parentDB, err = askDBForSubjects(idb, parentName, false,
		func(db *gorm.DB, qs model.Subject) (model.Subject, error) {
			return qs, db.Error
		},
		func(db *gorm.DB, qs model.Subject) (model.Subject, error) {
			return qs, errors.New("the parent subject does not exist in the iam")
		})

	if err != nil {
		return err
	}

	// Search child subject
	childDB, err = askDBForSubjects(idb, childName, false,
		func(db *gorm.DB, qs model.Subject) (model.Subject, error) {
			return qs, db.Error
		},
		func(db *gorm.DB, qs model.Subject) (model.Subject, error) {
			return qs, errors.New("the child subject does not exist in the iam")
		})

	if err != nil {
		return err
	}

	res := idb.DB().Where("id_subject_parent = ?", parentDB.ID).Where("id_subject_child = ?", childDB.ID).Take(&sLink)
	if res.Error != nil && !res.RecordNotFound() {
		return res.Error
	}

	sLink.IDSubjectParent = parentDB.ID
	sLink.IDSubjectChild = childDB.ID

	if res.RecordNotFound() {
		return fIfNotFound(res, sLink)
	}

	return fIfFound(res, sLink)
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
	_, err := askDBForSubjects(idb, s.Name, true,
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
//	- the subject is a parent or a child in subjectLinks
func RemoveSubject(
	idb database.IAMDatabase,
	s model.Subject,
) error {
	_, err := askDBForSubjects(idb, s.Name, true,
		func(db *gorm.DB, qs model.Subject) (model.Subject, error) {
			res := db.Delete(&s)
			return qs, res.Error
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
	if newName == "" {
		return errors.New("the new name cannot be empty")
	}

	_, err := askDBForSubjects(idb, s.Name, true,
		func(db *gorm.DB, qs model.Subject) (model.Subject, error) {
			res := db.Model(&s).Update("name", newName)
			return qs, res.Error
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
	return askDBForSubjects(idb, name, true,
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
//	- one of the subjects does not exists in the iam
//	- the link already exists
func AddSubjectLink(
	idb database.IAMDatabase,
	sParent model.Subject,
	sChild model.Subject,
) error {
	return askDBForSubjectLinks(idb, sParent.Name, sChild.Name, true,
		func(db *gorm.DB, qs model.SubjectLink) error {
			return errors.New("the connection link already exists")
		},
		func(db *gorm.DB, qs model.SubjectLink) error {
			db.Error = nil
			res := db.Create(&qs)
			return res.Error
		},
	)
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
	return askDBForSubjectLinks(idb, sParent.Name, sChild.Name, true,
		func(db *gorm.DB, qs model.SubjectLink) error {
			res := db.Delete(&qs)
			return res.Error
		},
		func(db *gorm.DB, qs model.SubjectLink) error {
			return errors.New("the connection link does not exist")
		},
	)
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
	// TODO: implement + test
	return nil
}
