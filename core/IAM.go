//Package core :
// This file contains all the API exposed that we will used
package core

import (
	"iam/core/constant"
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

// AddRole :
// See details in iam/core/engine/roles.go
func (iam IAM) AddRole(name string) error {
	return engine.AddRole(iam.idb, name)
}

// RemoveRole :
// See details in iam/core/engine/roles.go
func (iam IAM) RemoveRole(name string) error {
	return engine.RemoveRole(iam.idb, name)
}

// AddSubject :
// See details in iam/core/engine/items.go
func (iam IAM) AddSubject(name string) error {
	return engine.AddItem(
		iam.idb,
		model.ITEM_TYPE_SUBJ,
		name,
	)
}

// RemoveSubject :
// See details in iam/core/engine/items.go
func (iam IAM) RemoveSubject(name string) error {
	return engine.RemoveItem(
		iam.idb,
		model.ITEM_TYPE_SUBJ,
		name,
	)
}

// RenameSubject :
// See details in iam/core/engine/items.go
func (iam IAM) RenameSubject(name string, newName string) error {
	return engine.RenameItem(
		iam.idb,
		model.ITEM_TYPE_SUBJ,
		name,
		newName,
	)
}

// GetSubject :
// See details in iam/core/engine/items.go
func (iam IAM) GetSubject(name string) (model.Item, error) {
	return engine.GetItem(
		iam.idb,
		model.ITEM_TYPE_SUBJ,
		name,
	)
}

// AddSubjectLink :
// See details in iam/core/engine/items.go
func (iam IAM) AddSubjectLink(nP string, nC string) error {
	return engine.AddItemLink(
		iam.idb,
		model.ITEM_TYPE_SUBJ,
		nP,
		nC,
	)
}

// RemoveSubjectLink :
// See details in iam/core/engine/items.go
func (iam IAM) RemoveSubjectLink(nP string, nC string) error {
	return engine.RemoveItemLink(
		iam.idb,
		model.ITEM_TYPE_SUBJ,
		nP,
		nC,
	)
}

// AddSubjectArchitecture :
// See details in iam/core/engine/items.go
func (iam IAM) AddSubjectArchitecture(
	parents []string,
	childs []string,
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
// See details in iam/core/engine/items.go
func (iam IAM) AddObject(name string) error {
	return engine.AddItem(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		name,
	)
}

// RemoveObject :
// See details in iam/core/engine/items.go
func (iam IAM) RemoveObject(name string) error {
	return engine.RemoveItem(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		name,
	)
}

// RenameObject :
// See details in iam/core/engine/items.go
func (iam IAM) RenameObject(name string, newName string) error {
	return engine.RenameItem(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		name,
		newName,
	)
}

// GetObject :
// See details in iam/core/engine/items.go
func (iam IAM) GetObject(name string) (model.Item, error) {
	return engine.GetItem(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		name,
	)
}

// AddObjectLink :
// See details in iam/core/engine/items.go
func (iam IAM) AddObjectLink(nameP string, nameC string) error {
	return engine.AddItemLink(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		nameP,
		nameC,
	)
}

//RemoveObjectLink :
// See details in iam/core/engine/items.go
func (iam IAM) RemoveObjectLink(nameP string, nameC string) error {
	return engine.RemoveItemLink(
		iam.idb,
		model.ITEM_TYPE_OBJ,
		nameP,
		nameC,
	)
}

// AddObjectArchitecture :
// See details in iam/core/engine/items.go
func (iam IAM) AddObjectArchitecture(
	parents []string,
	childs []string,
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
// See details in iam/core/engine/items.go
func (iam IAM) AddDomain(name string) error {
	return engine.AddItem(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		name,
	)
}

// RemoveDomain :
// See details in iam/core/engine/items.go
func (iam IAM) RemoveDomain(name string) error {
	return engine.RemoveItem(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		name,
	)
}

// RenameDomain :
// See details in iam/core/engine/items.go
func (iam IAM) RenameDomain(name string, newName string) error {
	return engine.RenameItem(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		name,
		newName,
	)
}

// GetDomain :
// See details in iam/core/engine/items.go
func (iam IAM) GetDomain(name string) (model.Item, error) {
	return engine.GetItem(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		name,
	)
}

// AddDomainLink :
// See details in iam/core/engine/items.go
func (iam IAM) AddDomainLink(nameP string, nameC string) error {
	return engine.AddItemLink(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		nameP,
		nameC,
	)
}

// RemoveDomainLink :
// See details in iam/core/engine/items.go
func (iam IAM) RemoveDomainLink(nameP string, nameC string) error {
	return engine.RemoveItemLink(
		iam.idb,
		model.ITEM_TYPE_DOMAIN,
		nameP,
		nameC,
	)
}

// AddDomainArchitecture :
// See details in iam/core/engine/items.go
func (iam IAM) AddDomainArchitecture(
	parents []string,
	childs []string,
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

//AddAssignation :
// See details in iam/core/engine/assignations.go
func (iam IAM) AddAssignation(
	roleName string,
	subjName string,
	domainName string,
) error {
	return engine.AddAssignation(
		iam.idb,
		roleName,
		subjName,
		domainName,
	)
}

//RemoveAssignation :
// See details in iam/core/engine/assignations.go
func (iam IAM) RemoveAssignation(
	roleName string,
	subjName string,
	domainName string,
) error {
	return engine.RemoveAssignation(
		iam.idb,
		roleName,
		subjName,
		domainName,
	)
}

//AddPermission :
// See details in iam/core/engine/permissions.go
func (iam IAM) AddPermission(
	roleName string,
	domainName string,
	objName string,
	act constant.Action,
) error {
	return engine.AddPermission(
		iam.idb,
		roleName,
		domainName,
		objName,
		act,
	)
}

//RemovePermission :
// See details in iam/core/engine/permissions.go
func (iam IAM) RemovePermission(
	roleName string,
	domainName string,
	objName string,
	act constant.Action,
) error {
	return engine.RemovePermission(
		iam.idb,
		roleName,
		domainName,
		objName,
		act,
	)
}
