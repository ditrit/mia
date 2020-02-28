//Package core :
// This file contains all the API exposed that we will used
package core

import (
	"mia/core/constant"
	"mia/core/database"
	"mia/core/engine"
	"mia/core/model"
)

// MIA :
// The abstract struct that users will work with
type MIA struct {
	idb database.MIADatabase
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//							INITIALIZERS							//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

// NewMIA :
// Constructor : Declare a MIA Object
func NewMIA(p string, dropTables bool) MIA {
	res := MIA{
		idb: database.NewMIADatabase(p),
	}

	res.idb.Setup(dropTables)

	return res
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//								ROLES								//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

// AddRole :
// See details in mia/core/engine/roles.go
func (mia MIA) AddRole(name string) error {
	return engine.AddRole(mia.idb, true, name)
}

// RemoveRole :
// See details in mia/core/engine/roles.go
func (mia MIA) RemoveRole(name string) error {
	return engine.RemoveRole(mia.idb, true, name)
}

// GetRole :
// See details in mia/core/engine/roles.go
func (mia MIA) GetRole(name string) (model.Role, error) {
	return engine.GetRole(mia.idb, true, name)
}

// IsRoleExists :
// See details in mia/core/engine/roles.go
func (mia MIA) IsRoleExists(name string) (bool, error) {
	return engine.IsRoleExists(mia.idb, true, name)
}

// GetAllRoles :
// See details in mia/core/engine/roles.go
func (mia MIA) GetAllRoles() ([]string, error) {
	return engine.GetAllRoles(mia.idb, true)
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//							SUBJECTS								//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

// AddSubject :
// See details in mia/core/engine/items.go
func (mia MIA) AddSubject(name string, nameParent string) error {
	return engine.AddItem(
		mia.idb,
		true,
		model.ITEM_TYPE_SUBJ,
		name,
		nameParent,
	)
}

// AddSubjectToRoot :
// See details in mia/core/engine/items.go
// add subject litteraly under roots
func (mia MIA) AddSubjectToRoot(name string) error {
	rootName, _ := model.GetRootNameWithType(model.ITEM_TYPE_SUBJ)

	return engine.AddItem(
		mia.idb,
		true,
		model.ITEM_TYPE_SUBJ,
		name,
		rootName,
	)
}

// RemoveSubject :
// See details in mia/core/engine/items.go
func (mia MIA) RemoveSubject(name string) error {
	return engine.RemoveItem(
		mia.idb,
		true,
		model.ITEM_TYPE_SUBJ,
		name,
	)
}

// RenameSubject :
// See details in mia/core/engine/items.go
func (mia MIA) RenameSubject(name string, newName string) error {
	return engine.RenameItem(
		mia.idb,
		true,
		model.ITEM_TYPE_SUBJ,
		name,
		newName,
	)
}

// GetSubject :
// See details in mia/core/engine/items.go
func (mia MIA) GetSubject(name string) (model.Item, error) {
	return engine.GetItem(
		mia.idb,
		true,
		model.ITEM_TYPE_SUBJ,
		name,
	)
}

// AddSubjectLink :
// See details in mia/core/engine/items.go
func (mia MIA) AddSubjectLink(nP string, nC string) error {
	return engine.AddItemLink(
		mia.idb,
		true,
		model.ITEM_TYPE_SUBJ,
		nP,
		nC,
	)
}

// RemoveSubjectLink :
// See details in mia/core/engine/items.go
func (mia MIA) RemoveSubjectLink(nP string, nC string) error {
	return engine.RemoveItemLink(
		mia.idb,
		true,
		model.ITEM_TYPE_SUBJ,
		nP,
		nC,
	)
}

// GetSubjectArchitecture :
// See details in mia/core/engine/items.go
func (mia MIA) GetSubjectArchitecture() ([]string, map[string][]string, error) {
	return engine.GetItemArchitectureNameOnly(mia.idb, true, model.ITEM_TYPE_SUBJ)
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//								OBJECTS								//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

// AddObject :
// See details in mia/core/engine/items.go
func (mia MIA) AddObject(name string, nameParent string) error {
	return engine.AddItem(
		mia.idb,
		true,
		model.ITEM_TYPE_OBJ,
		name,
		nameParent,
	)
}

// AddObjectToRoot :
// See details in mia/core/engine/items.go
// add object litteraly under roots
func (mia MIA) AddObjectToRoot(name string) error {
	rootName, _ := model.GetRootNameWithType(model.ITEM_TYPE_OBJ)

	return engine.AddItem(
		mia.idb,
		true,
		model.ITEM_TYPE_OBJ,
		name,
		rootName,
	)
}

// RemoveObject :
// See details in mia/core/engine/items.go
func (mia MIA) RemoveObject(name string) error {
	return engine.RemoveItem(
		mia.idb,
		true,
		model.ITEM_TYPE_OBJ,
		name,
	)
}

// RenameObject :
// See details in mia/core/engine/items.go
func (mia MIA) RenameObject(name string, newName string) error {
	return engine.RenameItem(
		mia.idb,
		true,
		model.ITEM_TYPE_OBJ,
		name,
		newName,
	)
}

// GetObject :
// See details in mia/core/engine/items.go
func (mia MIA) GetObject(name string) (model.Item, error) {
	return engine.GetItem(
		mia.idb,
		true,
		model.ITEM_TYPE_OBJ,
		name,
	)
}

// AddObjectLink :
// See details in mia/core/engine/items.go
func (mia MIA) AddObjectLink(nameP string, nameC string) error {
	return engine.AddItemLink(
		mia.idb,
		true,
		model.ITEM_TYPE_OBJ,
		nameP,
		nameC,
	)
}

//RemoveObjectLink :
// See details in mia/core/engine/items.go
func (mia MIA) RemoveObjectLink(nameP string, nameC string) error {
	return engine.RemoveItemLink(
		mia.idb,
		true,
		model.ITEM_TYPE_OBJ,
		nameP,
		nameC,
	)
}

// GetObjectArchitecture :
// See details in mia/core/engine/items.go
func (mia MIA) GetObjectArchitecture() ([]string, map[string][]string, error) {
	return engine.GetItemArchitectureNameOnly(mia.idb, true, model.ITEM_TYPE_OBJ)
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//								DOMAINS								//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

// AddDomain :
// See details in mia/core/engine/items.go
func (mia MIA) AddDomain(name string, nameParent string) error {
	return engine.AddItem(
		mia.idb,
		true,
		model.ITEM_TYPE_DOMAIN,
		name,
		nameParent,
	)
}

// AddDomainToRoot :
// See details in mia/core/engine/items.go
// add domain litteraly under roots
func (mia MIA) AddDomainToRoot(name string) error {
	rootName, _ := model.GetRootNameWithType(model.ITEM_TYPE_DOMAIN)

	return engine.AddItem(
		mia.idb,
		true,
		model.ITEM_TYPE_DOMAIN,
		name,
		rootName,
	)
}

// RemoveDomain :
// See details in mia/core/engine/items.go
func (mia MIA) RemoveDomain(name string) error {
	return engine.RemoveItem(
		mia.idb,
		true,
		model.ITEM_TYPE_DOMAIN,
		name,
	)
}

// RenameDomain :
// See details in mia/core/engine/items.go
func (mia MIA) RenameDomain(name string, newName string) error {
	return engine.RenameItem(
		mia.idb,
		true,
		model.ITEM_TYPE_DOMAIN,
		name,
		newName,
	)
}

// GetDomain :
// See details in mia/core/engine/items.go
func (mia MIA) GetDomain(name string) (model.Item, error) {
	return engine.GetItem(
		mia.idb,
		true,
		model.ITEM_TYPE_DOMAIN,
		name,
	)
}

// AddDomainLink :
// See details in mia/core/engine/items.go
func (mia MIA) AddDomainLink(nameP string, nameC string) error {
	return engine.AddItemLink(
		mia.idb,
		true,
		model.ITEM_TYPE_DOMAIN,
		nameP,
		nameC,
	)
}

// RemoveDomainLink :
// See details in mia/core/engine/items.go
func (mia MIA) RemoveDomainLink(nameP string, nameC string) error {
	return engine.RemoveItemLink(
		mia.idb,
		true,
		model.ITEM_TYPE_DOMAIN,
		nameP,
		nameC,
	)
}

// GetDomainArchitecture :
// See details in mia/core/engine/items.go
func (mia MIA) GetDomainArchitecture() ([]string, map[string][]string, error) {
	return engine.GetItemArchitectureNameOnly(mia.idb, true, model.ITEM_TYPE_DOMAIN)
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//								ASSIGNMENTS							//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

//AddAssignment :
// See details in mia/core/engine/assignments.go
func (mia MIA) AddAssignment(
	roleName string,
	subjName string,
	domainName string,
) error {
	return engine.AddAssignment(
		mia.idb,
		true,
		roleName,
		subjName,
		domainName,
	)
}

//RemoveAssignment :
// See details in mia/core/engine/assignments.go
func (mia MIA) RemoveAssignment(
	roleName string,
	subjName string,
	domainName string,
) error {
	return engine.RemoveAssignment(
		mia.idb,
		true,
		roleName,
		subjName,
		domainName,
	)
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//							PERMISSIONS								//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

//AddPermission :
// See details in mia/core/engine/permissions.go
func (mia MIA) AddPermission(
	roleName string,
	domainName string,
	objName string,
	act constant.Action,
	eff bool,
) error {
	return engine.AddPermission(
		mia.idb,
		true,
		roleName,
		domainName,
		objName,
		act,
		eff,
	)
}

//RemovePermission :
// See details in mia/core/engine/permissions.go
func (mia MIA) RemovePermission(
	roleName string,
	domainName string,
	objName string,
	act constant.Action,
) error {
	return engine.RemovePermission(
		mia.idb,
		true,
		roleName,
		domainName,
		objName,
		act,
	)
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//							ENFORCE									//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

//Enforce :
// See details in mia/core/engine/enforce.go
func (mia MIA) Enforce(
	subjectName string,
	domainName string,
	objectName string,
	action constant.Action,
) (bool, error) {
	return engine.Enforce(
		mia.idb,
		subjectName,
		domainName,
		objectName,
		action,
	)
}
