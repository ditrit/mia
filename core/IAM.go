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
		model.ITEM_TYPE_SUBJ,
		s,
	)
}

// RemoveSubject :
// See details in iam/core/engine/subjects.go
func (iam IAM) RemoveSubject(s model.Item) error {
	return engine.RemoveItem(
		iam.idb,
		model.ITEM_TYPE_SUBJ,
		s,
	)
}

// RenameSubject :
// See details in iam/core/engine/subjects.go
func (iam IAM) RenameSubject(s model.Item, newName string) error {
	return engine.RenameItem(
		iam.idb,
		model.ITEM_TYPE_SUBJ,
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
		model.ITEM_TYPE_SUBJ,
		sParent,
		sChild,
	)
}

// RemoveSubjectLink :
// See details in iam/core/engine/subjects.go
func (iam IAM) RemoveSubjectLink(sParent model.Item, sChild model.Item) error {
	return engine.RemoveItemLink(
		iam.idb,
		model.ITEM_TYPE_SUBJ,
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
		model.ITEM_TYPE_SUBJ,
		parents,
		childs,
		ignoreAlreadyExistsSubject,
		ignoreAlreadyExistsSubjectLinks,
	)
}

// AddObject :
// See details in iam/core/engine/objects.go
func (iam IAM) AddObject(s model.Item) error {
	return engine.AddItem(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		s,
	)
}

// RemoveObject :
// See details in iam/core/engine/objects.go
func (iam IAM) RemoveObject(s model.Item) error {
	return engine.RemoveItem(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		s,
	)
}

// RenameObject :
// See details in iam/core/engine/objects.go
func (iam IAM) RenameObject(s model.Item, newName string) error {
	return engine.RenameItem(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		s,
		newName,
	)
}

// GetObject :
// See details in iam/core/engine/objects.go
func (iam IAM) GetObject(name string) (model.Item, error) {
	return engine.GetItem(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		name,
	)
}

// AddObjectLink :
// See details in iam/core/engine/objects.go
func (iam IAM) AddObjectLink(sParent model.Item, sChild model.Item) error {
	return engine.AddItemLink(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		sParent,
		sChild,
	)
}

//RemoveObjectLink :
// See details in iam/core/engine/objects.go
func (iam IAM) RemoveObjectLink(sParent model.Item, sChild model.Item) error {
	return engine.RemoveItemLink(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		sParent,
		sChild,
	)
}

// AddObjectArchitecture :
// See details in iam/core/engine/objects.go
func (iam IAM) AddObjectArchitecture(
	parents []model.Item,
	childs []model.Item,
	ignoreAlreadyExistsObject bool,
	ignoreAlreadyExistsObjectLinks bool,
) error {
	return engine.AddItemArchitecture(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		parents,
		childs,
		ignoreAlreadyExistsObject,
		ignoreAlreadyExistsObjectLinks,
	)
}

// AddDomain :
// See details in iam/core/engine/objects.go
func (iam IAM) AddDomain(s model.Item) error {
	return engine.AddItem(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		s,
	)
}

// RemoveDomain :
// See details in iam/core/engine/objects.go
func (iam IAM) RemoveDomain(s model.Item) error {
	return engine.RemoveItem(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		s,
	)
}

// RenameDomain :
// See details in iam/core/engine/objects.go
func (iam IAM) RenameDomain(s model.Item, newName string) error {
	return engine.RenameItem(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		s,
		newName,
	)
}

// GetDomain :
// See details in iam/core/engine/objects.go
func (iam IAM) GetDomain(name string) (model.Item, error) {
	return engine.GetItem(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		name,
	)
}

// AddDomainLink :
// See details in iam/core/engine/objects.go
func (iam IAM) AddDomainLink(sParent model.Item, sChild model.Item) error {
	return engine.AddItemLink(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		sParent,
		sChild,
	)
}

// RemoveDomainLink :
// See details in iam/core/engine/objects.go
func (iam IAM) RemoveDomainLink(sParent model.Item, sChild model.Item) error {
	return engine.RemoveItemLink(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		sParent,
		sChild,
	)
}

// AddDomainArchitecture :
// See details in iam/core/engine/objects.go
func (iam IAM) AddDomainArchitecture(
	parents []model.Item,
	childs []model.Item,
	ignoreAlreadyExistsObject bool,
	ignoreAlreadyExistsObjectLinks bool,
) error {
	return engine.AddItemArchitecture(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		parents,
		childs,
		ignoreAlreadyExistsObject,
		ignoreAlreadyExistsObjectLinks,
	)
}
