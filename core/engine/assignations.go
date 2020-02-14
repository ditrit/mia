// Package engine :
// Handling assignations queries from IAM
package engine

import (
	"iam/core/database"
)

//AddAssignation :
// adds an assignation in the iam
// returns an error if
//	- one of the role, the subject or the domain does not exist
//	- the assignation already exists
func AddAssignation(
	idb database.IAMDatabase,
	roleName string,
	subjName string,
	domainName string,
) error {
	//TODO
	return nil
}

//RemoveAssignation :
// remove the assignation from the iam
// returns an error if
//	- the assignation does not exist
func RemoveAssignation(
	idb database.IAMDatabase,
	roleName string,
	subjName string,
	domainName string,
) error {
	//TODO
	return nil
}
