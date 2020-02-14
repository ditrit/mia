// Package engine :
// Handling roles queries from IAM
package engine

import (
	"iam/core/database"
)

//AddRole :
// add role in the IAM
// returns an error if :
//	- the role already exists
//	- the name is not conform
func AddRole(
	idb database.IAMDatabase,
	name string,
) error {
	//TODO
	return nil
}

//RemoveRole :
// remove role in the IAM
// returns an error if :
//	- the role does not exist in the iam
//	- the role is present in a assignment
func RemoveRole(
	idb database.IAMDatabase,
	name string,
) error {
	//TODO
	return nil
}
