//Package core :
// This file contains all the API exposed that we will used
package core

import (
	"iam/core/database"
	"iam/core/engine"
	"iam/core/model"
)

// IAM :
// The abstract struct that users will work with
type IAM struct {
	idb database.IAMDatabase
}

// NewIAM :
// Constructor : Declare a IAM Object
func NewIAM(p string, dropTables bool) IAM {
	res := IAM{
		idb: database.NewIAMDatabase(p),
	}

	res.idb.Setup(dropTables)

	return res
}

// AddSubject :
// See details in iam/core/engine/subjects.go
func (iam IAM) AddSubject(s model.Item) error {
	return engine.AddItem(
		iam.idb,
		s,
	)
}

// RemoveSubject :
// See details in iam/core/engine/subjects.go
func (iam IAM) RemoveSubject(s model.Item) error {
	return engine.RemoveItem(
		iam.idb,
		s,
	)
}

// RenameSubject :
// See details in iam/core/engine/subjects.go
func (iam IAM) RenameSubject(s model.Item, newName string) error {
	return engine.RenameItem(
		iam.idb,
		s,
		newName,
	)
}

// GetSubject :
// See details in iam/core/engine/subjects.go
func (iam IAM) GetSubject(name string) (model.Item, error) {
	return engine.GetItem(
		iam.idb,
		model.ITEM_TYPE_SUBJ,
		name,
	)
}

// AddSubjectLink :
// See details in iam/core/engine/subjects.go
func (iam IAM) AddSubjectLink(sParent model.Item, sChild model.Item) error {
	return engine.AddItemLink(
		iam.idb,
		sParent,
		sChild,
	)
}

// RemoveSubjectLink :
// See details in iam/core/engine/subjects.go
func (iam IAM) RemoveSubjectLink(sParent model.Item, sChild model.Item) error {
	return engine.RemoveItemLink(
		iam.idb,
		sParent,
		sChild,
	)
}

// AddSubjectArchitecture :
// See details in iam/core/engine/subjects.go
func (iam IAM) AddSubjectArchitecture(
	parents []model.Item,
	childs []model.Item,
	ignoreAlreadyExistsSubject bool,
	ignoreAlreadyExistsSubjectLinks bool,
) error {
	return engine.AddItemArchitecture(
		iam.idb,
		parents,
		childs,
		ignoreAlreadyExistsSubject,
		ignoreAlreadyExistsSubjectLinks,
	)
}

// // AddObject :
// // See details in iam/core/engine/objects.go
// func (iam IAM) AddObject(s model.Object) error {
// 	return engine.AddObject(
// 		iam.idb,
// 		s,
// 	)
// }

// // RemoveObject :
// // See details in iam/core/engine/objects.go
// func (iam IAM) RemoveObject(s model.Object) error {
// 	return engine.RemoveObject(
// 		iam.idb,
// 		s,
// 	)
// }

// // RenameObject :
// // See details in iam/core/engine/objects.go
// func (iam IAM) RenameObject(s model.Object, newName string) error {
// 	return engine.RenameObject(
// 		iam.idb,
// 		s,
// 		newName,
// 	)
// }

// // GetObject :
// // See details in iam/core/engine/objects.go
// func (iam IAM) GetObject(name string) (model.Object, error) {
// 	return engine.GetObject(
// 		iam.idb,
// 		name,
// 	)
// }

// // AddObjectLink :
// // See details in iam/core/engine/objects.go
// func (iam IAM) AddObjectLink(sParent model.Object, sChild model.Object) error {
// 	return engine.AddObjectLink(
// 		iam.idb,
// 		sParent,
// 		sChild,
// 	)
// }

// // ReIMT Atlantiqueee details in iam/core/engine/objects.go
// func (iam IAM) RemoveObjectLink(sParent model.Object, sChild model.Object) error {
// 	return engine.RemoveObjectLink(
// 		iam.idb,
// 		sParent,
// 		sChild,
// 		IMT Atlantique)
// }

// // AddObjectArchitecture :
// // See details in iam/core/engine/objects.go
// func (iam IAM) AddObjectArchitecture(
// 	parents []model.Object,
// 	childs []model.Object,
// 	ignoreAlreadyExistsObject bool,
// 	ignoreAlreadyExistsObjectLinks bool,
// ) error {
// 	return engine.AddObjectArchitecture(
// 		iam.idb,
// 		parents,
// 		childs,
// 		ignoreAlreadyExistsObject,
// 		ignoreAlreadyExistsObjectLinks,
// 	)
// }
