// Package engine :
// Handling permissions queries from IAM
package engine

import (
	"errors"
	"iam/core/constant"
	"iam/core/database"
)

//AddPermission :
// adds an permission in the iam
// returns an error if
//	- one of the role, the subject or the domain does not exist
//	- the permission already exists
func AddPermission(
	idb database.IAMDatabase,
	roleName string,
	domainName string,
	objName string,
	act constant.Action,
) error {
	//TODO
	return errors.New("not implemented")
}

//RemovePermission :
// remove the permission from the iam
// returns an error if
//	- the permission does not exist
func RemovePermission(
	idb database.IAMDatabase,
	roleName string,
	domainName string,
	objName string,
	act constant.Action,
) error {
	//TODO
	return errors.New("not implemented")
}
