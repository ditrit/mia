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

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//							INITIALIZERS							//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

// NewIAM :
// Constructor : Declare a IAM Object
func NewIAM(p string, dropTables bool) IAM {
	res := IAM{
		idb: database.NewIAMDatabase(p),
	}

	res.idb.Setup(dropTables)

	return res
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//								ROLES								//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

// AddRole :
// See details in iam/core/engine/roles.go
func (iam IAM) AddRole(name string) error {
	return engine.AddRole(iam.idb, true, name)
}

// RemoveRole :
// See details in iam/core/engine/roles.go
func (iam IAM) RemoveRole(name string) error {
	return engine.RemoveRole(iam.idb, true, name)
}

// GetRole :
// See details in iam/core/engine/roles.go
func (iam IAM) GetRole(name string) (model.Role, error) {
	return engine.GetRole(iam.idb, true, name)
}

// IsRoleExists :
// See details in iam/core/engine/roles.go
func (iam IAM) IsRoleExists(name string) (bool, error) {
	return engine.IsRoleExists(iam.idb, true, name)
}

// GetAllRoles :
// See details in iam/core/engine/roles.go
func (iam IAM) GetAllRoles() ([]string, error) {
	return engine.GetAllRoles(iam.idb, true)
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//							SUBJECTS								//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

// AddSubject :
// See details in iam/core/engine/items.go
func (iam IAM) AddSubject(name string, nameParent string) error {
	return engine.AddItem(
		iam.idb,
		true,
		model.ITEM_TYPE_SUBJ,
		name,
		nameParent,
	)
}

// AddSubjectToRoot :
// See details in iam/core/engine/items.go
// add subject litteraly under roots
func (iam IAM) AddSubjectToRoot(name string) error {
	rootName, _ := model.GetRootNameWithType(model.ITEM_TYPE_SUBJ)

	return engine.AddItem(
		iam.idb,
		true,
		model.ITEM_TYPE_SUBJ,
		name,
		rootName,
	)
}

// RemoveSubject :
// See details in iam/core/engine/items.go
func (iam IAM) RemoveSubject(name string) error {
	return engine.RemoveItem(
		iam.idb,
		true,
		model.ITEM_TYPE_SUBJ,
		name,
	)
}

// RenameSubject :
// See details in iam/core/engine/items.go
func (iam IAM) RenameSubject(name string, newName string) error {
	return engine.RenameItem(
		iam.idb,
		true,
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
		true,
		model.ITEM_TYPE_SUBJ,
		name,
	)
}

// AddSubjectLink :
// See details in iam/core/engine/items.go
func (iam IAM) AddSubjectLink(nP string, nC string) error {
	return engine.AddItemLink(
		iam.idb,
		true,
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
		true,
		model.ITEM_TYPE_SUBJ,
		nP,
		nC,
	)
}

// GetSubjectArchitecture :
// See details in iam/core/engine/items.go
func (iam IAM) GetSubjectArchitecture() ([]string, map[string][]string, error) {
	return engine.GetItemArchitectureNameOnly(iam.idb, true, model.ITEM_TYPE_SUBJ)
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//								OBJECTS								//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

// AddObject :
// See details in iam/core/engine/items.go
func (iam IAM) AddObject(name string, nameParent string) error {
	return engine.AddItem(
		iam.idb,
		true,
		model.ITEM_TYPE_OBJ,
		name,
		nameParent,
	)
}

// AddObjectToRoot :
// See details in iam/core/engine/items.go
// add object litteraly under roots
func (iam IAM) AddObjectToRoot(name string) error {
	rootName, _ := model.GetRootNameWithType(model.ITEM_TYPE_OBJ)

	return engine.AddItem(
		iam.idb,
		true,
		model.ITEM_TYPE_OBJ,
		name,
		rootName,
	)
}

// RemoveObject :
// See details in iam/core/engine/items.go
func (iam IAM) RemoveObject(name string) error {
	return engine.RemoveItem(
		iam.idb,
		true,
		model.ITEM_TYPE_OBJ,
		name,
	)
}

// RenameObject :
// See details in iam/core/engine/items.go
func (iam IAM) RenameObject(name string, newName string) error {
	return engine.RenameItem(
		iam.idb,
		true,
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
		true,
		model.ITEM_TYPE_OBJ,
		name,
	)
}

// AddObjectLink :
// See details in iam/core/engine/items.go
func (iam IAM) AddObjectLink(nameP string, nameC string) error {
	return engine.AddItemLink(
		iam.idb,
		true,
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
		true,
		model.ITEM_TYPE_OBJ,
		nameP,
		nameC,
	)
}

// GetObjectArchitecture :
// See details in iam/core/engine/items.go
func (iam IAM) GetObjectArchitecture() ([]string, map[string][]string, error) {
	return engine.GetItemArchitectureNameOnly(iam.idb, true, model.ITEM_TYPE_OBJ)
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//								DOMAINS								//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

// AddDomain :
// See details in iam/core/engine/items.go
func (iam IAM) AddDomain(name string, nameParent string) error {
	return engine.AddItem(
		iam.idb,
		true,
		model.ITEM_TYPE_DOMAIN,
		name,
		nameParent,
	)
}

// AddDomainToRoot :
// See details in iam/core/engine/items.go
// add domain litteraly under roots
func (iam IAM) AddDomainToRoot(name string) error {
	rootName, _ := model.GetRootNameWithType(model.ITEM_TYPE_DOMAIN)

	return engine.AddItem(
		iam.idb,
		true,
		model.ITEM_TYPE_DOMAIN,
		name,
		rootName,
	)
}

// RemoveDomain :
// See details in iam/core/engine/items.go
func (iam IAM) RemoveDomain(name string) error {
	return engine.RemoveItem(
		iam.idb,
		true,
		model.ITEM_TYPE_DOMAIN,
		name,
	)
}

// RenameDomain :
// See details in iam/core/engine/items.go
func (iam IAM) RenameDomain(name string, newName string) error {
	return engine.RenameItem(
		iam.idb,
		true,
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
		true,
		model.ITEM_TYPE_DOMAIN,
		name,
	)
}

// AddDomainLink :
// See details in iam/core/engine/items.go
func (iam IAM) AddDomainLink(nameP string, nameC string) error {
	return engine.AddItemLink(
		iam.idb,
		true,
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
		true,
		model.ITEM_TYPE_DOMAIN,
		nameP,
		nameC,
	)
}

// GetDomainArchitecture :
// See details in iam/core/engine/items.go
func (iam IAM) GetDomainArchitecture() ([]string, map[string][]string, error) {
	return engine.GetItemArchitectureNameOnly(iam.idb, true, model.ITEM_TYPE_DOMAIN)
}

//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//
//	//								ASSIGNMENTS							//	//
//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//	//

//AddAssignment :
// See details in iam/core/engine/assignments.go
func (iam IAM) AddAssignment(
	roleName string,
	subjName string,
	domainName string,
) error {
	return engine.AddAssignment(
		iam.idb,
		true,
		roleName,
		subjName,
		domainName,
	)
}

//RemoveAssignment :
// See details in iam/core/engine/assignments.go
func (iam IAM) RemoveAssignment(
	roleName string,
	subjName string,
	domainName string,
) error {
	return engine.RemoveAssignment(
		iam.idb,
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
// See details in iam/core/engine/permissions.go
func (iam IAM) AddPermission(
	roleName string,
	domainName string,
	objName string,
	act constant.Action,
	eff bool,
) error {
	return engine.AddPermission(
		iam.idb,
		true,
		roleName,
		domainName,
		objName,
		act,
		eff,
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
// See details in iam/core/engine/enforce.go
func (iam IAM) Enforce(
	subjectName string,
	domainName string,
	objectName string,
	action constant.Action,
) (bool, error) {
	return engine.Enforce(
		iam.idb,
		subjectName,
		domainName,
		objectName,
		action,
	)
}
